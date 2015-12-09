import glob
import os
import shutil
import yaml

from masonry.core import Certification
from src import utils

class InventoryBuilder(Certification):
    def __init__(self, certification_yaml_path):
        super().__init__(certification_yaml_path)
        self.inventory = {}
        self.component_inventory = self.inventory_components()
        self.standard_inventory = self.inventory_standards()

    def analyze_attribute(self, attribute):
        """ Check how many elements an attribute has otherwise if it's a list
        if it's not a list return that it's present otherwise return "Missing """
        if isinstance(attribute, list) or isinstance(attribute, dict):
            return len(attribute)
        elif attribute:
            return "Present"
        return "Missing"

    def analyze_control(self, control):
        control_dict = {}
        if not control.justifications:
            control_dict = "Missing Justifications"
        for component in control.justifications:
            system_key = component.get('system', 'No System')
            component_key = component.get('component', 'No Name')
            if not system_key in control_dict:
                control_dict[system_key] = {}
            control_dict[system_key][component_key] = {
                'implementation_status': component.get('implementation_status', 'Missing'),
                'narrative': self.analyze_attribute(component.get('narrative')),
                'references': self.analyze_attribute(component.get('references'))
            }
        return control_dict

    def inventory_components(self):
        component_inventory = {}
        for system_key, system in self.components.items():
            component_inventory[system_key] = {}
            for component_key, component in system['components'].items():
                component_inventory[system_key][component_key] = {
                    'references': self.analyze_attribute(component.get('references')),
                    'verifications': self.analyze_attribute(component.get('verifications')),
                    'documentation_completed': component.get('documentation_complete'),
                }
        return component_inventory

    def inventory_standards(self):
        standard_inventory = {}
        for standard_key, standard in self.standards_dict.items():
            standard_inventory[standard_key] = {}
            for control_key, control in standard:
                standard_inventory[standard_key][control_key] = self.analyze_control(control)
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



if __name__ == "__main__":
    print(InventoryBuilder('exports/certifications/LATO.yaml').export('exports/inventory'))
