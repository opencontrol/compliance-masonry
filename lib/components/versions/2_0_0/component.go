package component

import (
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

// Component struct is an individual component requiring documentation
// Schema info: https://github.com/opencontrol/schemas#component-yaml
type Component struct {
	Name          string                        `yaml:"name" json:"name"`
	Key           string                        `yaml:"key" json:"key"`
	References    common.GeneralReferences      `yaml:"references" json:"references"`
	Verifications common.VerificationReferences `yaml:"verifications" json:"verifications"`
	Satisfies     []Satisfies                   `yaml:"satisfies" json:"satisfies"`
	SchemaVersion semver.Version                `yaml:"-" json:"-"`
}

// GetName returns the name of the component
func (c Component) GetName() string {
	return c.Name
}

// GetKey returns the key for the component (may not be unique). Useful for creating directories.
func (c Component) GetKey() string {
	return c.Key
}

// SetKey sets the key for the component. Useful for overriding.
func (c *Component) SetKey(key string) {
	c.Key = key
}

// GetVerifications get all the verifications.
func (c Component) GetVerifications() *common.VerificationReferences {
	return &c.Verifications
}

// GetReferences get all the references.
func (c Component) GetReferences() *common.GeneralReferences {
	return &c.References
}

// GetAllSatisfies gets all the Satisfies objects for the component.
func (c Component) GetAllSatisfies() []common.Satisfies {
	// Have to do manual conversion from this Component's Satisfies to the interface base.Satisfies.
	baseSatisfies := make([]common.Satisfies, len(c.Satisfies))
	for idx, value := range c.Satisfies {
		baseSatisfies[idx] = value
	}
	return baseSatisfies
}

// GetVersion returns the version
func (c Component) GetVersion() semver.Version {
	return c.SchemaVersion
}

// SetVersion sets the version for the component.
func (c *Component) SetVersion(version semver.Version) {
	c.SchemaVersion = version
}

// GetResponsibleRole gets the responsible party / role for the component.
func (c Component) GetResponsibleRole() string {
	return ""
}

// Satisfies struct contains data demonstrating why a specific component meets
// a control
// This struct is a one-to-one mapping of a `satisfies` item in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type Satisfies struct {
	ControlKey           string               `yaml:"control_key" json:"control_key"`
	StandardKey          string               `yaml:"standard_key" json:"standard_key"`
	Narrative            Narrative            `yaml:"narrative" json:"narrative"`
	CoveredBy            common.CoveredByList `yaml:"covered_by" json:"covered_by"`
	ImplementationStatus string               `yaml:"implementation_status" json:"implementation_status"`
}

// GetControlKey returns the control
func (s Satisfies) GetControlKey() string {
	return s.ControlKey
}

// GetStandardKey returns the standard
func (s Satisfies) GetStandardKey() string {
	return s.StandardKey
}

// GetNarratives gets all the general documentation for this particular standard and control
func (s Satisfies) GetNarratives() []common.Section {
	// Have to do manual conversion to the interface base.Section.
	// V2.0.0 only had one Narrative field, so if it actually exists, let's create a slice of 1 to return.
	var baseNarrative []common.Section
	if len(s.Narrative) > 0 {
		baseNarrative = make([]common.Section, 1)
		baseNarrative[0] = s.Narrative
	}

	return baseNarrative
}

// GetParameters gets all the parameters for this particular standard and control
func (s Satisfies) GetParameters() []common.Section {
	return nil
}

// GetCoveredBy gets the list of all the CoveredBy
func (s Satisfies) GetCoveredBy() common.CoveredByList {
	return s.CoveredBy
}

// GetControlOrigin returns the control origin (empty string for this version)
func (s Satisfies) GetControlOrigin() string {
	return ""
}

// GetControlOrigins returns all the control origins (empty slice for this version)
func (s Satisfies) GetControlOrigins() []string {
	return []string{}
}

// Narrative is the representation of the general documentation for a particular standard and control for the component.
type Narrative string

// GetKey returns a unique key (empty string for this version)
func (n Narrative) GetKey() string {
	return ""
}

// GetText returns the text for the section
func (n Narrative) GetText() string {
	return string(n)
}

// GetImplementationStatus returns the implementation status
func (s Satisfies) GetImplementationStatus() string {
	return s.ImplementationStatus
}

// GetImplementationStatuses returns all implementation statuses (just the only one for this version)
func (s Satisfies) GetImplementationStatuses() []string {
	return []string{s.ImplementationStatus}
}
