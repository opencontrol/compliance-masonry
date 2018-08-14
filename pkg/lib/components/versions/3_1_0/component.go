/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package component

import (
	"sort"

	"github.com/blang/semver"
	"github.com/fatih/set"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
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
	return c.ResponsibleRole
}

// Satisfies struct contains data demonstrating why a specific component meets
// a control
// This struct is a one-to-one mapping of a `satisfies` item in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type Satisfies struct {
	ControlKey             string               `yaml:"control_key" json:"control_key"`
	StandardKey            string               `yaml:"standard_key" json:"standard_key"`
	Narrative              []NarrativeSection   `yaml:"narrative" json:"narrative"`
	CoveredBy              common.CoveredByList `yaml:"covered_by" json:"covered_by"`
	Parameters             []Section            `yaml:"parameters" json:"parameters"`
	ControlOrigin          string               `yaml:"control_origin" json:"control_origin"`
	ControlOrigins         []string             `yaml:"control_origins" json:"control_origins"`
	ImplementationStatus   string               `yaml:"implementation_status" json:"implementation_status"`
	ImplementationStatuses []string             `yaml:"implementation_statuses" json:"implementation_statuses"`
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
	// Have to do manual conversion to the interface base.Section from NarrativeSection.
	baseSection := make([]common.Section, len(s.Narrative))
	for idx, value := range s.Narrative {
		baseSection[idx] = value
	}
	return baseSection
}

// GetParameters gets all the parameters for this particular standard and control
func (s Satisfies) GetParameters() []common.Section {
	// Have to do manual conversion to the interface base.Section from Section.
	baseSection := make([]common.Section, len(s.Parameters))
	for idx, value := range s.Parameters {
		baseSection[idx] = value
	}
	return baseSection
}

// GetCoveredBy gets the list of all the CoveredBy
func (s Satisfies) GetCoveredBy() common.CoveredByList {
	return s.CoveredBy
}

// GetControlOrigin returns the control origin (only the first one if multiple)
func (s Satisfies) GetControlOrigin() string {
	return s.ControlOrigin
}

// GetControlOrigins returns all the control origins
func (s Satisfies) GetControlOrigins() []string {
	controlOrigins := set.New(set.ThreadSafe)
	for i := range s.ControlOrigins {
		controlOrigins.Add(s.ControlOrigins[i])
	}
	if s.ControlOrigin != "" {
		controlOrigins.Add(s.ControlOrigin)
	}
	l := set.StringSlice(controlOrigins)
	sort.Strings(l)
	return l
}

// GetImplementationStatus returns the implementation status (only the first one if multiple)
func (s Satisfies) GetImplementationStatus() string {
	return s.ImplementationStatus
}

// GetImplementationStatuses returns all implementation statuses
func (s Satisfies) GetImplementationStatuses() []string {
	implementationStatuses := set.New(set.ThreadSafe)
	for i := range s.ImplementationStatuses {
		implementationStatuses.Add(s.ImplementationStatuses[i])
	}
	if s.ImplementationStatus != "" {
		implementationStatuses.Add(s.ImplementationStatus)
	}
	l := set.StringSlice(implementationStatuses)
	sort.Strings(l)
	return l
}

// NarrativeSection contains the key and text for a particular section.
// NarrativeSection can omit the key.
type NarrativeSection struct {
	Key  string `yaml:"key,omitempty" json:"key,omitempty"`
	Text string `yaml:"text" json:"text"`
}

// GetKey returns a unique key
func (ns NarrativeSection) GetKey() string {
	return ns.Key
}

// GetText returns the text for the section
func (ns NarrativeSection) GetText() string {
	return ns.Text
}

// Section contains the key and text for a particular section. Both are required.
type Section struct {
	Key  string `yaml:"key" json:"key"`
	Text string `yaml:"text" json:"text"`
}

// GetKey returns a unique key
func (s Section) GetKey() string {
	return s.Key
}

// GetText returns the text for the section
func (s Section) GetText() string {
	return s.Text
}
