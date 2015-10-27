import pytest
from click.testing import CliRunner
from src import cli
import os


@pytest.fixture
def runner():
    return CliRunner()


DATA_DIR = 'fixtures/'
CERTS_OUTPUT_DIR = 'fixtures/exports/'


def test_certs(runner):
    """ Check that certifications are created properly """
    certs_file = os.path.join(CERTS_OUTPUT_DIR, 'certifications/LATO.yaml')
    if os.path.exists(certs_file):
        os.remove(certs_file)
    result = runner.invoke(
        cli.main,
        ['certs', '-d{0}'.format(DATA_DIR), '-o{0}'.format(CERTS_OUTPUT_DIR)]
    )
    assert result.exit_code == 0
    assert not result.exception
    assert result.output.strip() == 'Certifications Created in: `fixtures/exports/certifications`'


def test_gitbook(runner):
    """ Check that gitbook is rendered properly """
    result = runner.invoke(
        cli.main,
        ['gitbook', 'LATO', '-c{0}'.format(CERTS_OUTPUT_DIR)]
    )
    assert result.exit_code == 0
    assert not result.exception
    assert result.output.strip() == 'Certifications Created in: `fixtures/exports/certifications`'
