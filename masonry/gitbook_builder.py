""" This script uses uses core masonry objects to create gitbook output """
import glob
import os
import shutil


from masonry.core import Certification, Standard, Control, System, Component
from src.utils import create_dir


def write_markdown(path, content):
    """ Exports content to the path in a format compatible with markdown """
    with open(path, 'w') as stream:
        stream.write(content)


def concat_markdowns(markdown_path, output_path):
    """ Add markdown content files to the gitbook directory and make the summary
    file the base summary string in order to join the markdown summary with
    the gitbook generated in this file. """
    for filename in glob.iglob(os.path.join(markdown_path, "*", "*")):
        # Get the output file path and create the directory before copying
        output_filepath = os.path.join(
            output_path, filename.replace(os.path.join(markdown_path, ''), '')
        )
        ouput_dir = os.path.dirname(output_filepath)
        create_dir(ouput_dir)
        shutil.copy(filename, output_filepath)
    summary_path = os.path.join(markdown_path, 'SUMMARY.md')
    with open(summary_path, 'r') as summary_file:
        main_summary = summary_file.read()
    return main_summary


def prepend_markdown(export_dir, markdown_path):
    if markdown_path and os.path.exists(markdown_path):
        return concat_markdowns(markdown_path, export_dir)
    return ''


class GitbookComponent(Component):
    """ GitbookComponent loads component and exports data in gitbook format """
    def __init__(self, component_directory=None, component_dict=None):
        super().__init__(
            component_directory=component_directory,
            component_dict=component_dict
        )
        self.import_dir = ''
        self.export_dir = ''

    def move_reference_file(self, reference):
        """ Moves the locally stored reference files into the new gitbook
        directory """
        rel_file_path = reference.get('path')
        import_path = os.path.join(self.import_dir, rel_file_path)
        export_dir = os.path.split(self.export_dir)[0]
        create_dir(os.path.join(self.export_dir, self.meta['component_key']))
        export_path = os.path.join(export_dir, rel_file_path)
        shutil.copy(import_path, export_path)

    def export_markdown_reference(self, reference):
        """ Exports the text of references in gitbook markdown format """
        text = ''
        if reference.get('type').lower() == 'url':
            text += '[{0}]({1})  \n'.format(
                reference.get('name'), reference.get('path')
            )
        elif reference.get('type').lower() == 'image':
            self.move_reference_file(reference)
            text += '![{0}]({1})  \n'.format(
                reference.get('name'), reference.get('path')
            )
        return text

    def get_reference_text(self):
        """ Gets the text of references in gitbook markdown format """
        references = self.meta.get('references')
        if not references:
            return ''
        text = '## References  \n\n'
        for reference in references:
            text += self.export_markdown_reference(reference)
        return text

    def get_verifications_text(self):
        """ Gets the text of verifications in gitbook markdown format """
        verifications = self.meta.get('verifications')
        if not verifications:
            return ''
        text = "## Verifications  \n\n"
        for verification_key, verification in verifications.items():
            text += "#### {0}\n".format(verification_key)
            text += self.export_markdown_reference(verification)
        return text

    def export_gitbook(self, export_dir, import_dir, key):
        """ Export components in gitbook format """
        self.import_dir = import_dir
        self.export_dir = export_dir
        export_path = '{0}-{1}.md'.format(export_dir, key)
        text = '# {0}  \n\n'.format(self.meta['name'])
        text += self.get_reference_text()
        text += '\n\n'
        text += self.get_verifications_text()
        write_markdown(path=export_path, content=text)
        return export_path


class GitbookSystem(System):
    """ GitbookSystem loads systems and exports data in gitbook format """

    def __init__(self, system_directory=None, system_dict=None):
        super().__init__(
            system_directory=system_directory,
            system_dict=system_dict,
            component_class=GitbookComponent
        )

    def export(self, export_dir, import_dir, key):
        """ Exports the systems into gitbook format and returns a list of
        components to help with the creation of a summary """
        export_path = os.path.join(export_dir, key)
        create_dir(export_path)
        system_summary = '# {0}  \n\n'.format(self.meta['name'])
        components = []
        for component_key in sorted(self.components):
            self.components[component_key].export_gitbook(
                export_dir=export_path,
                import_dir=import_dir,
                key=component_key
            )
            system_summary += '[{0}]({1})'.format(
                self.components[component_key].meta.get('name'),
                '{0}-{1}.md'.format(key, component_key)
            )
            components.append(component_key)
        summary_file_path = os.path.join(export_dir, '{0}.md'.format(key))
        write_markdown(path=summary_file_path, content=system_summary)
        return components


class GitbookControl(Control):
    """ GitbookControl loads controls and exports data in gitbook format """

    @staticmethod
    def extract_just_text(justification):
        """ Returns the text of the control in markdown format """
        text = '## {0}  \n'.format(justification['component'])
        text += '{0}  \n'.format(justification['narrative'])
        if 'references' in justification:
            text += '### Verified By:  \n'
            for reference in justification['references']:
                text += '[{0} in {1} {2}]({3})  \n'.format(
                    reference['verification'],
                    reference['system'],
                    reference['component'],
                    os.path.join(
                        '..',
                        'components',
                        '{0}-{1}.md'.format(
                            reference['system'],
                            reference['component']
                        )
                    )
                )
        return text

    def export_markdown(self, export_path):
        """ Export control in markdown format """
        text = '# {0}  \n'.format(self.meta['name'])
        system_text_dict = {}
        for justification in self.justifications:
            system_text_dict[justification['system']] = \
                system_text_dict.get(justification['system'], '') +\
                self.extract_just_text(justification)
        for system, system_text in system_text_dict.items():
            text += '## {0}  \n'.format(system)
            text += system_text

        write_markdown(path=export_path, content=text)

    def export_gitbook(self, export_dir, control_key):
        """ Export a control data in gitbook format returns the family of the
        control to help with the creation of the summary  """
        file_path = '{0}-{1}.md'.format(export_dir, control_key)
        self.export_markdown(file_path)
        return self.meta['family']


class GitbookStandard(Standard):
    """ GitbookStandard loads standards and exports data in gitbook format """
    def __init__(self, standards_yaml_path=None, standard_dict=None):
        """ Overwrites the Standard __init__ method to include a different
        class for exporting gitbooks """
        super().__init__(
            standards_yaml_path=standards_yaml_path,
            standard_dict=standard_dict,
            control_class=GitbookControl
        )

    def export_gitbook(self, export_dir, key):
        """ Exports standard the standard and controls and returns a summary
        """
        summary_text = ''
        family_dict = {}
        export_path = os.path.join(export_dir, key)
        relative_export_path = os.path.join('standards', key)

        for control_key, control in self.controls.items():
            family = control.export_gitbook(
                export_dir=export_path, control_key=control_key
            )
            family_dict[family] = family_dict.get(family, []) \
                + [str(control_key)]

        for family in sorted(family_dict):
            family_path = '{0}-{1}.md'.format(relative_export_path, family)
            family_summary_path = os.path.join(
                os.path.split(export_dir)[0], family_path
            )
            family_summary_text = '# {0}  \n\n## Controls  \n\n'.format(family)
            summary_text += '* [{0} - {1}]({2})\n'.format(
                key, family, family_path
            )
            for control in sorted(family_dict[family]):
                family_summary_text += '* [{0}]({1})\n'.format(
                    control,
                    '{0}-{1}.md'.format(key, control)
                )
                summary_text += '\t* [{0}]({1})\n'.format(
                    control,
                    '{0}-{1}.md'.format(relative_export_path, control)
                )
            write_markdown(
                path=family_summary_path,
                content=family_summary_text
            )
        return summary_text


class GitbookBuilder(Certification):
    """ GitbookBuilder loads certification data and exports a gitbook """
    def __init__(self, certification_yaml_path):
        """ Overwrites the Certification __init__ method to include a different
        class for exports """
        super().__init__(
            certification_yaml_path=certification_yaml_path,
            standard_class=GitbookStandard,
            system_class=GitbookSystem
        )

    def export_systems(self, export_dir):
        """ Exports the systems and returns a summary text of the export
        paths """
        summary_text = ''
        components_export_path = os.path.join(export_dir, 'components')
        create_dir(components_export_path)
        componets_import_path = os.path.split(self.certification_yaml_path)[0]
        for system_key in sorted(self.systems):
            components = self.systems[system_key].export(
                export_dir=components_export_path,
                import_dir=componets_import_path,
                key=system_key
            )
            summary_text += '* [{0}]({1})  \n'.format(
                system_key,
                os.path.join('components', '{0}.md'.format(system_key))
            )
            for component_key in sorted(components):
                summary_text += '\t* [{0}]({1})  \n'.format(
                    component_key,
                    os.path.join('components', '{0}-{1}.md'.format(
                        system_key, component_key
                    ))
                )
        return summary_text

    def export_standards(self, export_dir):
        """ Exports the standards and returns a standard summary """
        standard_path = os.path.join(export_dir, 'standards')
        create_dir(standard_path)
        standards_summary = ''.join([
            standard.export_gitbook(export_dir=standard_path, key=key)
            for key, standard in self.standards_dict.items()
        ])
        return standards_summary

    def export(self, export_dir, markdown_path=None):
        """ Exports a gitbook version of the ssp """
        summary_text = prepend_markdown(
            export_dir=export_dir, markdown_path=markdown_path
        )
        summary_text += self.export_standards(export_dir)
        summary_text += self.export_systems(export_dir)
        # Export summaries
        write_markdown(
            path=os.path.join(export_dir, 'README.md'),
            content=summary_text
        )
        write_markdown(
            path=os.path.join(export_dir, 'SUMMARY.md'),
            content=summary_text
        )
