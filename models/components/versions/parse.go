package versions

import (
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"gopkg.in/yaml.v2"
	v2 "github.com/opencontrol/compliance-masonry/models/components/versions/2_0_0"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/config"
	"fmt"
)

const (
	ComponentV2_0 = 2.0
)

func ParseComponent(componentData []byte) (base.Component, error) {
	b := base.Base{}
	err := yaml.Unmarshal(componentData, &b)
	if err != nil {
		return nil, constants.ErrMissingVersion
	}
	var component base.Component
	switch b.SchemaVersion {
	case ComponentV2_0:
		c := new(v2.Component)
		err = yaml.Unmarshal(componentData, c)
		component = c
	default:
		return nil, config.ErrUnknownSchemaVersion

	}
	if err != nil {
		return nil, fmt.Errorf(constants.ErrComponentSchemaParsef, b.SchemaVersion)
	}
	return component, nil
}