""" This script converts the Control-masonry excel file to
a series of yamls """

import os

from openpyxl import load_workbook
from yaml import dump


class ComponentMissingError(Exception):
    """ Custom error to notify user that spreadsheet does not
    list a specific component """
    pass


def validate_component(components, component_id, sheet_name):
    """ Check if component exists before appending data """
    if not components.get(component_id):
        msg = "{0} component is present in `{1}` but not in `Components` sheet"
        msg = msg.format(component_id, sheet_name)
        raise ComponentMissingError(msg)


def open_workbook(filename='data/xlsx/Control-masonry.xlsx'):
    """ Open the xlsx workbook containing control masonry information """
    return load_workbook(filename=filename)


def extract_components(workbook):
    """ Get the individual components from the xlsx workbook """
    components = dict()
    for row in workbook['Components'].rows[1:]:
        component_id = row[1].value.strip()
        component_name = row[2].value.strip()
        if not components.get(component_id):
            components[component_id] = {}
        components[component_id]['name'] = component_name
    return components


def layer_with_references(components, workbook):
    """ Extract the components worksheet information from the `References`
    worksheet and place the data into a dict """
    for row in workbook['References'].rows[1:]:
        component_id = row[0].value.strip()
        reference_name = row[1].value.strip()
        reference_url = row[2].value.strip()
        validate_component(
            components=components,
            component_id=component_id,
            sheet_name='References'
        )
        if not components[component_id].get('references'):
            components[component_id]['references'] = []
        components[component_id]['references'].append({
            'reference_name': reference_name,
            'reference_url': reference_url,
        })
    return components


def layer_with_governors(components, workbook):
    """ Layer the components data with data from the governors spreadsheet """
    for row in workbook['Governors'].rows[1:]:
        component_id = row[0].value.strip()
        governor_name = row[1].value.strip()
        governor_url = row[2].value.strip()
        validate_component(
            components=components,
            component_id=component_id,
            sheet_name='Governors'
        )
        if not components[component_id].get('governors'):
            components[component_id]['governors'] = []
        components[component_id]['governors'].append({
            'governor_name': governor_name,
            'governor_url': governor_url,
        })
    return components


def layer_with_justifications(components, workbook):
    """ Layer the components data with data from the `Justifications`
    spreadsheet to show which controls each componenet satisfies """
    for row in workbook['Justifications'].rows[1:]:
        control_id = row[0].value.strip()
        component_id = row[2].value.strip()
        narrative = row[3].value.strip()
        validate_component(
            components=components,
            component_id=component_id,
            sheet_name='Justifications'
        )
        if not components[component_id].get('satisfies'):
            components[component_id]['satisfies'] = dict()
        components[component_id]['satisfies'][control_id] = narrative
    return components


def split_into_systems(components, workbook):
    """ Splits the individual components into systems """
    systems = dict()
    for row in workbook['Components'].rows[1:]:
        system_id = row[0].value.strip()
        component_id = row[1].value.strip()
        if not systems.get(system_id):
            systems[system_id] = dict()
        systems[system_id][component_id] = components[component_id]
    return systems


def process_components(workbook):
    """ Collect data from the xlsx workbook and structure data
    in dict by system and then component """
    components = extract_components(workbook=workbook)
    components = layer_with_references(
        components=components, workbook=workbook)
    components = layer_with_governors(
        components=components, workbook=workbook)
    components = layer_with_justifications(
        components=components, workbook=workbook)
    return split_into_systems(
        components=components, workbook=workbook)


def process_certifications(workbook):
    """ Collect standards data from xlsx workbook and extract
    standards dict """
    certifications = {}
    certifications['LATO'] = {'name': 'LATO', 'standards': {}}
    certifications['FISMA-low'] = {'name': 'FISMA-low', 'standards': {}}
    certifications['FISMA-med'] = {'name': 'FISMA-med', 'standards': {}}

    def add_control_to_certification(
            certification, standard, control_id, status):
        """ Determins if control is included in certification and adds to the
        certification dict """
        if status:
            if not certifications[certification]['standards'].get(standard):
                certifications[certification]['standards'][standard] = {}
            certifications[certification]['standards'][standard][control_id] = {}

    for row in workbook['Controls'].rows[1:]:
        standard = row[0].value.strip().upper()
        control_id = row[1].value.strip()
        add_control_to_certification(
            'LATO', standard, control_id, status=row[3].value)
        add_control_to_certification(
            'FISMA-low', standard, control_id, status=row[4].value)
        add_control_to_certification(
            'FISMA-med', standard, control_id, status=row[6].value)
    return certifications


def process_standards(workbook):
    """ Collect certifications data from xlsx workbook and extract
    standards dict, along with controls needed for limited, low, and
    medium ATO certifications """
    standards = {'NIST-800-53': {'name': 'NIST-800-53'}}
    for row in workbook['Controls'].rows[1:]:
        standard = row[0].value.strip().upper()
        control_id = row[1].value.strip()
        name = row[2].value.strip()
        standards[standard][control_id] = {'name': name}
    return standards


def create_folder(directory):
    """ Create a folder if it doesn't exist """
    if not os.path.exists(directory):
        os.makedirs(directory)


def write_yaml_data(data, filename):
    """ Write component data to a yaml file """
    with open(filename, 'w') as yaml_file:
        yaml_file.write(dump(data, default_flow_style=False))


def export_component_yamls(data, base_dir='data/components/'):
    """ Create a series of yaml files for each component organized
    by system """
    create_folder(base_dir)
    for system in data:
        directory = os.path.join(base_dir, system.replace(' ', ''))
        create_folder(directory)
        for component in data[system]:
            filename = os.path.join(
                directory, component.replace(' ', '') + '.yaml')
            write_yaml_data(
                data=data[system][component],
                filename=filename
            )


def export_generic_yamls(data, base_dir):
    """ Create a series of yaml files for each key """
    create_folder(base_dir)
    for key in data:
        filename = os.path.join(
            base_dir, key.replace(' ', '') + '.yaml')
        write_yaml_data(data=data[key], filename=filename)


if __name__ == "__main__":
    workbook = open_workbook()
    export_component_yamls(data=process_components(workbook))
    export_generic_yamls(
        data=process_certifications(workbook),
        base_dir='data/certifications/'
    )
    export_generic_yamls(
        data=process_standards(workbook),
        base_dir='data/standards/'
    )
