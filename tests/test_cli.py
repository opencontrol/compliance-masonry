""" Integration tests for Control Masonry CLI """

import os
import yaml
import pytest
import tempfile

from click.testing import CliRunner
from src import cli


@pytest.fixture
def runner():
    return CliRunner()

TEMP_OUTPUT_DIR = tempfile.TemporaryDirectory()
DATA_DIR = os.path.join(os.path.dirname(__file__), '..', 'fixtures')
CERTS_OUTPUT_DIR = os.path.join(TEMP_OUTPUT_DIR.name, 'exports/')
CERTS_DATA_DIR = os.path.join(
    os.path.dirname(__file__), '..', 'fixtures', 'exports', 'certifications'
)
DOCS_OUTPUT_DIR = os.path.join(TEMP_OUTPUT_DIR.name, 'docs/')
COMPS_OUTPUT_DIR = os.path.join(TEMP_OUTPUT_DIR.name, 'components/')


def load_file(file_path):
    with open(file_path, 'r') as file_content:
        return file_content.read()


def load_yaml_file(yaml_path):
    with open(yaml_path, 'r') as yaml_file:
        data = yaml.load(yaml_file)
    return data


def test_certs_run(runner):
    """ Check that certs command runs properly """
    result = runner.invoke(
        cli.main,
        ['certs', '-d{0}'.format(DATA_DIR), '-o{0}'.format(CERTS_OUTPUT_DIR)]
    )
    output = 'Certifications Created in: `{0}`'.format(
        CERTS_OUTPUT_DIR + 'certifications'
    )
    assert result.exit_code == 0
    assert not result.exception
    assert result.output.strip() == output


def test_certs_yaml(runner):
    """ Check that the certification that was created has the correct
    attributes """
    certs_yaml_file = os.path.join(CERTS_OUTPUT_DIR, 'certifications/LATO.yaml')
    generated_yaml = load_yaml_file(certs_yaml_file)
    expected_yaml = load_yaml_file(os.path.join(DATA_DIR, 'exports/certifications/LATO.yaml'))
    assert generated_yaml == expected_yaml


def test_gitbook_runs(runner):
    """ Check that the gitbook command is runs properly """
    result = runner.invoke(
        cli.main,
        [
            'docs', 'gitbook', 'LATO',
            '-c{0}'.format(CERTS_DATA_DIR),
            '-o{0}'.format(DOCS_OUTPUT_DIR)
        ]
    )
    output = 'Gitbook Files Created in `{0}`'.format(DOCS_OUTPUT_DIR)
    assert result.exit_code == 0
    assert not result.exception
    assert result.output.strip() == output


def test_gitbook_files(runner):
    """ Check that the gitbook files were created properly """
    generated_files = [
        'SUMMARY.md', 'README.md',
        'content/NIST-800-53-AC-2.md',
        'content/NIST-800-53-AC-6.md',
    ]
    for generated_file in generated_files:
        generated_file_path = os.path.join(DOCS_OUTPUT_DIR, generated_file)
        expected_file_path = os.path.join('fixtures/docs', generated_file)
        assert os.path.exists(generated_file_path)
        assert load_file(generated_file_path) == load_file(expected_file_path)


def test_gitbook_catches_unsupported_type_error(runner):
    """ Check that the gitbook command is runs properly """
    result = runner.invoke(
        cli.main,
        [
            'docs', 'gitbooc', 'LATO',
            '-c{0}'.format(CERTS_DATA_DIR),
            '-o{0}'.format(DOCS_OUTPUT_DIR)
        ]
    )
    assert result.output.strip() == "gitbooc format is not supported yet..."


def test_new_component_runs(runner):
    """ Check that the new component command is runs properly """
    result = runner.invoke(
        cli.main,
        [
            'new', 'component', 'AWS', 'EC2',
            '-o{0}'.format(COMPS_OUTPUT_DIR)
        ]
    )
    output = 'New Component: `{0}`'.format(
        os.path.join(COMPS_OUTPUT_DIR, 'AWS', 'EC2.yaml')
    )
    assert result.exit_code == 0
    assert not result.exception
    assert result.output.strip() == output


def test_new_component_yaml(runner):
    """ Check that the certification that was created has the correct
    attributes """
    comp_yaml_file = os.path.join(COMPS_OUTPUT_DIR, 'AWS/EC2.yaml')
    generated_yaml = load_yaml_file(comp_yaml_file)
    expected_yaml = load_yaml_file(os.path.join(DATA_DIR, 'components/AWS/EC2.yaml'))
    assert generated_yaml == expected_yaml


TEMP_OUTPUT_DIR.cleanup()
