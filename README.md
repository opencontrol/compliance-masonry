# Control Masonry (AKA Compliance Masonry) - Alpha
[![Build Status](https://travis-ci.org/18F/control-masonry.svg)](https://travis-ci.org/18F/control-masonry)
[![Code Climate](https://codeclimate.com/github/18F/control-masonry/badges/gpa.svg)](https://codeclimate.com/github/18F/control-masonry)
## About
Control Masonry allows users to construct certification documentation, which is required for approval of government IT systems and applications.

Alpha Note: Control Masonry is an emerging project. We recognize that in its current state, Control Masonry's user experience still needs to mature. Nevertheless, we are "eating our own dog food" and working to make continuous improvements.

# Quick Start with CLI

### Installing
Only Tested on Python 3+
```bash
$ git clone https://github.com/18F/control-masonry.git
$ cd control-masonry
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
masonry certs
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

## Importing Control Masonry
```
import masonry

masonry.build_certifications(
  data_dir="data_directory", output_dir="output_directory"
)

masonry.build_gitbook(
  certification="certification_name",
  certification_dir="location of certification's directory",
  output_dir="location to output gitbook"
)
```

# Documentation Format

### Components Documentation
Component documentation contains information about individual system components and the standards they satisfy.

```yaml
name: User Account and Authentication (UAA) Server
system: CloudFoundry
documentation_complete: true
references:
- name: User Account and Authentication (UAA) Server
  url: http://docs.pivotal.io/pivotalcf/concepts/architecture/uaa.html
- name: Creating and Managing Users with the UAA CLI (UAAC)
  url: http://docs.pivotal.io/pivotalcf/adminguide/uaa-user-management.html
governors:
- name: Cloud Foundry Roles
  url: https://cf-p1-docs-prod.cfapps.io/pivotalcf/concepts/roles.html
- name: Cloud Foundry Org Access
  url: https://github.com/cloudfoundry/cloud_controller_ng/blob/master/spec/unit/access/organization_access_spec.rb
- name: Cloud Foundry Space Access
  url: https://github.com/cloudfoundry/cloud_controller_ng/blob/master/spec/unit/access/space_access_spec.rb
satisfies:
  NIST-800-53:
    AC-2: Cloud Foundry accounts are managed through the User Account and Authentication
      (UAA) Server.
    IA-2: The UAA is the identity management service for Cloud Foundry. Its primary
      role is as an OAuth2 provider, issuing tokens for client applications to use when
      they act on behalf of Cloud Foundry users.
    SC-13: All traffic from the public internet to the Cloud Controller and UAA happens
      over HTTPS and operators configure encryption of the identity store in the UAA
    SC-28 (1): Operators configure encryption of the identity store in the UAA. When
      users register an account with the Cloud Foundry platform, the UAA, acts as the
      user store and stores user passwords in the UAA database using bcrypt. Bcrypt
      is a blowfish encryption algorithm, which enables cloud foundry to store a secure
      hash of your users' passwords.
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
Modern applications are build on existing systems such as S3, EC2, and Cloud Foundry. Documentation for how these underlying systems fulfill NIST controls or PCI SSC Data Security Standards is a prerequisite for receiving authorization to operate (ATO). Unlike most [System Security Plan documentation](http://csrc.nist.gov/publications/nistpubs/800-18-Rev1/sp800-18-Rev1-final.pdf), Control Masonry documentation is organized by components making it easier for engineers and security teams to collaborate.

Control Masonry simplifies the process of certification documentations by providing:
1. a data store for certifications (ex FISMA), standards (ex NIST-800-53), and the individual system components (ex AWS-EC2).
2. a way for government project to edit existing files and also add new control files for their applications and organizations.
3. a pipeline for generating clean and standardized certification documentation


### Long Term Plan Diagram
![control-masonry](https://cloud.githubusercontent.com/assets/47762/9829499/08d2b1dc-58bb-11e5-8185-5dc617188ae7.png)
(Here's [the .gliffy source](https://gist.github.com/mogul/8d7cb123e03b0fe1b993).)

### Data Flow Diagram
![control_masonry](https://cloud.githubusercontent.com/assets/4596845/10542998/e6397422-73e9-11e5-8681-5539be8b8164.png)
