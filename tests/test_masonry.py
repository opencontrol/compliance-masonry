import os
import pytest
import tempfile
import glob


from masonry.core import Component


COMPONENT = Component('fixtures/components/AWS/EC2')


def test_component_load_file():
    """ Check that the component object can load a component yaml and
    store the correct attributes """
    assert COMPONENT.component_directory == 'fixtures/components/AWS/EC2'
    assert COMPONENT.component_key == 'EC2'
    assert COMPONENT.system_key == 'AWS'


def test_component_justification_mapping():
    """ Check that the justifications this component satisfies have been
    correctly stored """
    assert COMPONENT.justification_mapping == {
        'NIST-800-53': {
            'CM-2': [('AWS', 'EC2')]
        },
        'PCI-DSS-MAY-2015': {
            1.1: [('AWS', 'EC2')],
            2.1: [('AWS', 'EC2')],
            '1.1.1': [('AWS', 'EC2')]
        }
    }


def test_get_justifications():
    """ Given a standard and control check that the get_justifications method
    returns the justifications for that control.  """
    justification = COMPONENT.get_justifications('PCI-DSS-MAY-2015', 1.1)
    assert len(justification['references']) == 2


def test_export_references():
    """ Check if the method correctly determins which references were saved
    locally and saves those to the appropriate location in the export directory
    """
    TEMP_OUTPUT_DIR = tempfile.TemporaryDirectory()
    references = COMPONENT.meta.get('verifications')
    # Export files
    COMPONENT.export_references(references, TEMP_OUTPUT_DIR.name)
    # Check if files exists
    assert len(glob.glob(os.path.join(TEMP_OUTPUT_DIR.name, '*', '*', '*'))) == 1
    TEMP_OUTPUT_DIR.cleanup()


def test_export_component():
    """ Check if the method correctly returns the metadata needed for the
    certification """
    TEMP_OUTPUT_DIR = tempfile.TemporaryDirectory()
    # Export files
    data = COMPONENT.export_component(TEMP_OUTPUT_DIR.name)
    assert data['EC2']['name'] == 'Amazon Elastic Compute Cloud'


if __name__ == '__main__':
    pytest.main()
