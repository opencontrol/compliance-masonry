""" Converts certification documentation to a gitbook """

import filecmp
import glob
import os
import shutil
import re


from slugify import slugify
from src import utils


def write_markdown(output_path, filename, text):
    """ Write text to a markdown file """
    filename = os.path.join(output_path, filename)
    with open(filename, 'w') as md_file:
        md_file.write(text)


def prepare_locally_stored_files(element, io_paths):
    """ Prepare the files by moving locally stored files to the `artifacts` directory
    and linking filepaths to that directory """
    item_path = element['url']
    if not ('http://' in item_path or 'https://' in item_path):
        element['url'] = os.path.join('/artifacts', item_path).replace('\\', '/')
        if io_paths:
            output_path = os.path.join(io_paths['output'], 'artifacts', item_path)
            input_path = os.path.join(io_paths['input'], item_path)
            utils.create_dir(os.path.dirname(output_path))
            if not os.path.exists(output_path) or not filecmp.cmp(input_path, output_path):
                shutil.copy(input_path, output_path)


def convert_element(element, io_paths=None):
    """ Converts a dict with a name url and type to markdown """
    prepare_locally_stored_files(element, io_paths)
    if element['type'].lower() == 'image':
        base_text = '\n![{0}]({1})\n'
    else:
        base_text = '\n[{0}]({1})\n'
    return base_text.format(element['name'], element['url'])


def generate_text_narative(narative):
    """ Checks if the narrative is in dict format or in string format.
    If the narrative is in dict format the script converts it to to a
    string """
    text = ''
    if isinstance(narative, dict):
        for key in sorted(narative):
            text += '{0}. {1} \n '.format(key, narative[key])
    else:
        text = narative + '  \n'
    return text


def build_standards_summary(summaries, output_path):
    """ Construct the summary for the standards """
    main_summary = "## Standards  \n\n"
    for standard_key in natural_sort(summaries['standards']):
        for family_key in natural_sort(summaries['standards'][standard_key]):
            section_summary = '# {0}  \n'.format(family_key)
            main_summary += '* [{0} - {1}](content/{1}.md)\n'.format(standard_key, family_key)
            for control_key in natural_sort(summaries['standards'][standard_key][family_key]):
                control = summaries['standards'][standard_key][family_key][control_key]
                main_summary += '\t* [{0} - {1}](content/{2}.md)\n'.format(
                    control['family'],
                    control['control_name'],
                    control['slug']
                )
                section_summary += '* [{0} - {1}]({2}.md)\n'.format(
                    control['control'],
                    control['control_name'],
                    control['slug']
                )
            write_markdown(output_path, 'content/' + family_key + '.md', section_summary)
    return main_summary


def build_components_summary(summaries, output_path):
    """ Construct the summary for the components """
    main_summary = '\n## Systems  \n\n'
    for system_key in sorted(summaries['components']):
        main_summary += '* [{0}](content/{1}.md)\n'.format(system_key, system_key)
        section_summary = '# {0}  \n###Components  \n'.format(system_key)
        for component_key in sorted(summaries['components'][system_key]):
            component = summaries['components'][system_key][component_key]
            # Add the components url to main summary
            main_summary += '\t* [{0}](content/{1}.md)\n'.format(
                component['component_key'],
                component['slug']
            )
            # Add the components url to section summary
            section_summary += '* [{0}]({1}.md)\n'.format(
                component['component_key'],
                component['slug']
            )
        write_markdown(output_path, 'content/' + system_key + '.md', section_summary)
    return main_summary


def concat_markdowns(markdown_path, output_path):
    """ Add markdown content files to the gitbook directory and make the summary
    file the base summary string in order to join the markdown summary with
    the gitbook generated in this file. """
    for filename in glob.iglob(os.path.join(markdown_path, "*", "*")):
        # Get the output file path and create the directory before copying
        output_filepath = os.path.join(
            output_path, filename.replace(os.path.join(markdown_path, ''), '')
        )
        ouput_dir = os.path.dirname(output_filepath)
        utils.create_dir(ouput_dir)
        shutil.copy(filename, output_filepath)
    summary_path = os.path.join(markdown_path, 'SUMMARY.md')
    with open(summary_path, 'r') as summary_file:
        main_summary = summary_file.read()
    return main_summary


def build_summary(summaries, output_path, markdown_path):
    """ Construct a gitbook summary for the controls """
    if markdown_path and os.path.exists(markdown_path):
        main_summary = concat_markdowns(markdown_path, output_path)
        # load the main markdown
    else:
        main_summary = "# Summary  \n\n"
    main_summary += build_standards_summary(summaries, output_path)
    main_summary += build_components_summary(summaries, output_path)
    write_markdown(output_path, 'SUMMARY.md', main_summary)
    write_markdown(output_path, 'README.md', main_summary)


def document_cert_page(certification, standard_key, control_key):
    """ Create a new page dict. This item is a dictionary that
    contains the standard and control keys, a slug of the combined key, and the
    name of the control"""
    meta_data = certification['standards'][standard_key][control_key]['meta']

    slug = slugify('{0}-{1}'.format(standard_key, control_key))
    return {
        'control': control_key,
        'standard': standard_key,
        'family': meta_data.get('family'),
        'control_name': meta_data.get('name'),
        'slug': slug
    }


def document_component_page(certification, system_key, component_key):
    """ Create a new page dict. This item is a dictionary that
    contains the standard and control keys, a slug of the combined key, and the
    name of the control"""
    component = certification['components'][system_key]['components'][component_key]
    slug = slugify('{0}-{1}'.format(system_key, component_key))
    return {
        'system_key': system_key,
        'component_key': component_key,
        'component_name': component['name'],
        'slug': slug
    }


def fetch_component(reference, certification):
    """ Fetches a specific component from the certification dict,
    this component will be used to extract the component name and it's verifications
    when they are referenced """
    return certification['components'][reference['system']]['components'][reference['component']]


def fetch_verification(verification_ref, certification):
    """ Get the verfication component """
    component = fetch_component(verification_ref, certification)['verifications']
    return component[verification_ref['verification']]


def org_by_system_component(justifications):
    """ Organizes list of justifications in a dictionary of systems and
    components """
    justifications_dict = {}
    for justification in justifications:
        system_key = justification['system']
        component_key = justification['component']
        if system_key not in justifications_dict:
            justifications_dict[system_key] = {}
        justifications_dict[system_key][component_key] = justification
    return justifications_dict


def build_control_text(control, certification):
    """ Generate the markdown text from each `justification` """
    text = ''
    # Order the justifications by system and then component
    justifications_dict = org_by_system_component(control.get('justifications', []))
    for system_key in sorted(justifications_dict):
        text += '\n## {0}\n'.format(system_key)
        for component_key in sorted(justifications_dict[system_key]):
            justification = justifications_dict[system_key][component_key]
            component = fetch_component(justification, certification)
            text += '\n## {0}\n'.format(component.get('name'))
            text += generate_text_narative(justification.get('narrative'))
            verifications = justification.get('references')
            if verifications:
                for verification_ref in verifications:
                    text += convert_element(
                        fetch_verification(verification_ref, certification)
                    )
    return text


def build_component_text(component, io_paths):
    """ Create markdown output for component text """
    text = '\n### References  \n'
    references = component.get('references', [])
    if references:
        for reference in sorted(references, key=lambda k: k['name']):
            text += convert_element(reference, io_paths)
        text += '\n### Verifications  \n'
    verifications = component.get('verifications', [])
    if verifications:
        for verification_key in sorted(verifications):
            text += convert_element(component['verifications'][verification_key], io_paths)
    return text


def build_cert_page(page_dict, certification, output_path):
    """ Write a page for the gitbook """
    text = '# {0}'.format(page_dict['control_name'])
    control = certification['standards'][page_dict['standard']][page_dict['control']]
    text += build_control_text(control, certification)
    file_name = 'content/' + page_dict['slug'] + '.md'
    write_markdown(output_path, file_name, text)


def build_component_page(page_dict, certification, io_paths):
    """ Write a page for the gitbook """
    text = '# {0}'.format(page_dict['component_name'])
    component = certification['components'][page_dict['system_key']]['components'][page_dict['component_key']]
    text += build_component_text(component, io_paths)
    file_name = 'content/' + page_dict['slug'] + '.md'
    write_markdown(io_paths['output'], file_name, text)


def natural_sort(elements):
    """ Natural sorting algorithms for stings with text and numbers reference:
    stackoverflow.com/questions/4836710/
    """
    convert = lambda text: int(text) if text.isdigit() else text.lower()
    alphanum_key = lambda key: [convert(c) for c in re.split('([0-9]+)', str(key))]
    return sorted(elements, key=alphanum_key)


def build_standards_documentation(certification, output_path):
    """ Create the documentation for standards """
    summary = {}
    for standard_key in certification['standards']:
        summary[standard_key] = {}
        for control_key in certification['standards'][standard_key]:
            if 'justifications' in certification['standards'][standard_key][control_key]:
                page_dict = document_cert_page(certification, standard_key, control_key)
                build_cert_page(page_dict, certification, output_path)
                if page_dict['family'] not in summary[standard_key]:
                    summary[standard_key][page_dict['family']] = {}
                summary[standard_key][page_dict['family']][control_key] = page_dict
    return summary


def build_components_documentation(certification, io_paths):
    """ Create the documentation for the components """
    summary = {}
    for system_key in sorted(certification['components']):
        summary[system_key] = {}
        for component_key in sorted(certification['components'][system_key]['components']):
            page_dict = document_component_page(certification, system_key, component_key)
            build_component_page(page_dict, certification, io_paths)
            summary[system_key][component_key] = page_dict
    return summary


def create_gitbook_documentation(certification_path, output_path, markdown_path=None):
    """ Convert certification to pages format """
    summaries = {}
    io_paths = {
        'output': output_path,
        'input': os.path.dirname(certification_path)
    }
    certification = utils.yaml_loader(certification_path)
    summaries['components'] = build_components_documentation(certification, io_paths)
    summaries['standards'] = build_standards_documentation(certification, output_path)
    build_summary(summaries, output_path, markdown_path)
    return output_path
