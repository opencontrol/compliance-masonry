package opencontrol

import (
	"errors"
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol/versions/base"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"gopkg.in/yaml.v2"
)

var (
	// SchemaV1_0_0 is the semantic versioning representation in object form for version 1.0.0
	SchemaV1_0_0 = semver.Version{1, 0, 0, nil, nil}
)

const (
	// ErrMalformedBaseYamlPrefix is just the prefix to the error message for when the program is unable to parse
	// data into the base yaml struct.
	ErrMalformedBaseYamlPrefix = "Unable to parse yaml data"
)

// Parse will try to parse the data and determine which specific version of schema to further parse.
func Parse(parser base.SchemaParser, data []byte) (base.OpenControl, error) {
	if data == nil || len(data) == 0 {
		return nil, common.ErrNoDataToParse
	}
	b := base.Base{}
	err := yaml.Unmarshal(data, &b)
	if err != nil {
		return nil, errors.New(ErrMalformedBaseYamlPrefix + " - " + err.Error())
	}

	var opencontrol base.OpenControl
	var parseError error
	v, err := semver.Parse(b.SchemaVersion)
	if err != nil {
		return nil, common.ErrCantParseSemver
	}
	switch {
	case SchemaV1_0_0.Equals(v):
		opencontrol, parseError = parser.ParseV1_0_0(data)
	default:
		return nil, common.ErrUnknownSchemaVersion
	}
	if parseError != nil {
		return nil, parseError
	}

	return opencontrol, nil
}
