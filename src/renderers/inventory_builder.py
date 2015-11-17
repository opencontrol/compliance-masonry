""" Script for creating an inventory yaml """

import os

from src import utils


def prepare_cert_path(certification, certification_dir):
    """ Prepare the path for a specific certification """
    if not certification_dir:
        certification_dir = 'exports/certifications/'
    return os.path.join(certification_dir, '{0}.yaml'.format(certification))


def analyze_attribute(attribute):
    """ Check how many elements an attribute has otherwise return "Missing """
    if attribute:
        return len(attribute)
    return "Missing"


def analyze_component(component):
    """ Analyze a component to find gaps in governors and references """
    return {
        'references': analyze_attribute(component.get('references')),
        'governors': analyze_attribute(component.get('governors')),
        'documentation_completed': component.get('documentation_complete'),
    }


def catalog_control(inventory, control, standard_key, control_key):
    """ Adds all the components in the control into the inventory
    while determing the gaps """
    if 'justifications' in control:
        for component in control['justifications']:
            system = component.get('system', 'No System')
            name = component.get('name', 'No Name')
            # Catalog component in certification inventory
            if system not in inventory[standard_key][control_key]:
                inventory[standard_key][control_key][system] = []
            inventory[standard_key][control_key][system].append(name)
            # Catalog component in component inventory
            analysis = analyze_component(component)
            if system not in inventory['components']:
                inventory['components'][system] = {}
            inventory['components'][system][name] = analysis
    else:
        inventory[standard_key][control_key] = "Missing Justifications"


def build_inventory(certification_path):
    """ Create an inventory of components for a specific certification """
    certification = utils.yaml_loader(certification_path)
    inventory = {
        'certification': certification.get('name'),
        'components': {}
    }
    for standard_key in certification['standards']:
        inventory[standard_key] = {}
        for control_key in certification['standards'][standard_key]:
            inventory[standard_key][control_key] = {}
            control = certification['standards'][standard_key][control_key]
            catalog_control(inventory, control, standard_key, control_key)
    return inventory


def create_inventory(certification_path, output_path):
    """ Creates an inventory yaml """
    inventory = build_inventory(certification_path)
    inventory_path = os.path.join(
        output_path,
        inventory.get('certification') + '.yaml'
    )
    utils.yaml_writer(inventory, inventory_path)
    return inventory_path
