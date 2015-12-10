import glob
import os
import shutil
import yaml

from src import utils


class Component:
    """ Component stores data from a component yaml and handles
    the export of locally stored artifacts """

    def __init__(self, component_directory):
        """ Initialize a component object by identifying the system and
        component key, loading the metadata from the component.yaml, and
        creating a mapping of the controls it satisfies
        """
        self.component_directory = component_directory
        system_dir, self.component_key = os.path.split(component_directory)
        self.system_key = os.path.split(system_dir)[-1]
        self.load_metadata(component_directory)
        self.justification_mapping = self.prepare_justifications()


    def load_metadata(self, component_directory):
        """ Load metadata from components.yaml """
        self.meta = yaml.load(
            open(os.path.join(component_directory, 'component.yaml'))
        )

    def tag_references(self, control_justification):
        """ References that do not have component or system ID, point to
        the system and component in which they are located. This method tags
        the control justifications so they can be referenced later """
        for reference in control_justification['references']:
            if 'system' not in reference:
                reference['system'] = self.system_key
            if 'component' not in reference:
                reference['component'] = self.component_key

    def prepare_justifications(self):
        """ Create a mapping of the controls this component satisfies and
        ensure that all references have system and component tags """
        justifications = self.meta.get('satisfies', [])
        justification_mapping = {}
        for standard_key, standard in justifications.items():
            justification_mapping[standard_key] = {}
            for control_key, control in standard.items():
                justification_mapping[standard_key][control_key] = [
                    (self.system_key, self.component_key)
                ]
                if 'references' in control:
                    self.tag_references(control)
        return justification_mapping

    def get_justifications(self, standard_key, control_key):
        """ Return a list of the controls this component justifies while
        tagging them with the components standard and control key """
        justifications = self.meta.get('satisfies')[standard_key][control_key]
        justifications.update({
            'component': self.component_key, 'system': self.system_key
        })
        return justifications

    def export_references(self, references, export_dir):
        """ Given a list of references in either list or dict format,
        determin which references were saved locally and saves those to
        the appropriate location in the export directory  """
        relative_base_path = os.path.join(self.system_key, self.component_key)
        output_base_path = os.path.join(export_dir, relative_base_path)
        utils.create_dir(output_base_path)
        for reference in utils.inplace_gen(references):
            path = reference.get('path', 'NONE')
            file_import_path = os.path.join(self.component_directory, path)
            is_local = not ('http://' in file_import_path or 'https://' in file_import_path)
            if os.path.exists(file_import_path) and is_local:
                # Create dir and copy file
                file_output_path = os.path.join(output_base_path, path)
                utils.create_dir(os.path.dirname(file_output_path))
                shutil.copy(file_import_path, file_output_path)
                # Rename path
                file_relative_path = os.path.join(relative_base_path, path)
                reference['path'] = file_relative_path
        return references

    def export_component(self, export_dir):
        """ Return the metadata that is required in the certification documentation """
        return {
            self.component_key: {
                'documentation_complete': self.meta.get('documentation_complete'),
                'name': self.meta.get('name'),
                'verifications': self.export_references(self.meta.get('verifications'), export_dir),
                'references': self.export_references(self.meta.get('references'), export_dir)
            }
        }

    def __str__(self):
        return self.meta.get('name', 'Unnamed Component')


class System:
    """ System stores data from the system yaml along with a dict of Component
    objects that fall under the system """
    def __init__(self, system_directory):
        """ Initializes a System object by identifying the system yaml-file,
        loading the system key, metadata, and all the components under the system
        """
        self.system_directory = system_directory
        self.system_key = os.path.split(system_directory)[-1]
        self.load_metadata(self.system_directory)
        self.load_components(self.system_directory)

    def load_components(self, system_directory):
        """ Load the components under the system and store the data
        in individual component objects """
        components_glob = glob.iglob(
            os.path.join(system_directory, '*', 'component.yaml')
        )
        self.components = {}
        self.justification_mapping = {}
        for component_yaml_path in components_glob:
            component_dir_path = os.path.split(component_yaml_path)[0]
            component_key = os.path.split(component_dir_path)[-1]
            component = Component(component_directory=component_dir_path)
            utils.merge_justification(
                self.justification_mapping, component.justification_mapping
            )
            self.components[component_key] = component

    def load_metadata(self, system_directory):
        """ Load the component metadata """
        self.meta = yaml.load(
            open(os.path.join(system_directory, 'system.yaml'))
        )

    def export_system(self, export_dir):
        """ Export system data and component data """
        component_dict = {}
        for component in self.components:
            component_dict.update(self.components[component].export_component(export_dir))
        return {self.system_key: {'components': component_dict, 'meta': self.meta}}

    def __iter__(self):
        for component_key in self.components:
            yield self.components[component_key]

    def __getitem__(self, component_key):
        if self.components:
            return self.components[component_key]

    def __str__(self):
        return self.meta.get('name', 'Unnamed System')


class Control:
    """ Control stores both control metadata and justifications """
    def __init__(self, control_dict):
        """ Load a control depending of the type of control if control does not
        contain 'meta' store everything as meta data, otherwise store
        meta and justifications separately.  """
        self.justifications = {}
        if 'justifications' not in control_dict:
            self.meta = control_dict
        else:
            self.meta = control_dict.get('meta', {})
            self.justifications = control_dict.get('justifications', [])

    def update_metadata(self, new_control):
        """ Update control metadata with another control """
        self.meta.update(new_control.meta)

    def add_justifications(self, new_justifications):
        """ Add justifications to this control """
        self.justifications = new_justifications

    def export(self):
        """ Export justifications in dict format """
        return {'meta': self.meta, 'justifications': self.justifications}


class Standard:
    """ Standard stores control data from a standard yaml """
    def __init__(self, standards_yaml_path=None, standard_dict=None, control_class=Control):
        """ Given a standard yaml or standard dict load all of the controls"""
        if standards_yaml_path:
            standard_dict = yaml.load(open(standards_yaml_path))
        self.control_class = control_class
        self.standards_yaml_path = standards_yaml_path
        self.load_controls(standard_dict)
        if 'name' in standard_dict:
            self.name = standard_dict['name']

    def load_controls(self, standard_dict):
        """ Open the standars yaml and load all the controls """
        self.controls = {}
        for control_key, control_dict in standard_dict.items():
            if isinstance(control_dict, dict):
                self.controls[control_key] = self.control_class(control_dict)

    def export(self):
        """ Export standards in dict format """
        return {control_key: control.export() for control_key, control in self.controls.items()}

    def __getitem__(self, control_key):
        if self.controls:
            return self.controls[control_key]

    def __str__(self):
        return self.name

    def __iter__(self):
        for control_key in self.controls:
            yield (control_key, self.controls[control_key])


class Certification:
    """ Certification stores data from a certification yaml """
    def __init__(self, certification_yaml_path=None, standard_class=Standard):
        # Set Paths
        self.certification_yaml_path = certification_yaml_path
        self.name = os.path.split(os.path.splitext(certification_yaml_path)[0])[-1]
        # Set standard class
        self.standard_class = standard_class
        # Load Data
        certification_data = yaml.load(open(certification_yaml_path))
        self.load_standards(certification_data)
        self.load_components(certification_data)

    def load_standards(self, certification_dict):
        """ Load the standards inside a certification """
        self.standards_dict = {}
        for standard_key, standard in certification_dict['standards'].items():
            self.standards_dict[standard_key] = self.standard_class(standard_dict=standard)

    def load_components(self, certification_dict):
        self.components = certification_dict.get('components', {})

    def import_components(self, system_components_dict):
        """ Update the components directory in the certification """
        self.components.update(system_components_dict)

    def export(self):
        """ Export certification in dict format """
        return {
            'name': self.name,
            'standards': {
                standard_key: standard.export()
                for standard_key, standard in self.standards_dict.items()
            },
            'components': self.components
        }

    def __getitem__(self, standard_key):
        if self.standards_dict:
            return self.standards_dict[standard_key]

    def __iter__(self):
        for standard_key in self.standards_dict:
            yield (standard_key, self.standards_dict[standard_key])