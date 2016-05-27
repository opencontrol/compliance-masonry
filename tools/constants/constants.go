package constants

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

// DefaultFloat32Value is the exported constant value for the default float32 value for an uninitialized float32
// variable.
const (
	DefaultFloat32Value float32 = iota
)

const (
	// VersionNotNeeded is a place holder to indicate that the schema version does not need to be specified.
	VersionNotNeeded float32 = -18.0
	// MinComponentYAMLVersion is the minimum schema version for the component
	// YAML supported by this masonry toolchain.
	MinComponentYAMLVersion float32 = 3.0
	// MaxComponentYAMLVersion is the minimum schema version for the component
	// YAML supported by this masonry toolchain.
	MaxComponentYAMLVersion float32 = VersionNotNeeded
)

const (
	// WarningNoInformationFound is a warning to indicate that no information was found. Typically this will be
	// used when a template is being filled in and there is no information found for a particular section.
	WarningNoInformationFound = "No information found"
)