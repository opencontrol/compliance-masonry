import pytest
from click.testing import CliRunner
from control_masonry import cli
import os


@pytest.fixture
def runner():
    return CliRunner()


DATA_DIR = 'control_masonry/fixtures/'
OUTPUT_DIR = 'control_masonry/fixtures/exports/'


def test_certs(runner):
    """ Check that certifications are created properly """
    certs_file = os.path.join(OUTPUT_DIR, 'certifications/LATO.yaml')
    if os.path.exists(certs_file):
        os.remove(certs_file)
    result = runner.invoke(
        cli.main,
        ['certs', '-d{0}'.format(DATA_DIR), '-o{0}'.format(OUTPUT_DIR)]
    )
    assert result.exit_code == 0
    assert not result.exception
    assert result.output.strip() == 'Created Certifications'


'''
def test_cli_with_option(runner):
    result = runner.invoke(cli.main, ['--as-cowboy'])
    assert not result.exception
    assert result.exit_code == 0
    assert result.output.strip() == 'Howdy, world.'


def test_cli_with_arg(runner):
    result = runner.invoke(cli.main, ['Gabriel'])
    assert result.exit_code == 0
    assert not result.exception
    assert result.output.strip() == 'Hello, Gabriel.'
'''
