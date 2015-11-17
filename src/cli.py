import click
import os
import logging

from src.renderers import (
    yamls_to_certification, certifications_to_gitbook, inventory_builder
)
from src import template_generator
from src import utils


def verify_certification_path(certification, certs_path):
    """ Check if the certification exists. If the certification does not exist
    print a list of certifications that do exist, if it does return
    the certification path """
    cert_path, possible_certs = utils.check_certifications(certification, certs_path)
    if possible_certs:
        click.echo('{0} is not available options are:\n{1}'.format(
            certification, '\n'.join(possible_certs)
        ))
        return
    else:
        return cert_path


@click.group()
@click.option('--verbose', '-v', is_flag=True, help='Toggle logging')
def main(verbose):
    if verbose:
        click.echo('Verbose Mode On')
        logging.basicConfig(level=logging.DEBUG)
    else:
        logging.basicConfig(level=logging.CRITICAL)


@main.command()
@click.argument('certification')
@click.option(
    '--data-dir', '-d',
    type=click.Path(exists=True),
    default='data',
    help='Directory containing components, standards, and certifications data.'
)
@click.option(
    '--output-dir', '-o',
    type=click.Path(exists=False),
    default='exports/certifications',
    help='Directory where certifications is exported'
)
def certs(certification, data_dir, output_dir):
    """ Create certification yamls """
    utils.create_dir(output_dir)
    certs_dir = os.path.join(data_dir, 'certifications')
    if verify_certification_path(certification, certs_dir):
        output_path = yamls_to_certification.create_yaml_certifications(
            certification, data_dir, output_dir
        )
        click.echo('Certification created in: `{0}`'.format(output_path))


@main.command()
@click.argument('export-format')
@click.argument('certification')
@click.option(
    '--certs-dir', '-c',
    type=click.Path(exists=True),
    default='exports/certifications',
    help='Directory containing certification yamls'
)
@click.option(
    '--output-dir', '-o',
    type=click.Path(exists=False),
    default='exports',
    help='Directory where documentation is exported'
)
def docs(export_format, certification, certs_dir, output_dir):
    """ Create certification documentation """
    cert_path = verify_certification_path(certification, certs_dir)
    if cert_path:
        if export_format == 'gitbook':
            gitbook_output_dir = os.path.join(output_dir, 'gitbook')
            utils.create_dir(os.path.join(gitbook_output_dir, 'content'))
            output_path = certifications_to_gitbook.create_gitbook_documentation(
                cert_path, gitbook_output_dir
            )
            click.echo('Gitbook Files Created in `{0}`'.format(output_path))
        else:
            click.echo('{0} format is not supported yet...'.format(export_format))


@main.command()
@click.argument('certification')
@click.option(
    '--certs-dir', '-c',
    type=click.Path(exists=True),
    default='exports/certifications',
    help='Directory containing certification yamls'
)
@click.option(
    '--output-dir', '-o',
    type=click.Path(exists=False),
    default='exports/inventory',
    help='Directory where inventory is exported'
)
def inventory(certification, certs_dir, output_dir):
    """ Creates an inventory for a specific certification  """
    utils.create_dir(output_dir)
    cert_path = verify_certification_path(certification, certs_dir)
    if cert_path:
        output_path = inventory_builder.create_inventory(cert_path, output_dir)
        click.echo('Inventory yaml created at `{0}`'.format(output_path))


@main.command()
@click.option(
    '--directory', '-d',
    type=click.Path(exists=False),
    help='Directory where documentation is exported'
)
def init(directory):
    """ Initalize a new control masonry project """
    if not directory:
        directory = 'data'
    if not os.path.exists(directory):
        project_path = template_generator.init_project(directory)
        click.echo('New Project: `{0}`'.format(project_path))
    else:
        click.echo('Directory exists, please choose different directory with -d')


@main.command()
@click.argument('file-type')
@click.argument('system-name')
@click.argument('component-name')
@click.option(
    '--output-dir', '-o',
    type=click.Path(exists=False),
    help='Directory where documentation is exported'
)
def new(file_type, system_name, component_name, output_dir):
    """ Command for generating new yaml files """
    if file_type == "component":
        component_path = template_generator.create_new_component_yaml(
            system_name, component_name, output_dir
        )
        click.echo('New Component: `{0}`'.format(component_path))
    else:
        click.echo('Avaiable file-types: `component`')
