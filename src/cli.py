import click
import logging

from src.renderers import yamls_to_certification, certifications_to_gitbook


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
