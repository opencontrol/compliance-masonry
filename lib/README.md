# Lib Package

## Purpose

This package (and the sub packages) are the only packages that should
be imported by external projects. This lib package is very helpful for
those writing plugins to extend the functionality of Masonry.

Additionally, 1) the `lib/common` package contains all the interfaces for
workspace information and 2) the `lib/common/mocks` contains all the
mocks to help with tests.

The idea is once you run the `get` command and all the resources
(YAML, etc) are placed in a folder, those resources can be loaded into
 a `Workspace` via code and you'll have access to all of that info.
 
## Development

### Usage
 
There are many interfaces but the main ones are:

#### Workspace
Workspace is the representation of your working space with all the
resources gathered together.
- How to obtain a workspace:
  ```go
  import "github.com/opencontrol/compliance-masonry/lib"
  
  // some other code
  
  ws := lib.LoadData("where-get-command-placed-things", "certification-path")
  ```

#### Standard
Standard is representation of all the controls for a certain standard 
(e.g. NIST-800-53).

Once you have a workspace object, you can use `GetStandard` and provide
a standard key. For more information about the key value to used or more
information about standards, refer to the
[standard schema](https://github.com/opencontrol/schemas#standards).

#### Component
Component is a basic block of compliance information that corresponds to
a control or set of controls.

Once you have a workspace object, you can use `GetAllComponents` or
`GetComponent`and provide a component key. For more information about 
the key value to used or more information about the component, refer to
the [component schema](https://github.com/opencontrol/schemas#components)

#### Certification
Certification is a list of controls that make up a certain
"certification".

Once you have a workspace object, you can use `GetCertification`.
For more information about the certification, refer to the
[certification schema](https://github.com/opencontrol/schemas#certifications)

### Result Data
`Verification` is a data structure that is not represented in yaml but
rather a post-processed map of data to help quickly getting component
data for a particular control name - standard name combination.

### Mock generation
The `lib/common/mocks` is a package that is auto-generated via
 [`mockery`](https://github.com/vektra/mockery). Follow the
 instructions there to install `mockery`.
 
 Whenever a modification is made to an existing interface or a new 
 interface is created, you should use mockery to (re)generate the mock
 while inside the `lib/common` folder.

```sh
mockery -name NameOfInterface
# Example:
mockery -name Workspace
```
 
