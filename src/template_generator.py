""" Module for generating empty data templates """

import os
import shutil

from slugify import slugify
from src import utils


def get_template_dir():
    """ Finds the directory that the project templates are in """
    module_dir = os.path.dirname(__file__)
    return os.path.join(module_dir, 'templates/')


def get_file_path(system, name, output_dir):
    """ Creates the path for the directory that will contain the component
    if it doesn't exist and returns the file path of component yaml"""
    output_path = os.path.join(output_dir, system)
    utils.create_dir(output_path)
    return os.path.join(output_path, '{0}.yaml'.format(slugify(name)))


def create_component_dict(system, component):
    """ Generates a starting template for the component dictionary """
    return {
        'name': component,
        'references': [
            {'name': 'Reference Name', 'url': 'Refernce URL',  'type': 'Image'}
        ],
        'verifications': [
            {'name': 'Verification Name', 'url': 'Verification URL',  'type': 'Image'}
        ],
        'satisfies': {},
        'documentation_complete': False
    }


def create_new_component_yaml(system, component, output_dir):
    """ Create new component yaml """
    file_path = get_file_path(system, component, output_dir)
    component_dict = create_component_dict(system, component)
    utils.yaml_writer(component_dict, file_path)
    return file_path


def init_project(output_dir):
    """ Create a new control masonry project template """
    if not output_dir:
        output_dir = 'data'
    output_container, _ = os.path.split(output_dir)
    utils.create_dir(output_container)
    template_dir = get_template_dir()
    copy_to_path = os.path.join(os.getcwd(), output_dir)
    shutil.copytree(template_dir, copy_to_path)
    return output_dir
