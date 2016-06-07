package constants

import (
	"errors"
)

const (
	// DefaultStandardsFolder is the folder where to store standards.
	DefaultStandardsFolder = "standards"
	// DefaultCertificationsFolder is the folder where to store certifications.
	DefaultCertificationsFolder = "certifications"
	// DefaultComponentsFolder is the folder where to store components.
	DefaultComponentsFolder = "components"
	// DefaultDestination is the root folder where to store standards, certifications, and components.
	DefaultDestination = "opencontrols"
	// DefaultConfigYaml is the file name for the file to find config details
	DefaultConfigYaml = "opencontrol.yaml"
)

// ResourceType is a type to help tell when it should be of only types of resources.
type ResourceType string

const (
	// Standards is the placeholder for the resource type of standards
	Standards ResourceType = "Standards"
	// Certifications is the placeholder for the resource type of certifications
	Certifications ResourceType = "Certifications"
	// Components is the placeholder for the resource type of components
	Components ResourceType = "Components"
)

const (
	ErrVersionNotInSemverFormatf = "Version %v is not in semver format"
	ErrMissingVersion            = "Schema Version can not be found."
)

const (
	// WarningNoInformationAvailable is a warning to indicate that no information was found. Typically this will be
	// used when a template is being filled in and there is no information found for a particular section.
	WarningNoInformationAvailable = "No information available for component"

	// WarningUnknownStandardAndControlf is a string with two string specifiers to allow for injecting both
	// the standard and the control into the final string. This is a warning that is used for when the given
	// standard and control return no information at all. Which implies no component in the whole system has
	// addressed this combination of standard and control.
	WarningUnknownStandardAndControlf = "No information found for the combination of standard %s and control %s"
)

var (
	// ErrComponentFileDNE is raised when a component file does not exists
	ErrComponentFileDNE = errors.New("Component files does not exist")
)

const (
	// ErrComponentSchemaParsef is a formatted string for reporting which version schema to check.
	ErrComponentSchemaParsef = "Unable to parse component. Please check component.yaml schema for version %s"
)
