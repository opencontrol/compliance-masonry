package versions

import (
	"github.com/blang/semver"
	v2 "github.com/opencontrol/compliance-masonry/models/components/versions/2_0_0"
	v3 "github.com/opencontrol/compliance-masonry/models/components/versions/3_0_0"
	"gopkg.in/yaml.v2"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"github.com/opencontrol/compliance-masonry/config"
	"errors"
)

var (
	ComponentV2_0_0 = semver.MustParse("2.0.0")
	ComponentV3_0_0 = semver.MustParse("3.0.0")
)

func ParseComponent(componentData []byte) (base.Component, error) {
	b := base.Base{}
	err := yaml.Unmarshal(componentData, &b)
	if err != nil {
		return nil, errors.New(config.ErrMalformedBaseYamlPrefix + " - " + err.Error())
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
		return nil, err
	}
	return component, nil
}
