import click
import os
import logging

from src.renderers import (
    yamls_to_certification, certifications_to_gitbook, inventory_builder
)
from src import template_generator


@click.group()
@click.option('--verbose', '-v', is_flag=True, help='Toggle logging')
def main(verbose):
    if verbose:
        click.echo('Verbose Mode On')
        logging.basicConfig(level=logging.DEBUG)
    else:
        logging.basicConfig(level=logging.CRITICAL)


@main.command()
@click.option(
    '--data-dir', '-d',
    type=click.Path(exists=True),
    help='Directory containing components, standards, and certifications data.'
)
@click.option(
    '--output-dir', '-o',
    type=click.Path(exists=False),
    help='Directory where certifications is exported'
)
def certs(data_dir, output_dir):
    """ Create certification yamls """
    output_path = yamls_to_certification.create_yaml_certifications(
        data_dir, output_dir
    )
    click.echo('Certifications Created in: `{0}`'.format(output_path))


@main.command()
@click.argument('export-format')
@click.argument('certification')
@click.option(
    '--certs-dir', '-c',
    type=click.Path(exists=True),
    help='Directory containing certification yamls'
)
@click.option(
    '--output-dir', '-o',
    type=click.Path(exists=False),
    help='Directory where documentation is exported'
)
def docs(export_format, certification, certs_dir, output_dir):
    """ Create certification documentation """
    if export_format == 'gitbook':
        output_path = certifications_to_gitbook.create_gitbook_documentation(
            certification, certs_dir, output_dir
        )
        click.echo('Gitbook Files Created in `{0}`'.format(output_path))
    else:
        click.echo('{0} format is not supported yet...'.format(export_format))


@main.command()
@click.argument('certification')
@click.option(
    '--certs-dir', '-c',
    type=click.Path(exists=True),
    help='Directory containing certification yamls'
)
@click.option(
    '--output-dir', '-o',
    type=click.Path(exists=False),
    help='Directory where inventory is exported'
)
def inventory(certification, certs_dir, output_dir):
    """ Creates an inventory for a specific certification  """
    output_path, error = inventory_builder.create_inventory(
        certification, certs_dir, output_dir
    )
    if output_path:
        click.echo('Inventory yaml created at `{0}`'.format(output_path))
    else:
        click.echo(error)


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
