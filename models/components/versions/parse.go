package versions

import (
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"gopkg.in/yaml.v2"
	v2 "github.com/opencontrol/compliance-masonry/models/components/versions/2_0_0"
	"github.com/opencontrol/compliance-masonry/config"
	"fmt"
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"errors"
)

var (
	ComponentV2_0_0 = semver.MustParse("2.0.0")
)

func ParseComponent(componentData []byte) (base.Component, error) {
	b := base.Base{}
	err := yaml.Unmarshal(componentData, &b)
	if err != nil {
		// If we have a human friendly BaseComponentParseError, return it.
		switch err.(type) {
		case base.BaseComponentParseError:
			return nil, err
		}
		// Otherwise, just return a generic error about the schema.
		return nil, errors.New(constants.ErrComponentSchema)
	}
	var component base.Component
	switch {
	case ComponentV2_0_0.EQ(b.SchemaVersion):
		c := new(v2.Component)
		err = yaml.Unmarshal(componentData, c)
		component = c
	default:
		return nil, config.ErrUnknownSchemaVersion

	}
	if err != nil {
		return nil, fmt.Errorf("Unable to parse component. Please check component.yaml schema for version %s", b.SchemaVersion.String())
	}
	// Copy version from base because some versions of the component can not expect to parse directly into it's own struct
	// e.g. version 2.0.0 with 2.0 float
	component.SetVersion(b.SchemaVersion)
	return component, nil
}