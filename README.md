# Compliance Masonry

[![Go Report Card](https://goreportcard.com/badge/github.com/opencontrol/compliance-masonry)](https://goreportcard.com/report/github.com/opencontrol/compliance-masonry)
[![codecov.io](https://codecov.io/github/opencontrol/compliance-masonry/coverage.svg?branch=master)](https://codecov.io/github/opencontrol/compliance-masonry?branch=master)
[![Circle CI](https://circleci.com/gh/opencontrol/compliance-masonry/tree/master.svg?style=svg)](https://circleci.com/gh/opencontrol/compliance-masonry/tree/master)
[![Build status](https://ci.appveyor.com/api/projects/status/jjjo83ewacbwnthy/branch/master?svg=true)](https://ci.appveyor.com/project/opencontrol/compliance-masonry/branch/master)

Compliance Masonry is a command-line interface (CLI) that allows users to construct certification documentation using the [OpenControl Schema](https://github.com/opencontrol/schemas). See [Benefits](#benefits) for more explanation, and learn more [in our blog post about Compliance Masonry](https://18f.gsa.gov/2016/04/15/compliance-masonry-buildling-a-risk-management-platform/). If you're interested in working on the code, see [our developer documentation](CONTRIBUTING.md#development).

![screen shot 2016-04-12 at 12 22 02 pm](https://cloud.githubusercontent.com/assets/4596845/14469165/5d27495c-00b1-11e6-9d28-327938463adf.png)

## Quick start

1. Install Go 1.6, and ensure your `GOPATH` is set. Using [gvm](https://github.com/moovweb/gvm) is recommended.
1. Install the tool

    ```bash
    go get github.com/opencontrol/compliance-masonry
    ```

1. Run the CLI

    ```bash
    compliance-masonry
    ```

## Creating an OpenControl project

1. Start a fresh directory

    ```bash
    mkdir your-project-name && cd your-project-name
    ```

1. Create an opencontrol.yaml files

    ```bash
    touch opencontrol.yaml
    ```

1. Edit the opencontrol.yaml to contain the following data:

    ```yaml
    schema_version: "1.0.0" # 1.0.0 is the current opencontrol.yaml schema version
    name: Project_Name # Name of the project
    metadata:
      description: "A description of the system"
      maintainers:
        - maintainer_email@email.com
    components: # A list of paths to components written in the opencontrol format for more information view: https://github.com/opencontrol/schemas
      - ./component-1
    certifications: # An optional list of certifications for more information visit: https://github.com/opencontrol/schemas
      - ./cert-1.yaml
    standards: # An optional list of standards for more information visit: https://github.com/opencontrol/schemas
      - ./standard-1.yaml
    dependencies:
      certifications: # An optional list of certifications stored remotely
        - url: github.com/18F/LATO
          revision: master
      systems:  # An optional list of repos that contain an opencontrol.yaml stored remotely
        - url: github.com/18F/cg-compliance
          revision: master
      standards:   # An optional list of remote repos containing standards info that contain an opencontrol.yaml
        - url: github.com/18F/NIST-800-53
          revision: master
    ```

1. Collect dependencies

    ```bash
    compliance-masonry get
    ```

The `get` command will retrieve dependencies needed to compile documentation.

## Creating Gitbook Documentation

1. Update dependencies

    ```bash
    compliance-masonry get
    ```

1. Run the gitbook command

    ```bash
    compliance-masonry docs gitbook LATO
    # Or
    compliance-masonry docs gitbook FedRAMP-low
    ```

The `gitbook` command by default will create a folder called `exports` that contains the files needed to create a gitbook. Visit the [gitbook documentation](https://github.com/GitbookIO/gitbook-cli) for more information on creating gitbooks via the cli

## Create Docx template

1. Create a Word Document template that uses the following template tag format:

    ```
    Documentation for Standard: NIST-800-53 and Control: CM-2 will be rendered below
    {{ getControl "NIST-800-53@CM-2"}}

    Documentation for Standard: NIST-800-53 and Control: AC-2 will be rendered below
    {{ getControl "NIST-800-53@AC-2"}}
    ```

1. Run the docx command.

    ```bash
    compliance-masonry docs docx -t path/to/template.docx
    ```

Running the `docx` command will by default create a file named `export.docx` in the local directory.

## Gap Analysis

Use Gap Analysis to determine the difference between how many controls you have documented versus the total controls for a given certification. This should be used continually as you work to indicate your compliance progress.

Given you have an `opencontrol.yaml` for your project and you have already collected your dependencies via the `compliance-masonry get` command, run `compliance-masonry diff <the-certification>`:

```bash
# Example
$ compliance-masonry diff FedRAMP-moderate
Number of missing controls: 5
NIST-800-53@CP-7 (1)
NIST-800-53@PS-2
NIST-800-53@PS-3 (3)
NIST-800-53@MP-5
NIST-800-53@PS-7
```

## Examples

Compliance Masonry examples in the wild:

* [cloud.gov compliance data repository](https://github.com/18F/cg-compliance) + [documentation generated by Compliance Masonry](https://compliance.cloud.gov/)

## Documentation Format

Compliance Masonry uses the [OpenControl v2 Schema](https://github.com/opencontrol/schemas).

## Benefits

Modern applications are built on existing systems such as S3, EC2, and Cloud Foundry. Documentation for how these underlying systems fulfill NIST controls or PCI SSC Data Security Standards is a prerequisite for receiving authorization to operate (ATO). Unlike most [System Security Plan documentation](http://csrc.nist.gov/publications/nistpubs/800-18-Rev1/sp800-18-Rev1-final.pdf), Compliance Masonry documentation is built using [OpenControl Schema](https://github.com/opencontrol/schemas), a machine readable format for storing compliance documentation.

Compliance Masonry simplifies the process of certification documentations by providing:

1. a data store for certifications (ex FISMA), standards (ex NIST-800-53), and the individual system components (ex AWS-EC2).
1. a way for government project to edit existing files and also add new control files for their applications and organizations.
1. a pipeline for generating clean and standardized certification documentation.
