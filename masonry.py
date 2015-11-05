""" An interface for the renderers to be used when importing """

from src.renderers import (
    yamls_to_certification, certifications_to_gitbook, inventory_builder
)


def build_certifications(data_dir, output_dir):
    """ Interface for the yamls_to_certification renderer, which creates yaml
    certifications """
    return yamls_to_certification.create_yaml_certifications(
        data_dir=data_dir, output_dir=output_dir
    )


def build_gitbook(certification, certification_dir, output_dir):
    """ Interface for the certifications_to_gitbook renderer, which creates
    certification documentation in gitbook form """
    return certifications_to_gitbook.create_gitbook_documentation(
        certification=certification,
        certification_dir=certification_dir,
        output_path=output_dir
    )


def build_inventory(certification_path):
    """ Interface for inventory_builder renderer, which creates an inventory
    of components for a specific certification """
    return inventory_builder.build_inventory(certification_path)
