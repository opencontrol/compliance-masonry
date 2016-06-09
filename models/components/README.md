# Component Development

## Overview
Components will evolve over time to incorporate more details and in different ways.
However, since this is an actively used tool, it should support all valid component schemas, past and present.
The way this is accomplished is by abstracting access to specific versions of components to generic interfaces located in [versions/base/component.go](versions/base/component.go).

## Development
When creating your component struct, place the appropriate YAML tags on the fields as you would normally for the specific version.
Now, whenever you unmarshal your data, it will fill all the fields you want for that version.

Now you must implement the [interfaces](versions/base/component.go) inside your component structs.

In the case that you want to change how much access you want to give externally, you will need to edit the methods on the [`interfaces`](versions/base/component.go).

### Adding A New Version

1. Create a folder `versions/X_Y_Z` where the `X_Y_Z` is the version in semver notation.
1. Create your Component struct and any other structs that will implement the different interfaces inside [here](versions/base/component.go).
1. Add another variable to represent your version inside of [`versions/parse.go`](versions/parse.go)
1. Add a check in the switch-case block for your new version that is represented by the variable created in step 3.
    1. Follow the same logic seen in the other versions inside the switch-case block.
1. Add tests case fixtures with valid and invalid data for your version along with the other fixtures.
1. Add those cases to the [`versions/parse_test.go`](versions/parse_test.go)


### Editing The Interface

1. Change the interface(s) inside of [interfaces](versions/base/component.go).
1. Change the implementations for all the component versions (located in `versions`) to accomodate the interface changes.

For details on the component scheams, consult [the schemas repository](https://github.com/opencontrol/schemas)

## tl;dr
All component versions must implement the [interface](versions/base/component.go).
