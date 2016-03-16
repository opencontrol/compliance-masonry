package yaml

import (
	"errors"
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
	"gopkg.in/yaml.v2"
	"github.com/blang/semver"
)

var (
	// ErrNoDataToParse represents the case that there is no data to be found to be parsed (either nil or empty).
	ErrNoDataToParse        = errors.New("No data to parse")
	// ErrUnknownSchemaVersion is thrown when the schema version is unknown to the parser.
	ErrUnknownSchemaVersion = errors.New("Unknown schema version")
	// ErrCantParseSemver is thrown when the semantic versioning can not be parsed.
	ErrCantParseSemver = errors.New("Can't parse semantic versioning of schema_version")
)

var (
	SchemaV1_0_0 = semver.Version{1, 0, 0, nil, nil}
)

// Parse will try to parse the data and determine which specific version of schema to further parse.
func Parse(parser common.SchemaParser, data []byte) (common.BaseSchema, error) {
	if data == nil || len(data) == 0 {
		return nil, ErrNoDataToParse
	}
	base := common.Base{}
	err := yaml.Unmarshal(data, &base)
	if err != nil {
		return nil, err
	}

	var schema common.BaseSchema
	var parseError error
	v, err := semver.Parse(base.SchemaVersion)
	if err != nil {
		return nil, ErrCantParseSemver
	}
	switch {
	case SchemaV1_0_0.Equals(v):
		schema, parseError = parser.ParseV1_0_0(data)
	default:
		return nil, ErrUnknownSchemaVersion
	}
	if parseError != nil {
		return nil, parseError
	}

	return schema, nil
}
