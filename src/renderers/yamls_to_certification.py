""" This script converts the components and standards yamls into
certifications """

import copy
import glob
import logging
import os

from yaml import dump, load


def yaml_writer(component_data, filename):
    """ Write component data to a yaml file """
    with open(filename, 'w') as yaml_file:
        yaml_file.write(dump(component_data, default_flow_style=False))


def yaml_loader(glob_path):
    """ Creates a generator for loading yaml files into dicts """
    for filename in glob.iglob(glob_path):
        with open(filename, 'r') as yaml_file:
            yield load(yaml_file)


def prepare_data_paths(data_dir):
    """ Create glob paths for data directory includes glob path for
    certifications, components, and standards """
    if not data_dir:
        data_dir = 'data'
    certifications_path = os.path.join(data_dir, 'certifications/*.yaml')
    components_path = os.path.join(data_dir, 'components/*/*.yaml')
    standards_path = os.path.join(data_dir, 'standards/*.yaml')
    return certifications_path, components_path, standards_path


def prepare_cert_output_path(output_dir):
    """ Creates a path for the certifications exports directory """
    if not output_dir:
        output_dir = 'exports'
    output_path = os.path.join(output_dir, 'certifications')
    if not os.path.exists(output_path):
        os.makedirs(output_path)
    return output_path


def create_standards_dic(standards_path):
    """ Create a standards dictionary for later merging with the
    certifications documentation """
    return {
        standard['name']: standard for standard in yaml_loader(standards_path)
    }


def check_and_add_key(new_dict, old_dict, key):
    """ Check if the dict has a key, otherwise issues a warning """
    if key in old_dict:
        new_dict[key] = copy.deepcopy(old_dict.get(key))
    else:
        logging.warning(
            "Component `%s` is missing `%s` data", old_dict.get('name'), key)


def prepare_component(component_dict):
    """ Creates a deep copy of the component dict, but only keeps the name,
    references, and governors data"""
    new_component_dict = dict()
    for key in ['name', 'references', 'governors']:
        check_and_add_key(
            new_dict=new_component_dict, old_dict=component_dict, key=key)
    return new_component_dict


def convert_to_bystandards(component_dict, bystandards_dict):
    """ Adds each component dictionary to a dictionary organized by
    by the control it satisfies deep copies are used because a component
    can meet multiple standards"""
    for standard in component_dict['satisfies']:
        if not bystandards_dict.get(standard):
            bystandards_dict[standard] = dict()
        for control in component_dict['satisfies'][standard]:
            if not bystandards_dict[standard].get(control):
                bystandards_dict[standard][control] = list()
            preped_component = prepare_component(component_dict)
            preped_component['narative'] = component_dict['satisfies'][standard][control]
            bystandards_dict[standard][control].append(preped_component)


def create_bystandards_dict(components_path):
    """ Open component files and organize them by the standards/controls
    each satisfies """
    bystandards_dict = dict()
    for component_dict in yaml_loader(components_path):
        convert_to_bystandards(
            component_dict=component_dict, bystandards_dict=bystandards_dict)
    return bystandards_dict


def merge_components(certification, components, standard, control):
    """ Adds the components to the certification control and warns
    user if control has no documentation """
    control_justification = components.get(standard).get(control)
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


def build_certifications(certifications_path, components, standards):
    """ Merges the components and standards data with the certification
    data """
    for certification in yaml_loader(certifications_path):
        for standard in sorted(certification['standards']):
            for control in sorted(certification['standards'][standard]):
                # Create a reference to the certification control
                certification['standards'][standard][control] = dict()
                merge_components(certification, components, standard, control)
                merge_standard(certification, standards, standard, control)
        yield certification['name'], certification


def create_yaml_certifications(data_dir, output_dir):
    """ Generate certification yamls from data """
    certifications_path, components_path, standards_path = prepare_data_paths(data_dir)
    output_path = prepare_cert_output_path(output_dir)
    standards = create_standards_dic(standards_path)
    components = create_bystandards_dict(components_path)
    certifications = build_certifications(
        certifications_path, components, standards)
    for name, certification in certifications:
        filename = os.path.join(output_path, name + '.yaml')
        yaml_writer(component_data=certification, filename=filename)
    return output_path
