""" Module for generating empty data templates """

import os
import shutil
import yaml

from glob import iglob
from slugify import slugify

def get_template_dir():
    """ Finds the directory that the project templates are in """
    module_dir = os.path.dirname(__file__)
    return os.path.join(module_dir, 'templates/')


def create_output_dirs(output_path):
    """ Create the directories on the output path if they don't exist """
    if output_path and not os.path.exists(output_path):
        os.makedirs(output_path)


def get_file_path(system, name, output_dir):
    """ Creates the path for the directory that will contain the component
    if it doesn't exist and returns the file path of component yaml"""
    if not output_dir:
        output_dir = 'data/components'
    output_path = os.path.join(output_dir, system)
    create_output_dirs(output_path)
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


def init_project(output_dir):
    """ Create a new control masonry project template """
    if not output_dir:
        output_dir = 'data'
    output_container, _ = os.path.split(output_dir)
    create_output_dirs(output_container)
    template_dir = get_template_dir()
    copy_to_path = os.path.join(os.getcwd(), output_dir)
    shutil.copytree(template_dir, copy_to_path)
    return output_dir
