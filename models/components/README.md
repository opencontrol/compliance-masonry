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

*TODO: Future patches will enable 1) more component.yaml and 2) using multiple versions. More details to come then.*

For details on the component scheams, consult [the schemas repository](https://github.com/opencontrol/schemas)

## tl;dr
All component versions must implement the [interface](versions/base/component.go).
