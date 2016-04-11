# opencontrol.yaml README

opencontrol.yaml is a configuration file that describes the makeup of a system.
It describes the components it directly uses as well as any dependent systems / components.

In order to support backwards comparability, we version the configuration file with a field named: `schema_version`.

The version is expressed in a string format using [semantic versioning](http://semver.org/).

## Adding a new schema version.

When adding a new schema version. (in this example we will add v2.0.0
- Create a new numbered version folder under the `versions` folder.
  - e.g. `mkdir versions/2.0.0/`.
  - Make the package name for the go files `schema`.
  - Ensure there is a function that implements this prototype: `func (s *Schema) Parse(data []byte) error`
- In `parse.go`
  - Create a new Schema Version variable for clean code sake.
    - e.g. `SchemaV2_0_0 = semver.Version{2, 0, 0, nil, nil}`
  - In the switch case, add a new case for 2.0.0 that calls ParseV2_0_0
  
  ```
  	case SchemaV2_0_0.Equals(v):
  		schema, parseError = parser.ParseV2_0_0(data)
  ```

- Add a new `ParseV2_0_0` function to the `parser/parser.go` file.
  - It should satisfy the `func (p Parser) ParseVX_Y_Z(data []byte) (common.BaseSchema, error) {` prototype.
  - Alias the import so that you can easily refer to the particular version.
  ```
  import (
    v1_0_0 "github.com/opencontrol/compliance-masonry/yaml/versions/1.0.0"
    v2_0_0 "github.com/opencontrol/compliance-masonry/yaml/versions/2.0.0"
  )
  ```
  - Now it should call the parse function that your schema version has.
  
  ```
  config := v2_0_0.Schema{}
  config.Parse(data)
  ```