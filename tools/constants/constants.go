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

const (
	// ErrMissingVersion reports that the schema version cannot be found.
	ErrMissingVersion  = "Schema Version can not be found."
	// ErrComponentFileDNE is raised when a component file does not exists
	ErrComponentFileDNE = "Component files does not exist"
)

const (
	// WarningNoInformationAvailable is a warning to indicate that no information was found. Typically this will be
	// used when a template is being filled in and there is no information found for a particular section.
	WarningNoInformationAvailable = "No information available for component"
)