package component

import (
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
)

// Component struct is an individual component requiring documentation
// Schema info: https://github.com/opencontrol/schemas#component-yaml
type Component struct {
	Name          string                  `yaml:"name" json:"name"`
	Key           string                  `yaml:"key" json:"key"`
	References    *common.GeneralReferences      `yaml:"references" json:"references"`
	Verifications *common.VerificationReferences `yaml:"verifications" json:"verifications"`
	Satisfies     []Satisfies          `yaml:"satisfies" json:"satisfies"`
	SchemaVersion semver.Version          `yaml:"schema_version" json:"schema_version"`
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
	return c.Verifications
}

func (c Component) GetReferences() *common.GeneralReferences {
	return c.References
}

func (c Component) GetAllSatisfies() []base.Satisfies {
	// Have to do manual conversion
	baseSatisfies := make([]base.Satisfies, len(c.Satisfies))
	for idx, value := range c.Satisfies {
		baseSatisfies[idx] = value
	}
	return baseSatisfies
}

// Satisfies struct contains data demonstrating why a specific component meets
// a control
// This struct is a one-to-one mapping of a `satisfies` item in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type Satisfies struct {
	ControlKey  string             `yaml:"control_key" json:"control_key"`
	StandardKey string             `yaml:"standard_key" json:"standard_key"`
	Narrative   Narrative `yaml:"narrative" json:"narrative"`
	CoveredBy   common.CoveredByList      `yaml:"covered_by" json:"covered_by"`
}

func (s Satisfies)GetControlKey() string {
	return s.ControlKey
}

func (s Satisfies)GetStandardKey() string {
	return s.StandardKey
}

func (s Satisfies) GetNarratives() []base.Narrative {
	// Have to do manual conversion
	var baseNarrative []base.Narrative
	if len(s.Narrative) > 0 {
		baseNarrative := make([]base.Narrative, 1)
		baseNarrative[0] = s.Narrative
	}

	return baseNarrative
}

func (s Satisfies) GetCoveredBy() common.CoveredByList {
	return s.CoveredBy
}

type Narrative string

func (n Narrative) GetKey() string {
	return ""
}

func (n Narrative) GetText() string {
	return string(n)
}
