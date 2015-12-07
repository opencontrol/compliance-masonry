""" Module for generating empty data templates """

import os
import shutil

from slugify import slugify
from src import utils


def get_template_dir():
    """ Finds the directory that the project templates are in """
    module_dir = os.path.dirname(__file__)
    return os.path.join(module_dir, 'templates/')


def get_file_path(output_dir, system_key, component_key=None):
    """ Creates the path for the directory that will contain the component
    if it doesn't exist and returns the file path of component yaml"""
    filepath = os.path.join(output_dir, system_key)
    filename = 'system.yaml'
    if component_key:
        filepath = os.path.join(filepath, component_key)
        filename = 'component.yaml'
    utils.create_dir(filepath)
    return os.path.join(filepath, filename)


def create_data_dict(system_key, component_key):
    """ Generates a starting template for the component dictionary """
    if component_key:
        return {
            'name': component_key,
            'references': [
                {
                    'name': 'Reference Name',
                    'path': 'http://dummyimage.com/600x400',
                    'type': 'Image'
                }
            ],
            'verifications': {
                'Verification_ID': {
                    'name': 'Verification Name',
                    'path': 'http://dummyimage.com/600x400',
                    'type': 'Image'
                }
            },
            'satisfies': {},
            'documentation_complete': False
        }
    else:
        return {'name': system_key}


def create_new_data_yaml(output_dir, system_key, component_key=None):
    """ Create new component yaml """
    system_key = slugify(system_key)
    if component_key:
        component_key = slugify(component_key)
    file_path = get_file_path(output_dir, system_key, component_key)
    data_dict = create_data_dict(system_key, component_key)
    utils.yaml_writer(data_dict, file_path)
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
