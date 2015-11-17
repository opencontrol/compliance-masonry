""" This script converts the components and standards yamls into
certifications """

import copy
import logging
import os

from src import utils


def prepare_data_paths(certification, data_dir):
    """ Create the default glob paths for certifications, components, and standards """
    certifications_path = os.path.join(
        data_dir, 'certifications/{0}.yaml'.format(certification)
    )
    components_path = os.path.join(data_dir, 'components/*/*.yaml')
    standards_path = os.path.join(data_dir, 'standards/*.yaml')
    return certifications_path, components_path, standards_path


def prepare_output_path(output_path):
    """ Creates a path for the certifications exports directory """

    return output_path


def create_standards_dic(standards_path):
    """ Create a standards dictionary for later merging with the
    certifications documentation """
    return {
        standard['name']: standard for standard in utils.yaml_gen_loader(standards_path)
    }


def copy_key(new_dict, old_dict, key):
    """ Copy value of key if it exists from old dict and add it to a
    new dictionary  """
    if key in old_dict:
        new_dict[key] = copy.deepcopy(old_dict.get(key))
    else:
        logging.warning(
            "Component `%s` is missing `%s` data", old_dict.get("name"), key)


def extract_components(main_dict, component_dict):
    """ Extracts the fields that pertain to the individual components and
    adds to an existing dict (main_dict) """
    if not component_dict['system'] in main_dict:
        main_dict[component_dict['system']] = {}
    if not component_dict['component'] in main_dict[component_dict['system']]:
        main_dict[component_dict['system']][component_dict['component']] = {
            'name': component_dict['name'],
            'documentation_complete': component_dict['documentation_complete'],
            'references':  component_dict['references'],
            'verifications': component_dict['verifications'],
        }


def prepare_references(component_dict, standard, control):
    """ Creates a dict containing references for the specific
    components standard-control's justification references which don't have
    a specific component and system are assigned the component and system of the
    system and component they derived from """
    references = component_dict['satisfies'][standard][control].get('references')
    if references:
        for idx, reference in enumerate(references):
            if 'component' not in reference:
                reference['component'] = component_dict['component']
            if 'system' not in reference:
                reference['system'] = component_dict['system']
    return references


def extract_standards(main_dict, component_dict):
    """ Extracts the fields that pertain to the st components and adds
    them to an existing dict (main_dict) organized by by the standard and
    control each satisfies. Deep copies are used because a component can meet
    multiple standards """
    for standard in component_dict['satisfies']:
        if not main_dict.get(standard):
            main_dict[standard] = dict()
        for control in component_dict['satisfies'][standard]:
            if not main_dict[standard].get(control):
                main_dict[standard][control] = list()
            main_dict[standard][control].append({
                'system': component_dict['system'],
                'component': component_dict['component'],
                'narrative': component_dict['satisfies'][standard][control]['narrative'],
                'implementation_status': component_dict['satisfies'][standard][control]['implementation_status'],
                'references': prepare_references(component_dict, standard, control)
            })


def parse_components(components_path):
    """ Open component files and organize them into two dicts on with
    components data and the other with standards data """
    components_dict = dict()
    bystandards_dict = dict()
    for component_dict in utils.components_loader(components_path):
        extract_components(main_dict=components_dict, component_dict=component_dict)
        extract_standards(main_dict=bystandards_dict, component_dict=component_dict)
    return components_dict, bystandards_dict


def merge_components(certification, components, standard, control):
    """ Adds the components to the certification control and warns
    user if control has no documentation """
    control_justification = components.get(standard, {}).get(control)
    if control_justification:
        certification['standards'][standard][control]['justifications'] = \
            control_justification
    else:
        logging.warning(
            "`%s` certification is missing `%s %s` justifications",
            certification.get('name'), standard, control
        )


def merge_standard(certification, standards, standard, control):
    """ Adds information data to the certification control and warns
    user if control has no information data """
    control_info = standards[standard].get(control)
    if control_info:
        certification['standards'][standard][control]['meta'] = control_info
    else:
        logging.warning(
            "`%s` certification is missing `%s %s` info",
            certification.get('name'), standard, control
        )


def build_certification(certifications_path, bystandards_dict, standards):
    """ Merges the components and standards data with the certification
    data """
    certification = utils.yaml_loader(certifications_path)
    for standard in certification['standards']:
        for control in sorted(certification['standards'][standard]):
            # Create a reference to the certification control
            certification['standards'][standard][control] = dict()
            merge_components(certification, bystandards_dict, standard, control)
            merge_standard(certification, standards, standard, control)
    return certification['name'], certification


def create_yaml_certifications(certification, data_dir, output_dir):
    """ Generate certification yamls from data """
    certifications_path, components_path, standards_path = prepare_data_paths(certification, data_dir)
    standards = create_standards_dic(standards_path)
    components_dict, bystandards_dict = parse_components(components_path)
    name, certification = build_certification(
        certifications_path, bystandards_dict, standards
    )
    certification['components'] = components_dict
    filename = os.path.join(output_dir, name + '.yaml')
    utils.yaml_writer(component_data=certification, filename=filename)
    return filename
