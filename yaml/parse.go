package yaml

import (
	"errors"
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
	"gopkg.in/yaml.v2"
)

var (
	// ErrNoDataToParse represents the case that there is no data to be found to be parsed (either nil or empty).
	ErrNoDataToParse        = errors.New("No data to parse")
	// ErrUnknownSchemaVersion is thrown when the schema version is unknown to the parser.
	ErrUnknownSchemaVersion = errors.New("Unknown schema version")
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
	switch base.SchemaVersion {
	case 1.0:
		schema, parseError = parser.ParseV1_0(data)
	default:
		return nil, ErrUnknownSchemaVersion
	}
	if parseError != nil {
		return nil, parseError
	}

	return schema, nil
}
