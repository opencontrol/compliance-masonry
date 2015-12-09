import glob
import os
import shutil
import yaml

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
    def inventory(self):
        control_dict = {}
        if not self.justifications:
            control_dict = "Missing Justifications"
        for component in self.justifications:
            system_key = component.get('system', 'No System')
            component_key = component.get('component', 'No Name')
            if not system_key in control_dict:
                control_dict[system_key] = {}
            control_dict[system_key][component_key] = {
                'implementation_status': component.get('implementation_status', 'Missing'),
                'narrative': analyze_attribute(component.get('narrative')),
                'references': analyze_attribute(component.get('references'))
            }
        return control_dict


class InventoryStandard(Standard):
    def __init__(self, standards_yaml_path=None, standard_dict=None):
        super().__init__(
            standards_yaml_path=standards_yaml_path,
            standard_dict=standard_dict,
            control_class=InventoryControl
        )

    def inventory(self):
        control_inventory = {}
        for control_key, control in self:
            control_inventory[control_key] = control.inventory()
        return control_inventory


class InventoryBuilder(Certification):
    def __init__(self, certification_yaml_path):
        super().__init__(certification_yaml_path, standard_class=InventoryStandard)
        self.inventory = {}
        self.component_inventory = self.inventory_components()
        self.standard_inventory = self.inventory_standards()

    def inventory_components(self):
        component_inventory = {}
        for system_key, system in self.components.items():
            component_inventory[system_key] = {}
            for component_key, component in system['components'].items():
                component_inventory[system_key][component_key] = {
                    'references': analyze_attribute(component.get('references')),
                    'verifications': analyze_attribute(component.get('verifications')),
                    'documentation_completed': component.get('documentation_complete'),
                }
        return component_inventory

    def inventory_standards(self):
        standard_inventory = {}
        for standard_key, standard in self.standards_dict.items():
            standard_inventory[standard_key] = standard.inventory()
        return standard_inventory

    def make_export_dict(self):
        return {
            'certification': self.name,
            'components': self.component_inventory,
            'standards': self.standard_inventory
        }

    def export(self, export_path):
        inventory_path = os.path.join(
            export_path,
            self.name + '.yaml'
        )
        utils.yaml_writer(self.make_export_dict(), inventory_path)
        return inventory_path
