import click
import os
import logging

from src.renderers import certifications_to_gitbook

from masonry.certification_builder import CertificationBuilder
from masonry.inventory_builder import InventoryBuilder
from masonry.doc_builder import DocumentBuilder

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
    default=os.path.join('exports', 'certifications'),
    help='Directory where certifications is exported'
)
def certs(certification, data_dir, output_dir):
    """ Create certification yamls """
    utils.create_dir(output_dir)
    certs_dir = os.path.join(data_dir, 'certifications')
    if verify_certification_path(certification, certs_dir):
        builder = CertificationBuilder(data_dir)
        builder.export_certification(certification, output_dir)
        click.echo('Certification created in: `{0}`'.format(output_dir))


@main.command()
@click.argument('export-format')
@click.argument('certification')
@click.option(
    '--exports-dir', '-e',
    type=click.Path(exists=True),
    default='exports',
    help='Directory containing certification yamls'
)
@click.option(
    '--data-dir', '-d',
    type=click.Path(exists=False),
    default='data',
    help='Directory containing components, standards, markdowns and certifications data.'
)
@click.option(
    '--output-dir', '-o',
    type=click.Path(exists=False),
    default='exports',
    help='Directory where documentation is exported'
)
def docs(export_format, certification, exports_dir, data_dir, output_dir):
    """ Create certification documentation """
    certs_dir = os.path.join(exports_dir, 'certifications')
    cert_path = verify_certification_path(certification, certs_dir)
    markdown_dir = os.path.join(data_dir, 'markdowns')
    if cert_path:
        if export_format == 'gitbook':
            gitbook_output_dir = os.path.join(output_dir, 'gitbook')
            gitbook_markdown_dir = os.path.join(markdown_dir, 'gitbook')
            utils.create_dir(os.path.join(gitbook_output_dir, 'content'))
            output_path = certifications_to_gitbook.create_gitbook_documentation(
                cert_path, gitbook_output_dir, gitbook_markdown_dir
            )
            click.echo('Gitbook Files Created in `{0}`'.format(output_path))
        elif export_format == 'docx':
            docx_output_dir = os.path.join(output_dir, 'docx')
            utils.create_dir(os.path.join(docx_output_dir))
            output_path = DocumentBuilder(cert_path).export(docx_output_dir)
            click.echo('Docx created at `{0}`'.format(output_path))
        else:
            click.echo('{0} format is not supported yet...'.format(export_format))


@main.command()
@click.argument('certification')
@click.option(
    '--exports-dir', '-e',
    type=click.Path(exists=True),
    default='exports',
    help='Directory containing certification yamls'
)
@click.option(
    '--output-dir', '-o',
    type=click.Path(exists=False),
    default='exports/inventory',
    help='Directory where inventory is exported'
)
def inventory(certification, exports_dir, output_dir):
    """ Creates an inventory for a specific certification  """
    certs_dir = os.path.join(exports_dir, 'certifications')
    utils.create_dir(output_dir)
    cert_path = verify_certification_path(certification, certs_dir)
    if cert_path:
        output_path = InventoryBuilder(cert_path).export(output_dir)
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
@click.argument('system-key')
@click.argument('component-key', required=False)
@click.option(
    '--data-dir', '-d',
    type=click.Path(exists=False),
    default='data',
    help='Directory where documentation is exported'
)
def new(file_type, system_key, component_key, data_dir):
    """ Command for generating new yaml files """
    output_dir = os.path.join(data_dir, 'components')
    if file_type == "system":
        system_path = template_generator.create_new_data_yaml(
            output_dir=output_dir, system_key=system_key, component_key=None
        )
        click.echo('New System: `{0}`'.format(system_path))
    elif file_type == "component":
        component_path = template_generator.create_new_data_yaml(
            output_dir=output_dir, system_key=system_key, component_key=component_key
        )
        click.echo('New Component: `{0}`'.format(component_path))
    else:
        click.echo('Avaiable file-types: `system` or `component`')
