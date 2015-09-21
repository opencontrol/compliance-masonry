""" This script coverts the Control-masonry excel file to
a series of yamls """

from openpyxl import load_workbook


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


def open_workbook(filename='Control-masonry.xlsx'):
    """ Open the xlsx workbook containing control masonry information """
    return load_workbook(filename=filename)


def extract_components(workbook):
    """ Get the individual components from the xlsx workbook """
    components = dict()
    for row in workbook['Components'].rows:
        component_id = row[1].value
        component_name = row[2].value
        if not components.get(component_id):
            components[component_id] = {}
        components[component_id]['name'] = component_name
    return components


def layer_with_references(components, workbook):
    """ Extract the components worksheet information from the `References`
    worksheet and place the data into a dict """
    for row in workbook['References'].rows:
        component_id = row[0].value
        reference_name = row[1].value
        reference_url = row[2].value
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
    for row in workbook['Governors'].rows:
        component_id = row[0].value
        governor_name = row[1].value
        governor_url = row[2].value
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
    for row in workbook['Justifications'].rows:
        control_id = row[0].value
        component_id = row[2].value
        narrative = row[3].value
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
    for row in workbook['Components'].rows:
        system_id = row[0].value
        component_id = row[1].value
        if not systems.get(system_id):
            systems[system_id] = dict()
        systems[system_id][component_id] = components[component_id]
    return systems


def create_yamls():
    """ Create yaml files containing data from the control masonry workbook """
    workbook = open_workbook()
    components = extract_components(workbook=workbook)
    components = layer_with_references(
        components=components, workbook=workbook)
    components = layer_with_governors(
        components=components, workbook=workbook)
    components = layer_with_justifications(
        components=components, workbook=workbook)
    systems = split_into_systems(
        components=components, workbook=workbook)
    return systems


if __name__ == "__main__":
    create_yamls()
