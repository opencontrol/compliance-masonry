""" Module for generating empty data templates """

import os
import yaml

from slugify import slugify


def get_file_path(system, name, output_dir):
    """ Creates the path for the directory that will contain the component
    if it doesn't exist and returns the file path of component yaml"""
    if not output_dir:
        output_dir = 'data/components'
    output_path = os.path.join(output_dir, system)
    if not os.path.exists(output_path):
        os.makedirs(output_path)
    return os.path.join(output_path, '{0}.yaml'.format(slugify(name)))


def create_component_dict(system, name):
    """ Generates a starting template for the component dictionary """
    return {
        'system': system,
        'name': name,
        'references': [{'name': 'Reference Name'}, {'url': 'Refernce URL'}],
        'governors': [{'name': 'Governor Name'}, {'url': 'Governor URL'}],
        'satisfies': {}
    }


def create_new_component_yaml(system, name, output_dir):
    """ Create new component yaml """
    file_path = get_file_path(system, name, output_dir)
    component_dict = create_component_dict(system, name)
    with open(file_path, 'w') as yaml_file:
        yaml_file.write(yaml.dump(component_dict, default_flow_style=False))
    return file_path
