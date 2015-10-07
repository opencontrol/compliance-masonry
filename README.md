# Control Masonry

## About
Control Masonry allows users to construct certification documentation, which is required for approval of government IT systems and applications.

## Benefits
Modern applications are build on existing systems such as S3, EC2, and Cloud Foundry. Documentation for how these underlying systems fulfill NIST controls or PCI SSC Data Security Standards is a prerequisite for receiving authorization to operate (ATO).

Control Masonry simplifies the process of certification documentations by providing:
1. a data store for certifications (ex FISMA), standards (ex NIST-800-53), and the individual system components (ex AWS-EC2).
2. a way for government project to edit existing files and also add new control files for their applications and organizations.
3. a pipeline for generating clean and standardized certification documentation

## Creating Documentation
Generating certification documentation from components and standards.
```
python renderers/yamls_to_certification.py
```

# Documentation Format

### Components Documentation
Component documentation contains information about individual system components and the standards they satisfy.

```yaml
name: User Account and Authentication (UAA) Server
references:
- reference_name: User Account and Authentication (UAA) Server
  reference_url: http://docs.pivotal.io/pivotalcf/concepts/architecture/uaa.html
- reference_name: Creating and Managing Users with the UAA CLI (UAAC)
  reference_url: http://docs.pivotal.io/pivotalcf/adminguide/uaa-user-management.html
governors:
- governor_name: Cloud Foundry Roles
  governor_url: https://cf-p1-docs-prod.cfapps.io/pivotalcf/concepts/roles.html
- governor_name: Cloud Foundry Org Access
  governor_url: https://github.com/cloudfoundry/cloud_controller_ng/blob/master/spec/unit/access/organization_access_spec.rb
- governor_name: Cloud Foundry Space Access
  governor_url: https://github.com/cloudfoundry/cloud_controller_ng/blob/master/spec/unit/access/space_access_spec.rb
satisfies:
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
