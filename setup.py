"""
Create SSP Documentation
"""
from setuptools import find_packages, setup

dependencies = ['click', 'pyyaml', 'awesome-slugify', 'pykwalify']

setup(
    name='compliance-masonry',
    version='2.1.0',
    url='https://github.com/18F/control-masonry',
    license='BSD',
    author='Gabriel Ramirez',
    author_email='gabriel.e.ramirez@gmail.com',
    description='Build SSP Documentation',
    long_description=__doc__,
    packages=find_packages(exclude=['tests']),
    include_package_data=True,
    zip_safe=False,
    platforms='any',
    install_requires=dependencies,
    entry_points={
        'console_scripts': [
            'masonry = masonry.cli:main',
        ],
    },
    package_data={
        'masonry': [
            '*.py',
        ],
        'helpers': [
            'helpers/*.py',
            'helpers/templates/certifications/*',
            'helpers/templates/components/*/*',
            'helpers/templates/standards/*',
        ]
    },
    classifiers=[
        # As from http://pypi.python.org/pypi?%3Aaction=list_classifiers
        'Development Status :: 3 - Alpha',
        # 'Development Status :: 4 - Beta',
        # 'Development Status :: 5 - Production/Stable',
        # 'Development Status :: 6 - Mature',
        # 'Development Status :: 7 - Inactive',
        'Environment :: Console',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: BSD License',
        # 'Operating System :: POSIX',
        'Operating System :: MacOS',
        'Operating System :: Unix',
        # 'Operating System :: Windows',
        'Programming Language :: Python',
        'Programming Language :: Python :: 2',
        'Programming Language :: Python :: 3',
        'Topic :: Software Development :: Libraries :: Python Modules',
    ]
)
