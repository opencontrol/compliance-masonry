package schema

import (
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

const (
	// ErrMalformedV1_0_0YamlPrefix is just the prefix to the error message for when the program is unable to parse
	// data into the v1.0.0 yaml struct.
	ErrMalformedV1_0_0YamlPrefix = "Unable to parse yaml data"
)

// OpenControl contains the structs for the v1.0.0 schema
type OpenControl struct {
	Meta           Metadata     `yaml:"metadata"`
	Name           string       `yaml:"name"`
	Components     []string     `yaml:",flow"`
	Certifications []string     `yaml:",flow"`
	Standards      []string     `yaml:",flow"`
	Dependencies   Dependencies `yaml:"dependencies"`
}

// Dependencies contains all the dependencies for the system
type Dependencies struct {
	Certifications []VCSEntry `yaml:"certifications"`
	Systems        []VCSEntry `yaml:",flow"`
	Standards      []VCSEntry `yaml:",flow"`
}

// Metadata contains metadata about the system.
type Metadata struct {
	Description string   `yaml:"description"`
	Maintainers []string `yaml:",flow"`
}

// VCSEntry is a generic holder for handling the specific location and revision of a resource.
type VCSEntry struct {
	URL      string `yaml:"url"`
	Revision string `yaml:"revision"`
	ContextDir string `yaml:"contextdir"`
	Path     string `yaml:"path"`
}

// GetCertifications retrieves the list of certifications
func (o OpenControl) GetCertifications() []string {
	return o.Certifications
}

// GetComponents retrieves the list of components
func (o OpenControl) GetComponents() []string {
	return o.Components
}

// GetStandards retrieves the list of standards
func (o OpenControl) GetStandards() []string {
	return o.Standards
}

// GetCertificationsDependencies retrieves the list of certifications that this config will inherit.
func (o OpenControl) GetCertificationsDependencies() []common.RemoteSource {
	// Have to do manual conversion to the interface common.RemoteSource from VCSEntry.
	entries := make([]common.RemoteSource, len(o.Dependencies.Certifications))
	for idx, value := range o.Dependencies.Certifications {
		entries[idx] = value
	}
	return entries
}

// GetComponentsDependencies retrieves the list of components / systems that this config will inherit.
func (o OpenControl) GetComponentsDependencies() []common.RemoteSource {
	// Have to do manual conversion to the interface common.RemoteSource from VCSEntry.
	entries := make([]common.RemoteSource, len(o.Dependencies.Systems))
	for idx, value := range o.Dependencies.Systems {
		entries[idx] = value
	}
	return entries
}

// GetStandardsDependencies retrieves the list of standards that this config will inherit.
func (o OpenControl) GetStandardsDependencies() []common.RemoteSource {
	// Have to do manual conversion to the interface common.RemoteSource from VCSEntry.
	entries := make([]common.RemoteSource, len(o.Dependencies.Standards))
	for idx, value := range o.Dependencies.Standards {
		entries[idx] = value
	}
	return entries
}

// GetConfigFile is a getter for the config file name. Will return DefaultConfigYaml value if none has been set.
func (e VCSEntry) GetConfigFile() string {
	if e.Path == "" {
		return constants.DefaultConfigYaml
	}
	return e.Path
}

// GetRevision returns the specific revision of the vcs resource.
func (e VCSEntry) GetRevision() string {
	return e.Revision
}

// GetURL returns the URL of the vcs resource.
func (e VCSEntry) GetURL() string {
	return e.URL
}

// GetContextDir returns the dir containing content in the vcs resource.
func (e VCSEntry) GetContextDir() string {
        return e.ContextDir
}
