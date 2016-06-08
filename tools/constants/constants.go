package constants

import "errors"

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
	// ErrComponentSchemaParsef is a formatted string for reporting which version schema to check.
	ErrComponentSchemaParsef = "Unable to parse component. Please check component.yaml schema for version %v"
)

var (
	// ErrMissingVersion reports that the schema version cannot be found.
	ErrMissingVersion            = errors.New("Schema Version can not be found.")
	// ErrComponentFileDNE is raised when a component file does not exists
	ErrComponentFileDNE = errors.New("Component files does not exist")
	// ErrControlSchema is raised a control cannot be parsed
	ErrComponentSchema = errors.New("Unable to parse component")
)