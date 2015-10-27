import click
import os
import logging


from control_masonry.renderers import (
    yamls_to_certification, certifications_to_gitbook
)


def prepare_data_paths(data_dir):
    if not data_dir:
        data_dir = 'data'
    certifications = os.path.join(data_dir, 'certifications/*.yaml')
    components = os.path.join(data_dir, 'components/*/*.yaml')
    standards = os.path.join(data_dir, 'standards/*.yaml')
    return certifications, components, standards


def prepare_cert_output_dir(output_dir):
    if not output_dir:
        output_dir = 'exports'
    return os.path.join(output_dir, 'certifications')


def prepare_cert_path(certs_output_path, certification):
    if not certs_output_path:
        certs_output_path = 'exports/certifications/'
    return os.path.join(certs_output_path, '{0}.yaml'.format(certification))


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
    help='Directory where components, standards, and certifications are located.'
)
@click.option(
    '--output-dir', '-o',
    type=click.Path(exists=False),
    help='Directory where certifications and documentation are exported to.'
)
def certs(data_dir, output_dir):
    """ Create certification yamls """
    certs_data_path, comps_data_path, standards_data_path = prepare_data_paths(data_dir)
    certs_output_path = prepare_cert_output_dir(output_dir)
    yamls_to_certification.create_certifications(
        certifications_path=certs_data_path, components_path=comps_data_path,
        standards_path=standards_data_path, output_path=certs_output_path
    )
    click.echo('Created Certifications')


@main.command()
@click.argument('export-format')
@click.argument('certification')
@click.option(
    '--certs-dir', '-c',
    type=click.Path(exists=True),
    help='Directory containing certification yamls'
)
def docs(export_format, certification, certs_dir):
    """ Create certification documentation """
    certification_path = prepare_cert_path(certs_dir, certification)
    if export_format == 'gitbook':
        certifications_to_gitbook.convert_certifications(
            certification_path=certification_path,
            output_path="exports/Pages",
        )
        click.echo('{0} created'.format(export_format))
    else:
        click.echo('{0} format is not supported yet...'.format(export_format))
