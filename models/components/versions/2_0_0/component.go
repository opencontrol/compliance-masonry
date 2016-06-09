package component

import (
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"github.com/opencontrol/compliance-masonry/models/common"
)

// Component struct is an individual component requiring documentation
// Schema info: https://github.com/opencontrol/schemas#component-yaml
type Component struct {
	Name          string                  `yaml:"name" json:"name"`
	Key           string                  `yaml:"key" json:"key"`
	References    common.GeneralReferences      `yaml:"references" json:"references"`
	Verifications common.VerificationReferences `yaml:"verifications" json:"verifications"`
	Satisfies     []Satisfies          `yaml:"satisfies" json:"satisfies"`
	SchemaVersion float32                 `yaml:"schema_version" json:"schema_version"`
}

func (c Component) GetName() string {
	return c.Name
}

func (c Component) GetKey() string {
	return c.Key
}

func (c *Component) SetKey(key string) {
	c.Key = key
}

func (c Component) GetVerifications() *common.VerificationReferences {
	return &c.Verifications
}

func (c Component) GetReferences() *common.GeneralReferences {
	return &c.References
}

func (c Component) GetAllSatisfies() []base.Satisfies {
	// Have to do manual conversion from this Component's Satisfies to the interface base.Satisfies.
	baseSatisfies := make([]base.Satisfies, len(c.Satisfies))
	for idx, value := range c.Satisfies {
		baseSatisfies[idx] = value
	}
	return baseSatisfies
}

func (c Component) GetVersion() float32 {
	return c.SchemaVersion
}

// Satisfies struct contains data demonstrating why a specific component meets
// a control
// This struct is a one-to-one mapping of a `satisfies` item in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type Satisfies struct {
	ControlKey  string        `yaml:"control_key" json:"control_key"`
	StandardKey string        `yaml:"standard_key" json:"standard_key"`
	Narrative   string        `yaml:"narrative" json:"narrative"`
	CoveredBy   common.CoveredByList `yaml:"covered_by" json:"covered_by"`
}

func (s Satisfies) GetControlKey() string {
	return s.ControlKey
}

func (s Satisfies) GetStandardKey() string {
	return s.StandardKey
}

func (s Satisfies) GetNarrative() string {
	return s.Narrative
}

func (s Satisfies) GetCoveredBy() common.CoveredByList {
	return s.CoveredBy
}
