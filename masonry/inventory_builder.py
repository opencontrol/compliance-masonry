""" This script uses uses core masonry objects to analyze missing parts of the
certification """
import os

from masonry.core import Certification, Standard, Control
from src import utils


def analyze_attribute(attribute):
    """ Check how many elements an attribute has otherwise if it's a list
    if it's not a list return that it's present otherwise return "Missing """
    if isinstance(attribute, list) or isinstance(attribute, dict):
        return len(attribute)
    elif attribute:
        return "Present"
    return "Missing"


class InventoryControl(Control):
    """ InventoryControl inherits from the Control class and adds a method
    to create an inventory of justifications """
    def inventory(self):
        """ Create a catalog for a specific controls """
        control_dict = {}
        if not self.justifications:
            control_dict = "Missing Justifications"
        for component in self.justifications:
            system_key = component.get('system', 'No System')
            component_key = component.get('component', 'No Name')
            if system_key not in control_dict:
                control_dict[system_key] = {}
            control_dict[system_key][component_key] = {
                'implementation_status': component.get('implementation_status', 'Missing'),
                'narrative': analyze_attribute(component.get('narrative')),
                'references': analyze_attribute(component.get('references'))
            }
        return control_dict


class InventoryStandard(Standard):
    """ InventoryStandard inherits from Standard, while overriding the Control
    class with InventoryControl as the default storage for control data. This
    class also adds a method to help analyze missing certification gaps """
    def __init__(self, standards_yaml_path=None, standard_dict=None):
        super().__init__(
            standards_yaml_path=standards_yaml_path,
            standard_dict=standard_dict,
            control_class=InventoryControl
        )

    def inventory(self):
        """ Creates a catalog of controls in the system """
        control_inventory = {}
        for control_key, control in self:
            control_inventory[control_key] = control.inventory()
        return control_inventory


class InventoryBuilder(Certification):
    """ InventoryBuilder load certification data and exports a yaml gap analysis """
    def __init__(self, certification_yaml_path):
        super().__init__(certification_yaml_path, standard_class=InventoryStandard)
        self.inventory = {}
        self.systems_inventory = self.inventory_systems()
        self.standard_inventory = self.inventory_standards()

    def inventory_systems(self):
        """ Creates an system/components catalog """
        systems_inventory = {}
        for system_key, system in self.systems.items():
            systems_inventory[system_key] = {}
            for component in system:
                systems_inventory[system_key][component.component_key] = {
                    'references': analyze_attribute(component.meta.get('references')),
                    'verifications': analyze_attribute(component.meta.get('verifications')),
                    'documentation_completed': component.meta.get('documentation_complete'),
                }
        return systems_inventory

    def inventory_standards(self):
        """ Creates an standards/controls catalog """
        standard_inventory = {}
        for standard_key, standard in self.standards_dict.items():
            standard_inventory[standard_key] = standard.inventory()
        return standard_inventory

    def make_export_dict(self):
        """ Creates a dict version of the inventory report """
        return {
            'certification': self.name,
            'components': self.systems_inventory,
            'standards': self.standard_inventory
        }

    def export(self, export_path):
        """ Exports the inventory report to a yaml file """
        inventory_path = os.path.join(
            export_path,
            self.name + '.yaml'
        )
        utils.yaml_writer(self.make_export_dict(), inventory_path)
        return inventory_path
