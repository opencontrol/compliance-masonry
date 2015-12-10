""" This script uses the core masonry objects to construct certifications and
export artifacts to certification folder """

import glob
import os
import yaml

from src import utils
from masonry.core import System, Standard, Certification


class CertificationBuilder:
    """ CertificationBuilder loads certification data and exports certification
    documentation """

    def __init__(self, data_directory=None):
        """ Given a data directory loads the systems and standards
        given a certification yaml loads a certification """
        self.data_directory = data_directory
        self.load_systems(self.data_directory)
        self.load_standards(self.data_directory)

    def load_systems(self, data_directory):
        """ Load all the systems in the data data directory """
        systems_glob = glob.iglob(
            os.path.join(data_directory, 'components', '*', 'system.yaml')
        )
        self.systems = {}
        self.justification_mapping = {}
        for system_yaml_path in systems_glob:
            system_dir_path = os.path.split(system_yaml_path)[0]
            system_key = os.path.split(system_dir_path)[-1]
            system = System(system_directory=system_dir_path)
            utils.merge_justification(self.justification_mapping, system.justification_mapping)
            self.systems[system_key] = system

    def load_standards(self, data_directory):
        """ Load standards in the data directory """
        standards_glob = glob.iglob(
            os.path.join(data_directory, 'standards', '*.yaml')
        )
        self.standards = {}
        for standards_yaml_path in standards_glob:
            standard_key = os.path.splitext(
                os.path.split(standards_yaml_path)[-1]
            )[0]
            self.standards[standard_key] = Standard(standards_yaml_path=standards_yaml_path)

    def get_justifications(self, standard_key, control_key):
        """ Given a standard and control return all the justification from the
        components data """
        if standard_key in self.justification_mapping:
            if control_key in self.justification_mapping[standard_key]:
                for system, component in self.justification_mapping[standard_key][control_key]:
                    yield self.systems[system][component].get_justifications(
                        standard_key, control_key
                    )

    def prepare_certification_controls(self, certification):
        """ Prepare a specific certification for export by merging all
        justifications from systems and components """
        for standard_key, standard in certification:
            for control_key, control in standard:
                control.update_metadata(self.standards[standard_key][control_key])
                control.add_justifications(list(self.get_justifications(standard_key, control_key)))

    def create_certification(self, certification):
        """ Create a certification object by updating the controls and loading systems """
        certification_yaml_path = os.path.join(
            self.data_directory, 'certifications', certification + '.yaml'
        )
        certification = Certification(certification_yaml_path=certification_yaml_path)
        certification.import_systems(self.systems)
        self.prepare_certification_controls(certification)
        return certification

    def export_certification(self, certification, export_dir):
        """ Given a certification name and an export directory exports all the
        locally stored references to the directory along with the certification
        yaml. """
        certification_obj = self.create_certification(certification)
        with open(os.path.join(export_dir, certification + '.yaml'), 'w') as cert_file:
            cert_file.write(
                yaml.dump(
                    certification_obj.export(export_dir),
                    default_flow_style=False,
                    indent=2
                )
            )
