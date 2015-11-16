import os
import glob
import yaml


def create_dir(output_path):
    """ Given output path create directories """
    if not os.path.exists(output_path):
        os.makedirs(output_path)


def yaml_writer(component_data, filename):
    """ Write component data to a yaml file """
    with open(filename, 'w') as yaml_file:
        yaml_file.write(yaml.dump(component_data, default_flow_style=False))


def yaml_loader(filename):
    """ Load a yaml file """
    with open(filename) as yaml_file:
        return yaml.load(yaml_file)


def yaml_gen_loader(glob_path):
    """ Creates a generator for loading yaml files into dicts """
    for filename in glob.iglob(glob_path):
        with open(filename, 'r') as yaml_file:
            yield yaml.load(yaml_file)
