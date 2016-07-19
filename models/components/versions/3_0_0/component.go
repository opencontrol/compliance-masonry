package component

import (
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
)

// Component struct is an individual component requiring documentation
// Schema info: https://github.com/opencontrol/schemas#component-yaml
type Component struct {
	Name            string                        `yaml:"name" json:"name"`
	Key             string                        `yaml:"key" json:"key"`
	References      common.GeneralReferences      `yaml:"references" json:"references"`
	Verifications   common.VerificationReferences `yaml:"verifications" json:"verifications"`
	Satisfies       []Satisfies                   `yaml:"satisfies" json:"satisfies"`
	ResponsibleRole string                        `yaml:"responsible_role" json:"responsible_role"`
	SchemaVersion   semver.Version                `yaml:"-" json:"-"`
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

func (c Component) GetVersion() semver.Version {
	return c.SchemaVersion
}

func (c *Component) SetVersion(version semver.Version) {
	c.SchemaVersion = version
}

func (c Component) GetResponsibleRole() string {
	return c.ResponsibleRole
}

// Satisfies struct contains data demonstrating why a specific component meets
// a control
// This struct is a one-to-one mapping of a `satisfies` item in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type Satisfies struct {
	ControlKey           string               `yaml:"control_key" json:"control_key"`
	StandardKey          string               `yaml:"standard_key" json:"standard_key"`
	Narrative            []NarrativeSection   `yaml:"narrative" json:"narrative"`
	CoveredBy            common.CoveredByList `yaml:"covered_by" json:"covered_by"`
	Parameters           []Section            `yaml:"parameters" json:"parameters"`
	ControlOrigin        string               `yaml:"control_origin" json:"control_origin"`
	ImplementationStatus string               `yaml:"implementation_status" json:"implementation_status"`
}

func (s Satisfies) GetControlKey() string {
	return s.ControlKey
}

func (s Satisfies) GetStandardKey() string {
	return s.StandardKey
}

func (s Satisfies) GetNarratives() []base.Section {
	// Have to do manual conversion to the interface base.Section from NarrativeSection.
	baseSection := make([]base.Section, len(s.Narrative))
	for idx, value := range s.Narrative {
		baseSection[idx] = value
	}
	return baseSection
}

func (s Satisfies) GetParameters() []base.Section {
	// Have to do manual conversion to the interface base.Section from Section.
	baseSection := make([]base.Section, len(s.Parameters))
	for idx, value := range s.Parameters {
		baseSection[idx] = value
	}
	return baseSection
}

func (s Satisfies) GetCoveredBy() common.CoveredByList {
	return s.CoveredBy
}

func (s Satisfies) GetControlOrigin() string {
	return s.ControlOrigin
}

func (s Satisfies) GetControlOrigins() []string {
	return []string{s.ControlOrigin}
}

func (s Satisfies) GetImplementationStatus() string {
	return s.ImplementationStatus
}

func (s Satisfies) GetImplementationStatuses() []string {
	return []string{}
}

// NarrativeSection contains the key and text for a particular section.
// NarrativeSection can omit the key.
type NarrativeSection struct {
	Key  string `yaml:"key,omitempty" json:"key,omitempty"`
	Text string `yaml:"text" json:"text"`
}

func (ns NarrativeSection) GetKey() string {
	return ns.Key
}

func (ns NarrativeSection) GetText() string {
	return ns.Text
}

// Section contains the key and text for a particular section. Both are required.
type Section struct {
	Key  string `yaml:"key" json:"key"`
	Text string `yaml:"text" json:"text"`
}

func (s Section) GetKey() string {
	return s.Key
}

func (s Section) GetText() string {
	return s.Text
}
