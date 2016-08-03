# Lib Package

This package (and the sub packages) are the only packages that should be imported by external projects. This lib package is very helpful for those writing plugins to extend the functionality of Masonry.

## `Workspace`
Workspace is the starting point for all the information available to plugins. For those using Go plugins, using the interface methods are all you need. It represents all the information collected into the `opencontrols` folder after running the `get` command. 

### YAML files
The yaml files are represented by:

- `Components`
- `Certifications`
- `Standards`

### Result Data
`Justifications` is a data structure that is not represented in yaml but rather a post-processed map of data to help quickly getting component data for a particular control name - standard name combination.
 

## Plugin Developer Guide
Developers should not have to worry about loading real data / workspaces for their unit tests (they should for integration tests).

There is an example of developing your Go plugin and tests in `exampleplugin/example.go` and `exampleplugin/example_test.go` respectively.