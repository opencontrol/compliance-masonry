import os
import re

from slugify import slugify
from src import utils


def write_markdown(output_path, filename, text):
    """ Write text to a markdown file """
    filename = os.path.join(output_path, filename)
    with open(filename, 'w') as md_file:
        md_file.write(text)


def convert_name_url(references):
    """ Converts references data to markdown url bullet point. """
    text = ''
    for reference in references:
        text += '\n* [{0}]({1})\n'.format(
            reference['name'], reference['url'])
    return text


def generate_text_narative(narative):
    """ Checks if the narrative is in dict format or in string format.
    If the narrative is in dict format the script converts it to to a
    string """
    text = ''
    if type(narative) == dict:
        for key in sorted(narative):
            text += '{0}. {1} \n '.format(key, narative[key])
    else:
        text = narative
    return text


def build_summary(summary, output_path):
    """ Construct a gitbook summary for the controls """
    main_summary = "# Summary\n\n"
    last_family = ''
    section_summary = ''
    for control in summary:
        new_family = control['family']
        if last_family != new_family:
            control_family_name = control['control_name']
            main_summary += '* [{0} - {1} - {2}](content/{1}.md)\n'.format(
                control['standard'],
                control['family'],
                control_family_name,
            )
            # Write the section summary
            if last_family:
                write_markdown(output_path, 'content/' + last_family + '.md', section_summary)
            # Start a new section summary
            section_summary = '# {0} - {1}\n\n'.format(new_family, control_family_name)
        # Add the control url to main summary
        main_summary += '\t* [{0} - {1}](content/{2}.md)\n'.format(
            control['control'],
            control['control_name'],
            control['slug']
        )
        # Add the control url to section summary
        section_summary += '* [{0} - {1}]({2}.md)\n'.format(
            control['control'],
            control['control_name'],
            control['slug']
        )
        last_family = new_family
    # Export the last family
    write_markdown(output_path, 'content/' + last_family + '.md', section_summary)
    write_markdown(output_path, 'SUMMARY.md', main_summary)
    write_markdown(output_path, 'README.md', main_summary)


def document_page(summary, certification, standard_key, control_key):
    """ Create a new page dict. This item is a dictionary that
    contains the standard and control keys, a slug of the combined key, and the
    name of the control"""
    control_name = certification['standards'][standard_key][control_key]['meta']['name']
    slug = slugify('{0}-{1}'.format(standard_key, control_key))
    return {
        'control': control_key,
        'standard': standard_key,
        'family': control_key.split('-')[0],
        'control_name': control_name,
        'slug': slug
    }


def create_content(control):
    """ Generate the markdown text from each `justification` """
    text = ''
    for justification in control.get('justifications', []):
        text += '\n## {0}\n'.format(justification.get('name'))
        text += generate_text_narative(justification.get('narative'))
        references = justification.get('references')
        if references:
            text += '\n### References\n'
            text += convert_name_url(references)
        governors = justification.get('governors')
        if governors:
            text += '\n### Governors\n'
            text += convert_name_url(governors)
        text += "\n--------\n"
    return text


def build_page(page_dict, certification, output_path):
    """ Write a page for the gitbook """
    text = '# {0}'.format(page_dict['control_name'])
    control = certification['standards'][page_dict['standard']][page_dict['control']]
    text += create_content(control)
    file_name = 'content/' + page_dict['slug'] + '.md'
    write_markdown(output_path, file_name, text)


def natural_sort(elements):
    """ Natural sorting algorithms for stings with text and numbers reference:
    stackoverflow.com/questions/4836710/
    """
    convert = lambda text: int(text) if text.isdigit() else text.lower()
    alphanum_key = lambda key: [convert(c) for c in re.split('([0-9]+)', key)]
    return sorted(elements, key=alphanum_key)


def create_gitbook_documentation(certification_path, output_path):
    """ Convert certification to pages format """
    summary = []
    certification = utils.yaml_loader(certification_path)
    for standard_key in natural_sort(certification['standards']):
        for control_key in natural_sort(certification['standards'][standard_key]):
            if 'justifications' in certification['standards'][standard_key][control_key]:
                page_dict = document_page(summary, certification, standard_key, control_key)
                summary.append(page_dict)
                build_page(page_dict, certification, output_path)
    build_summary(summary, output_path)
    return output_path
