# Control Masonry

## About
Control Masonry allows users to construct certification documentation, which is required for approval of government IT systems and applications.

## Benefits
Modern applications are build on existing systems such as AWS, EC2, and Cloud Foundry. Documentation for how these underlying systems fulfil NIST controls or PCI SSC Data Security Standards is a prerequisite for receiving authorization to operate (ATO).

Control Masonry simplifies the process of certification documentations by providing:
1. a control justification store for multiple systems.
2. a system for government project to edit existing files and also add new control files for their applications and organizations.
3. a script for combining the files into a single base file
4. a pipeline for generating clean and standardized certification documentation

## My Cloud.gov/18F Controls/ Your Controls Documentation
Without being combined with the standards and certifications yamls the control yamls can be use to generate readable documentation w/ gitbook, etc...

```yaml
## Cloudgov.yaml
CF_UAA:
  name: Cloud Foundry User Authentication and Authorization (UAA)
  references:
    - name: UAA design doc
      url: https://asdfasdf
    - name: Some other doc
      url: https://boobarbazbat
  governors:
    - name: UAA configuration
      url: https://pathtogitrepohead
    - name: Live test results
      url: https://dashboardwithupdatedtestresults
  satisfies:
    - standard:
        AC-2: Description of now CF_UUA meets point control X sub mod. a
```

##### Edit control files
Individual control files are located in `/data/controls/`. They can be expanded or edited to meet the needs of individual projects.

##### To build centralized certification files
```
node build_certification.js
=======

## Standards Documentation
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

## Certifications
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

## Yaml output
Centralized yaml for a specific certification, can be used to render matrix.csv, gitbook.md, ssp.docx... This is were we will be able to see if any pieces are missing.
```yaml
# NIST-800-53.yaml
AC-2:
  a:
  - title: Title of control requirement justifications
    justifications:
    - id: CF_UAA
      name: Cloud Foundry User Authentication and Authorization (UAA)
      narrative: Description of now CF_UUA meets point control X sub mod. a
      references:
        - name: UAA design doc
          url: https://asdfasdf
        - name: Some other doc
          url: https://boobarbazbat
      governors:
        - name: UAA configuration
          url: https://pathtogitrepohead
        - name: Live test results
          url: https://dashboardwithupdatedtestresults
```
