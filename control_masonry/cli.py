import click
import os
import logging


from renderers import (
    yamls_to_certification, certifications_to_pages, certifications_to_gitbook
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
        return ''
    return os.path.join(certs_output_path, '{0}.yaml'.format(certification))


@click.command()
@click.option(
    '--data-dir', '-d',
    help='Directory where components, standards, and certifications are located.'
)
@click.option(
    '--output-dir', '-o',
    help='Directory where certifications and documentation are exported to.'
)
@click.option(
    '--certification', '-c',
    help='Specific certification used to create documentation.'
)
@click.option(
    '--debug', '-debug',
    is_flag=True,
    help='Specific certification used to create documentation.'
)
@click.argument('command', required=True)
def main(command, data_dir, output_dir, certification, debug):
    """ Create certification yamls """

    if debug:
        logging.basicConfig(level=logging.DEBUG)
    else:
        logging.basicConfig(level=logging.CRITICAL)

    certs_data_path, components_path, standards_path = prepare_data_paths(data_dir)
    certs_output_path = prepare_cert_output_dir(output_dir)
    certification_path = prepare_cert_path(certs_output_path, certification)

    if command == "certs":
        yamls_to_certification.create_certifications(
            certifications_path=certs_data_path, components_path=components_path,
            standards_path=standards_path, output_path=certs_output_path
        )
    elif command == "pages":
        certifications_to_pages.convert_certifications(
            certification_path=certification_path,
            output_path="exports/Pages",
        )
    elif command == "gitbook":
        certifications_to_pages.convert_certifications(
            certification_path=certification_path,
            output_path="exports/Pages",
        )
