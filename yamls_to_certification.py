import glob
import os

from yaml import dump, load


def write_yaml_data(component_data, filename):
    """ Write component data to a yaml file """
    with open(filename, 'w') as yaml_file:
        yaml_file.write(dump(component_data, default_flow_style=False))


def component_dict_gen(components_path):
    """ Open yaml files and creates a dict """
    for filename in glob.iglob(components_path):
        with open(filename, 'r') as yaml_file:
            yield load(yaml_file)


def check_and_add_to_dict(new_dict, old_dict, key):
    """ Check if the dict has a key, othewise issues a warning """
    if key in old_dict:
        new_dict[key] = old_dict.get(key)
    else:
        print(old_dict.get('name'), "is missing data:", key)


def prepare_component(component_dict):
    """ Creates a deep copy of the comonent dict, strips away uneeded data """
    new_component_dict = dict()
    for key in ['name', 'references', 'governors']:
        check_and_add_to_dict(
            new_dict=new_component_dict, old_dict=component_dict, key=key)
    return new_component_dict


def merge_compontents_standards(component_dict, standards_dict):
    """ Adds each component dictionary to a dictionary organized by
    by the control it satisfies deep copies are used because a component
    can meet multiple standards"""
    for control in component_dict['satisfies']:
        if not standards_dict.get(control):
            standards_dict[control] = list()
        preped_component = prepare_component(component_dict)
        preped_component['narative'] = component_dict['satisfies'][control]
        standards_dict[control].append(preped_component)


def create_standards_dict(components_path):
    """ Open component files and organize them by the standards/controls
    each satisfies """
    standards_dict = dict()
    for component_dict in component_dict_gen(components_path):
        merge_compontents_standards(
            component_dict=component_dict, standards_dict=standards_dict)
    return standards_dict


def certifications_gen(certifications_path):
    """ Open the certifications files and yields a dict """
    for filename in glob.iglob(certifications_path):
        with open(filename, 'r') as yaml_file:
            yield load(yaml_file)


def merge_standards_certifications(certifications_path, controls_path):
    """ Combine the components organized by controls and each certification
    to create a dict that shows how each control in a certification has
    been met"""
    standards = create_standards_dict(controls_path)
    for certification in certifications_gen(certifications_path):
        for standard in certification['standards']:
            for control in certification['standards'][standard]:
                if not certification['standards'][standard][control]:
                    certification['standards'][standard][control] = list()
                if standards.get(control):
                    certification['standards'][standard][control].append(
                        standards[control])
                else:
                    print(control, "missing")
        yield certification['name'], certification


def create_certifications(certifications_path, components_path, output_path):
    """ Generate certification yamls from data """
    certifications = merge_standards_certifications(
        certifications_path, components_path)
    for name, certification in certifications:
        filename = os.path.join(output_path, name + '.yaml')
        write_yaml_data(component_data=certification, filename=filename)


if __name__ == "__main__":
    create_certifications(
        certifications_path='data/certifications/*.yaml',
        components_path='data/components/*/*.yaml',
        output_path='completed_certifications'
    )
