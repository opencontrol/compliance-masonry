""" Script for creating an inventory yaml """

import os

from src import utils


def prepare_cert_path(certification, certification_dir):
    """ Prepare the path for a specific certification """
    if not certification_dir:
        certification_dir = 'exports/certifications/'
    return os.path.join(certification_dir, '{0}.yaml'.format(certification))


def analyze_attribute(attribute):
    """ Check how many elements an attribute has otherwise if it's a list
    if it's not a list return that it's present otherwise return "Missing """
    if isinstance(attribute, list) or isinstance(attribute, dict):
        return len(attribute)
    elif attribute:
        return "Present"
    return "Missing"


def analyze_component(component):
    """ Analyze a component to find gaps in governors and references """
    return {
        'references': analyze_attribute(component.get('references')),
        'verifications': analyze_attribute(component.get('verifications')),
        'documentation_completed': component.get('documentation_complete'),
    }


def catalog_control(inventory, control, standard_key, control_key):
    """ Adds all the components in the control into the inventory
    while determing the gaps """
    if 'justifications' in control:
        for component in control['justifications']:
            system_key = component.get('system', 'No System')
            component_key = component.get('component', 'No Name')
            # Catalog component in certification inventory
            if system_key not in inventory[standard_key][control_key]:
                inventory[standard_key][control_key][system_key] = {}
            if component_key not in inventory[standard_key][control_key][system_key]:
                inventory[standard_key][control_key][system_key][component_key] = {}
            inventory[standard_key][control_key][system_key][component_key] = {
                'implementation_status': component.get('implementation_status', 'Missing'),
                'narrative': analyze_attribute(component.get('narrative')),
                'references': analyze_attribute(component.get('references'))
            }
    else:
        inventory[standard_key][control_key] = "Missing Justifications"


def catalog_component(component, inventory, system_key, component_key):
    """ Summarizes the data in the components dict """
    inventory['components'][system_key][component_key] = analyze_component(component)


def inventory_standards(certification, inventory):
    """ Populate the inventory for standards """
    for standard_key in certification['standards']:
        inventory[standard_key] = {}
        for control_key in certification['standards'][standard_key]:
            inventory[standard_key][control_key] = {}
            control = certification['standards'][standard_key][control_key]
            catalog_control(inventory, control, standard_key, control_key)


def inventory_components(certification, inventory):
    """ Populate the inventory for components """
    for system_key in certification['components']:
        if system_key not in inventory['components']:
            inventory['components'][system_key] = {}
        for component_key in certification['components'][system_key]:
            catalog_component(
                certification['components'][system_key][component_key],
                inventory,
                system_key,
                component_key
            )


def build_inventory(certification_path):
    """ Create an inventory of components for a specific certification """
    certification = utils.yaml_loader(certification_path)
    inventory = {
        'certification': certification.get('name'),
        'components': {}
    }
    inventory_standards(certification, inventory)
    inventory_components(certification, inventory)
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
