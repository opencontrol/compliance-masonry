""" This script converts the components and standards yamls into
certifications """

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


def create_standards_dic(standards_path):
    """ Create a standards dictionary for later merging with the
    certifications documentation """
    return {
        standard['name']: standard for standard in yaml_loader(standards_path)
    }


def check_and_add_key(new_dict, old_dict, key):
    """ Check if the dict has a key, otherwise issues a warning """
    if key in old_dict:
        new_dict[key] = old_dict.get(key)
    else:
        logging.warning("Component `%s` is missing `%s` data", old_dict.get('name'), key)


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
    for control in component_dict['satisfies']:
        if not bystandards_dict.get(control):
            bystandards_dict[control] = list()
        preped_component = prepare_component(component_dict)
        preped_component['narative'] = component_dict['satisfies'][control]
        bystandards_dict[control].append(preped_component)


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
    control_justification = components.get(control)
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
        for standard in certification['standards']:
            for control in certification['standards'][standard]:
                # Create a reference to the certification control
                certification['standards'][standard][control] = dict()
                merge_components(certification, components, standard, control)
                merge_standard(certification, standards, standard, control)
        yield certification['name'], certification


def create_certifications(
        certifications_path, components_path, output_path, standards_path):
    """ Generate certification yamls from data """
    standards = create_standards_dic(standards_path)
    components = create_bystandards_dict(components_path)
    certifications = build_certifications(
        certifications_path, components, standards)
    for name, certification in certifications:
        filename = os.path.join(output_path, name + '.yaml')
        yaml_writer(component_data=certification, filename=filename)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    create_certifications(
        certifications_path='data/certifications/*.yaml',
        components_path='data/components/*/*.yaml',
        standards_path='data/standards/*.yaml',
        output_path='completed_certifications'
    )

