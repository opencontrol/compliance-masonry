""" Unit tests for yamls_to_certification.py  """

import pytest

from src.renderers import yamls_to_certification


def test_prepare_data_paths_default():
    """ Check that default data paths are created given no data dir """
    certs, components, standards = yamls_to_certification.prepare_data_paths('LATO', 'data')
    assert certs == 'data/certifications/LATO.yaml'
    assert components == 'data/components/*/*.yaml'
    assert standards == 'data/standards/*.yaml'


def test_prepare_data_paths_custom():
    """ Check that default data paths are created given a data directory """
    certs, components, standards = yamls_to_certification.prepare_data_paths('LATO', 't')
    assert certs == 't/certifications/LATO.yaml'
    assert components == 't/components/*/*.yaml'
    assert standards == 't/standards/*.yaml'


def test_copy_key_success():
    """ Check that key is transfered to new dict using a deep copy """
    old_dict = {'test_key': 'test_value'}
    new_dict_1 = {}
    new_dict_2 = {}
    key = 'test_key'
    yamls_to_certification.copy_key(new_dict_1, old_dict, key)
    yamls_to_certification.copy_key(new_dict_2, old_dict, key)
    assert old_dict == new_dict_1
    assert old_dict == new_dict_2


def test_copy_key_failure():
    """ Check that key that copy_key doesn't break when key doesn't exist """
    old_dict = {'test_key': 'test_value'}
    new_dict = {}
    key = 'test_key_fake'
    yamls_to_certification.copy_key(new_dict, old_dict, key)
    assert new_dict == {}


if __name__ == '__main__':
    pytest.main()
