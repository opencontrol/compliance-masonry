package versions

import (
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/config"
	v2 "github.com/opencontrol/compliance-masonry/models/components/versions/2_0_0"
	v3 "github.com/opencontrol/compliance-masonry/models/components/versions/3_0_0"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"gopkg.in/yaml.v2"
)

var (
	ComponentV2_0_0 = semver.MustParse("2.0.0")
	ComponentV3_0_0 = semver.MustParse("3.0.0")
)

func ParseComponent(componentData []byte) (base.Component, error) {
	b := base.Base{}
	err := yaml.Unmarshal(componentData, &b)
	if err != nil {
		return nil, errors.New(constants.ErrMissingVersion)
	}
	var component base.Component
	switch {
	case ComponentV2_0_0.Equals(b.SchemaVersion):
		c := new(v2.Component)
		err = yaml.Unmarshal(componentData, c)
		component = c
	case ComponentV3_0_0.Equals(b.SchemaVersion):
		c := new(v3.Component)
		err = yaml.Unmarshal(componentData, c)
		component = c
	default:
		return nil, config.ErrUnknownSchemaVersion
	}
	if err != nil {
		return nil, fmt.Errorf(constants.ErrComponentSchemaParsef, b.SchemaVersion.String())
	}
	return component, nil
}
