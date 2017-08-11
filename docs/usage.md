# Usage

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

## Docker

Compliance Masonry has also been packaged as a Docker image and published on [Docker Hub](https://hub.docker.com/r/opencontrolorg/compliance-masonry). Commands can be run with Docker in the directory containing `opencontrol.yaml` as follows:

```bash
docker run --rm -v "$PWD":/opencontrol -w /opencontrol opencontrolorg/compliance-masonry get
```

## GitBook

To view the compliance documentation as a web site or a PDF, see the [GitBook](gitbook.md) documentation.

## Create Docx template

While there used to be Word document templating logic in Masonry, the team [found](https://github.com/opencontrol/compliance-masonry/issues/153) that it could be done more effectively with rendering code tailored to the specifics of the destination `*.docx`. See the [FedRAMP templater](https://github.com/opencontrol/fedramp-templater) for an example of using Compliance Masonry as a library to inject [OpenControl-formatted](https://github.com/opencontrol/schemas) documentation into a Word doc.

## Gap Analysis

***Experimental.*** *[Does not take control origination into account.](https://github.com/opencontrol/schemas/issues/24)*

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

## JSON/YAML Extract

***Experimental.*** *[Work performed by other Team; no issue opened yet.]*

You may use `compliance-masonry` to generate conglomerate output in JSON or YAML.

This is two-step process:

1. Use `compliance-masonry get` to gather schema information.
1. Use the `extract` function to render the gathered schema as consolidated output.

### JSON Extract

In this example, transform the gathered input schema from the `[path-to-opencontrols-dir]`, send output to STDOUT (`-d -`), use JSON format, and feed the output through `jq` for readability:

```
compliance-masonry extracto [path-to-opencontrols-dir] -d - -f json fedramp-moderate | jq '.'
```

### YAML Extract

In this example, transform the gathered input schema from the `[path-to-opencontrols-dir]`, send output to STDOUT (`-d -`), and use YAML format:

```
compliance-masonry extracto [path-to-opencontrols-dir] -d - -f yaml fedramp-moderate
```

## Documentation format

Compliance Masonry uses the [OpenControl schema](https://github.com/opencontrol/schemas).

| Type | Supported versions |
|---|---|
| [Components](https://github.com/opencontrol/schemas#components) | [2.0.0](https://github.com/opencontrol/schemas/blob/master/kwalify/component/v2.0.0.yaml), [3.0.0](https://github.com/opencontrol/schemas/blob/master/kwalify/component/v3.0.0.yaml), 3.1.0 |
| [Standards](https://github.com/opencontrol/schemas#standards) | 1.0.0 |
| [Certifications](https://github.com/opencontrol/schemas#certifications) | 1.0.0 |
| [opencontrol.yaml](https://github.com/opencontrol/schemas#opencontrolyaml) | [1.0.0](https://github.com/opencontrol/schemas/blob/master/kwalify/opencontrol/v1.0.0.yaml) |
