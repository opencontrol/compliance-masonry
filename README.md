# Compliance Masonry

[![Go Report Card](https://goreportcard.com/badge/github.com/opencontrol/compliance-masonry)](https://goreportcard.com/report/github.com/opencontrol/compliance-masonry)
[![Code Climate](https://codeclimate.com/github/opencontrol/compliance-masonry/badges/gpa.svg)](https://codeclimate.com/github/opencontrol/compliance-masonry)
[![codecov.io](https://codecov.io/github/opencontrol/compliance-masonry/coverage.svg?branch=master)](https://codecov.io/github/opencontrol/compliance-masonry?branch=master)
[![Circle CI](https://circleci.com/gh/opencontrol/compliance-masonry/tree/master.svg?style=svg)](https://circleci.com/gh/opencontrol/compliance-masonry/tree/master)
[![Build status](https://ci.appveyor.com/api/projects/status/jjjo83ewacbwnthy/branch/master?svg=true)](https://ci.appveyor.com/project/opencontrol/compliance-masonry/branch/master)

Compliance Masonry is a command-line interface (CLI) that allows users to construct certification documentation using the [OpenControl Schema](https://github.com/opencontrol/schemas). Additional documentation:

* [Benefits](#benefits)
* [18F blog post about Compliance Masonry](https://18f.gsa.gov/2016/04/15/compliance-masonry-buildling-a-risk-management-platform/)
* [Compliance Masonry for the Compliance Literate](docs/masonry-for-the-compliance-literate.md)
* [Developer documentation](CONTRIBUTING.md#development)

![screen shot 2016-04-12 at 12 22 02 pm](docs/assets/data_flow.png)


## Installation
Compliance Masonry is packaged into a downloadable executable program for those
who want to use Compliance Masonry without the need to install any external
dependencies or programs. This is extremely beneficial for those who may to
jump right in without worrying about much technical setup.

### Mac OS X

In your terminal, run the following:

```sh
cd ~/Downloads
curl -L https://github.com/opencontrol/compliance-masonry/releases/download/v1.1.2/compliance-masonry_1.1.2_darwin_amd64.zip -o compliance-masonry.zip
unzip compliance-masonry.zip
cp compliance-masonry_1.1.2_darwin_amd64/compliance-masonry /usr/local/bin
```

### Windows

1. Go to [the Github Release](https://github.com/opencontrol/compliance-masonry/releases/tag/v1.1.2).
1. Download the package that corresponds to your machine and operating system.
    - For 32 Bit Windows, you'll want the file ending `_windows_386.zip`
    - For 64 Bit Windows, you'll want the file ending `_windows_amd64.zip`
1. Double-click on the downloaded package to unzip the archive. The resulting folder should contain a file called `compliance-masonry.exe`.
1. Create a folder, e.g. `C:\Masonry\bin`.
1. Drag `compliance-masonry.exe` into the new folder.
1. Open PowerShell.
    * Search your Start menu / Cortana for it.
1. [Add `C:\Masonry\bin` to your `PATH`.](https://www.java.com/en/download/help/path.xml)

### Linux

The instructions below are for 64-bit architectures. See the [releases](https://github.com/opencontrol/compliance-masonry/releases) page for others.

In your terminal, run the following:

```sh
curl -L https://github.com/opencontrol/compliance-masonry/releases/download/v1.1.2/compliance-masonry_1.1.2_linux_amd64.tar.gz -o compliance-masonry.tar.gz
tar -xf compliance-masonry.tar.gz
cp compliance-masonry_1.1.2_linux_amd64/compliance-masonry /usr/local/bin
```

## Creating an OpenControl project

1. Start a fresh directory

    ```bash
    mkdir your-project-name && cd your-project-name
    ```

1. Create an [`opencontrol.yaml`](https://github.com/opencontrol/schemas#opencontrolyaml)
1. Collect dependencies

    ```bash
    compliance-masonry get
    ```

The `get` command will retrieve dependencies needed to compile documentation in an `opencontrols/` folder. You will probably want to exclude this from your version control system (e.g. add `opencontrols/` to your `.gitignore`).

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

The `gitbook` command by default will create a folder called `exports` that contains the files needed to create a gitbook. Visit the [gitbook documentation](https://github.com/GitbookIO/gitbook-cli) for more information on creating gitbooks via the CLI.

### Adding additional markdown content to Gitbook documentation (optional)
Security documentation usually requires unstructured information that is not captured in the control documentation. The `markdowns` directory can be used to add this supplemental information.

1. Create a `markdowns` folder in the same directory as the `opencontrol.yaml`.

    ```bash
    mkdir markdowns
    ```

2. Create the `markdowns/SUMMARY.md` and `markdowns/README.md` documents.
    ```bash
    touch markdowns/SUMMARY.md
    touch markdowns/README.md
    ```

The content of the `markdowns/SUMMARY.md` and `markdowns/README.md` files and the files they reference is prepended to the Gitbook documentation.

For more information on using the `SUMMARY.md` and `README.md` files visit the [Gitbook documentation.](http://toolchain.gitbook.com/structure.html) For an example `markdowns` directory visit the [cloud.gov `markdowns`](https://github.com/18F/cg-compliance/tree/master/markdowns).

### Viewing gitbook locally in browser

Requires [NodeJS](https://nodejs.org/). After running the steps above,

1. Install the gitbook CLI

    ```bash
    npm install -g gitbook-cli
    ```

1. Navigate to the `exports` directory

    ```bash
    cd exports
    ```

1. Serve the gitbook site locally

    ```bash
    gitbook serve
    ```

1. Open the site: http://localhost:4000

After making any edits, view the changes by running

```bash
compliance-masonry get && compliance-masonry docs gitbook <certification>
```

### Export gitbook as a PDF

1. Following [the steps above](#creating-gitbook-documentation)
1. Navigate to the `exports` directory

    ```bash
    cd exports
    ```

1. Follow [these instructions](http://toolchain.gitbook.com/ebook.html)

## Create Docx template

While there used to be Word document templating logic in Masonry, the team [found](https://github.com/opencontrol/compliance-masonry/issues/153) that it could be done more effectively with rendering code tailored to the specifics of the destination `*.docx`. See the [FedRAMP templater](https://github.com/opencontrol/fedramp-templater) for an example of using Compliance Masonry as a library to inject [OpenControl-formatted](https://github.com/opencontrol/schemas) documentation into a Word doc.

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

Compliance Masonry examples in the wild (in order of completeness):

* cloud.gov [data](https://github.com/18F/cg-compliance) and [gitbook](https://compliance.cloud.gov/)
* [18F Identity-IdP (Upaya)](https://github.com/18F/identity-idp/tree/master/docs/security)
* [18F Confidential Survey](https://github.com/18F/cg-compliance/pull/33) ([needs update](https://github.com/18F/compliance-toolkit/issues/23))

## Documentation Format

Compliance Masonry uses the [OpenControl schema](https://github.com/opencontrol/schemas).

| Type | Supported versions |
|---|---|
| [Components](https://github.com/opencontrol/schemas#components) | [2.0.0](https://github.com/opencontrol/schemas/blob/master/kwalify/component/v2.0.0.yaml), [3.0.0](https://github.com/opencontrol/schemas/blob/master/kwalify/component/v3.0.0.yaml), 3.1.0 |
| [Standards](https://github.com/opencontrol/schemas#standards) | 1.0.0 |
| [Certifications](https://github.com/opencontrol/schemas#certifications) | 1.0.0 |
| [opencontrol.yaml](https://github.com/opencontrol/schemas#opencontrolyaml) | [1.0.0](https://github.com/opencontrol/schemas/blob/master/kwalify/opencontrol/v1.0.0.yaml) |

## Benefits

Modern applications are built on existing systems such as S3, EC2, and Cloud Foundry. Documentation for how these underlying systems fulfill NIST controls or PCI SSC Data Security Standards is a prerequisite for receiving authorization to operate (ATO). Unlike most [System Security Plan documentation](http://csrc.nist.gov/publications/nistpubs/800-18-Rev1/sp800-18-Rev1-final.pdf), Compliance Masonry documentation is built using [OpenControl Schema](https://github.com/opencontrol/schemas), a machine readable format for storing compliance documentation.

Compliance Masonry simplifies the process of certification documentations by providing:

1. a data store for certifications (ex FISMA), standards (ex NIST-800-53), and the individual system components (ex AWS-EC2).
1. a way for government project to edit existing files and also add new control files for their applications and organizations.
1. a pipeline for generating clean and standardized certification documentation.
