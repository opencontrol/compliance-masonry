package opencontrol

import (
	"errors"
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	v1_0_0 "github.com/opencontrol/compliance-masonry/pkg/lib/opencontrol/versions/1.0.0"
	"gopkg.in/yaml.v2"
)

var (
	// SchemaV1_0_0 is the semantic versioning representation in object form for version 1.0.0
	SchemaV1_0_0 = semver.Version{Major: 1, Minor: 0, Patch: 0, Pre: nil, Build: nil}
)

const (
	// ErrMalformedBaseYamlPrefix is just the prefix to the error message for when the program is unable to parse
	// data into the base yaml struct.
	ErrMalformedBaseYamlPrefix = "Unable to parse yaml data"
)

// YAMLParser is the concrete implementation of parsing different schema versions in YAML format.
type YAMLParser struct{}

// Parse will try to parse the data and determine which specific version of schema to further parse.
func (parser YAMLParser) Parse(data []byte) (common.OpenControl, error) {
	if data == nil || len(data) == 0 {
		return nil, common.ErrNoDataToParse
	}
	b := Base{}
	err := yaml.Unmarshal(data, &b)
	if err != nil {
		return nil, errors.New(ErrMalformedBaseYamlPrefix + " - " + err.Error())
	}

	var opencontrol common.OpenControl
	var parseError error
	v, err := semver.Parse(b.SchemaVersion)
	if err != nil {
		return nil, common.ErrCantParseSemver
	}
	switch {
	case SchemaV1_0_0.Equals(v):
		opencontrol = new(v1_0_0.OpenControl)
		parseError = yaml.Unmarshal(data, opencontrol)
	default:
		return nil, common.ErrUnknownSchemaVersion
	}
	if parseError != nil {
		return nil, parseError
	}

	return opencontrol, nil
}
