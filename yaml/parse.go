package yaml

import (
	"gopkg.in/yaml.v2"
	"errors"
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
)

var (
	ErrNoDataToParse = errors.New("No data to parse")
	ErrUnknownSchemaVersion = errors.New("Unknown schema version")
)

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
	switch base.GetSchemaVersion() {
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