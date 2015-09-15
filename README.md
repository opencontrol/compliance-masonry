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

## Installation & Usage
Control Masonry runs on Node.js

##### To install requirements
```bash
npm install
```

##### Edit control files
Individual control files are located in `/data/controls/`. They can be expanded or edited to meet the needs of individual projects.

##### To build centralized certification files
```
node build_certification.js
```
