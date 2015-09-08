---
permalink: /
title: Introduction
---

# Control Masonry

## About
Control Masonry allows users to construct [NIST Control 800-53](https://web.nvd.nist.gov/view/800-53/home
) documentation, which is required for approval of government IT systems and applications.

## Benefits
Modern applications are build on existing systems such as AWS, EC2, and Cloud Foundry. Documentation for how these underlying systems fulfil NIST controls is a prerequisite for receiving authorization to operate (ATO).

Control Masonry simplifies the process of creating control documentations by providing:
1. a control justification store for each of multiple systems in individual yaml files
2. a system for government project to edit existing files and also add new control files for their applications and organizations.
3. a script for combining the files into a single base file
4. a pipeline for generating clean and standardized documentation

## Installation & Usage
Control Masonry runs on Node.js

##### To install requirements
```bash
npm install
```
##### Edit control files
Individual control files are located in `/controls/systems/`. They can be expanded or edited to meet the needs of individual projects.

##### To build centralized control justification file
```
node build_controls.js
```

##### To update control documentation
```
node render_controls.js
```

##### To serve documentation locally
```
cd docs
./go serve
```

For more documentation on the front-end visit [18F Guides Templates](https://github.com/18F/guides-template)

## Control requirement justification structure
This file explains the structure of the individual control justification files. The data is stored in Yaml format for easy management.

YAML Format
```
---
control:
  a:
  - title: Title of control requirement justifications
    justifications:
    - text: Text justification
      image:
        text: Image text
        url: url
      link:
        text: URL text
        url: url
```

JSON format mapping
```json
{
  "control": {
    "a": [
      {
        "title": "Title of control requirement justifications",
        "justifications": [
          {
            "text": "Text justification",
            "image": {"text": "Image text", "url": "url"},
            "link": {"text": "URL text", "url": "url"}
          }
        ]
      }
    ]
  }
}
```

Markdown Format mapping
### Control
#### a
- ##### Title of control requirement justifications
  - Text justification
  - Image Text  
![Image text](http://dummyimage.com/300x100/ffffff/131961.jpg&text=Image+Justification)
  - [URL text](https://18f.gsa.gov/)
