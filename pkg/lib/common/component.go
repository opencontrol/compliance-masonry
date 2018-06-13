package common

import "github.com/blang/semver"

//go:generate mockery -name Component
//go:generate mockery -name Satisfies
//go:generate mockery -name Section

// Component is the basic building block for all systems.
// Schema info: https://github.com/opencontrol/schemas#component-yaml
//
// GetName returns the name of the component
//
// GetKey returns the key for the component (may not be unique). Useful for creating directories.
//
// SetKey sets the key for the component. Useful for overriding.
//
// GetAllSatisfies gets all the Satisfies objects for the component.
//
// GetVerifications get all the verifications.
//
// GetReferences get all the references.
//
// GetVersion returns the version
//
// SetVersion sets the version for the component.
//
// GetResponsibleRole gets the responsible party / role for the component.
type Component interface {
	GetName() string
	GetKey() string
	SetKey(string)
	GetAllSatisfies() []Satisfies
	GetVerifications() *VerificationReferences
	GetReferences() *GeneralReferences
	GetVersion() semver.Version
	SetVersion(semver.Version)
	GetResponsibleRole() string
}

// Satisfies contains information regarding how the component satisfies a given standard and control
//
// GetStandardKey returns the standard
//
// GetControlKey returns the control
//
// GetNarratives gets all the general documentation for this particular standard and control
//
// GetParameters gets all the parameters for this particular standard and control
//
// GetCoveredBy gets the list of all the CoveredBy
//
// GetControlOrigin returns the control origin (only the first one if multiple)
//
// GetControlOrigins returns all the control origins
//
// GetImplementationStatus returns the implementation status (only the first one if multiple)
//
// GetImplementationStatuses returns all implementation statuses
type Satisfies interface {
	GetStandardKey() string
	GetControlKey() string
	GetNarratives() []Section
	GetParameters() []Section
	GetCoveredBy() CoveredByList
	GetControlOrigin() string
	GetControlOrigins() []string
	GetImplementationStatus() string
	GetImplementationStatuses() []string
}

// Section is a general holder that allows it to be used in something like a map
//
// GetKey returns a unique key
//
// GetText returns the text for the section
type Section interface {
	GetKey() string
	GetText() string
}
