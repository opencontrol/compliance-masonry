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
        yaml_file.write(yaml.dump(component_data, default_flow_style=False, indent=2))


def yaml_loader(filename):
    """ Load a yaml file """
    with open(filename) as yaml_file:
        return yaml.load(yaml_file)


def yaml_gen_loader(glob_path):
    """ Creates a generator for loading yaml files into dicts """
    for filename in glob.iglob(glob_path):
        with open(filename, 'r') as yaml_file:
            yield yaml.load(yaml_file)


def components_loader(glob_path):
    """ Generator for loading component yaml files. Attaches a system and
    components keys to dictionary """
    for filename in glob.iglob(glob_path):
        system_path, component = os.path.split(filename)
        component = component.replace('.yaml', '')
        _, system = os.path.split(system_path)
        with open(filename, 'r') as yaml_file:
            data = yaml.load(yaml_file)
            data['component'] = component
            data['system'] = system
            yield data


def fetch_available_certifications(data_dir):
    """ Return a list of avaiable certifications """
    return [
        os.path.split(path.replace('.yaml', ''))[-1] for path in glob.iglob(os.path.join(data_dir, '*.yaml'))
    ]


def check_certifications(certification, data_dir):
    """ Checks if certification is present, if the certification is present
    returns the certification path otherwise returns a list of the certifications
    that are avaiable """
    certification_path = os.path.join(data_dir, certification + ".yaml")
    if os.path.exists(certification_path):
        return certification_path, None
    return None, fetch_available_certifications(data_dir)
