# Compliance Masonry - Alpha
[![Build Status](https://travis-ci.org/opencontrol/compliance-masonry.svg?branch=master)](https://travis-ci.org/18F/control-masonry)
[![Code Climate](https://codeclimate.com/github/opencontrol/compliance-masonry/badges/gpa.svg)](https://codeclimate.com/github/opencontrol/compliance-masonry)
## About
Compliance Masonry allows users to construct certification documentation, which is required for approval of government IT systems and applications.

Alpha Note: Compliance Masonry is an emerging project. We recognize that in its current state, Compliance Masonry's user experience still needs to mature. Nevertheless, we are "eating our own dog food" and working to make continuous improvements.

# Quick Start with CLI

### Installing
Only Tested on Python 3+
```bash
$ https://github.com/opencontrol/compliance-masonry.git
$ cd compliance-masonry
$ python setup.py install
```

### New Masonry Project
```bash
masonry init
```
New data directory will be created called `data` containing certifications, standards, and components folders.

### New Component template
```bash
masonry new component system_name component_name
```
New component template will be created as `data/components/system_name/component_name.yaml`

### Create certification yamls
```bash
masonry certs FedRAMP-low
```
Creates certification yamls in `exports/certifications`

### Create documentations
```bash
masonry docs gitbook FedRAMP-low
```
Generates the markdowns for a gitbook.

### Create Inventory
```bash
masonry inventory FedRAMP-low
```
Generates a yaml inventory of listing  missing certification and components documentation.

### Add General Documentation
General documentation can be concatenated to gitbook documentation placing gitbook compatible format in the `data/markdowns/gitbook` directory.  

# Documentation Format

### Components Documentation
Component documentation contains information about individual system components and the standards they satisfy.

```yaml
name: Amazon Elastic Compute Cloud # Name of the component
documentation_complete: false # Manual check if the documentation is complete (for gap analysis)
references:
  - name: Reference  # Name of the reference ie. EC2 website
    url: Refernce URL  # Url of the reference
    type: URL # type of reference (will affect how it's rendered in the documentation)
verifications:
  EC2_Verification_1: # ID of verification
    name: EC2 Verification 1  # Name of verification
    url: Verification 1 URL #  URL of the verification
    type: URL # type of reference (will affect how it's rendered in the documentation)
  EC2_Verification_2:
    name: EC2 Governor 2
    url: Verification 2 URL
    type: Image
satisfies:
  NIST-800-53:
    CM-2:
      narrative: Justification in narrative form # Justification text
      implementation_status: partial # Manual status of implementation (for gap analysis)
      references:
        - verification: EC2_Verification_1 # The specific verification ID that the reference links, no component or system is needed for internal references
        - system: CloudFoundry  # System name of the verification (can link to other systems / components)
          component: UAA  # System name of the verification (can link to other systems / components)
          verification: UAA_Verification_1 # The specific verification ID that the reference links to
```

### Standards Documentation
Contain information about security standards.

```yaml
# nist-800-53.yaml
standards:
  C-2:
    name: User Access
    description: There is an affordance for managing access by...

# PCI.yaml
standards:
  Regulation-6:
    name: User Access PCI
    description: There is an affordance for managing access by...
```

### Certifications
Empty yaml for creating certification documentation. Serve as a template for combining controls and standards yamls.

```yaml
# Fisma.yaml
standards:
  nist-800-53:
    C-2:
    C-3:
  PCI:
    6:
```

## Benefits
Modern applications are build on existing systems such as S3, EC2, and Cloud Foundry. Documentation for how these underlying systems fulfill NIST controls or PCI SSC Data Security Standards is a prerequisite for receiving authorization to operate (ATO). Unlike most [System Security Plan documentation](http://csrc.nist.gov/publications/nistpubs/800-18-Rev1/sp800-18-Rev1-final.pdf), Compliance Masonry documentation is organized by components making it easier for engineers and security teams to collaborate.

Compliance Masonry simplifies the process of certification documentations by providing:
1. a data store for certifications (ex FISMA), standards (ex NIST-800-53), and the individual system components (ex AWS-EC2).
2. a way for government project to edit existing files and also add new control files for their applications and organizations.
3. a pipeline for generating clean and standardized certification documentation


### Long Term Plan Diagram
![compliance-masonry](https://cloud.githubusercontent.com/assets/47762/9829499/08d2b1dc-58bb-11e5-8185-5dc617188ae7.png)
(Here's [the .gliffy source](https://gist.github.com/mogul/8d7cb123e03b0fe1b993).)

### Data Flow Diagram
![compliance-masonry](https://cloud.githubusercontent.com/assets/4596845/10542998/e6397422-73e9-11e5-8681-5539be8b8164.png)
