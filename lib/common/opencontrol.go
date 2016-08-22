package common

// OpenControl is an interface that every OpenControl yaml version should implement.
// Schema info: https://github.com/opencontrol/schemas#opencontrolyaml
//
// GetCertifications retrieves the list of certifications
//
// GetStandards retrieves the list of standards
//
// GetComponents retrieves the list of components
//
// GetCertificationsDependencies retrieves the list of certifications that this config will inherit.
//
// GetStandardsDependencies retrieves the list of standards that this config will inherit.
//
// GetComponentsDependencies retrieves the list of components / systems that this config will inherit.
type OpenControl interface {
	GetCertifications() []string
	GetStandards() []string
	GetComponents() []string
	GetCertificationsDependencies() []RemoteSource
	GetStandardsDependencies() []RemoteSource
	GetComponentsDependencies() []RemoteSource
}

// RemoteSource is an interface that any remote sources should implement in order to know how to download them.
//
// GetURL returns the URL of the resource.
//
// GetRevision returns the specific revision of the resource.
//
// GetConfigFile returns the config file to look at once the resource is downloaded.
type RemoteSource interface {
	GetURL() string
	GetRevision() string
	GetConfigFile() string
}
