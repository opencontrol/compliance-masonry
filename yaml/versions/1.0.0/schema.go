package schema

import (
	"errors"
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	// ErrMalformedV1_0_0YamlPrefix is just the prefix to the error message for when the program is unable to parse
	// data into the v1.0.0 yaml struct.
	ErrMalformedV1_0_0YamlPrefix = "Unable to parse yaml data"
)

// Schema contains the structs for the v1.0.0 schema
type Schema struct {
	common.Base    `yaml:",inline"`
	Meta           Metadata     `yaml:"metadata"`
	Name           string       `yaml:"name"`
	Components     []string     `yaml:",flow"`
	Certifications []string     `yaml:",flow"`
	Standards      []string     `yaml:",flow"`
	Dependencies   Dependencies `yaml:"dependencies"`
}

// Dependencies contains all the dependencies for the system
type Dependencies struct {
	Certifications []common.Entry `yaml:"certifications"`
	Systems        []common.Entry `yaml:",flow"`
	Standards      []common.Entry `yaml:",flow"`
}

// Metadata contains metadata about the system.
type Metadata struct {
	Description string   `yaml:"description"`
	Maintainers []string `yaml:",flow"`
}

// Parse will parse using it's own schema. In this case the v1.0.0 schema.
func (s *Schema) Parse(data []byte) error {
	err := yaml.Unmarshal(data, s)
	if err != nil {
		return errors.New(ErrMalformedV1_0_0YamlPrefix + " - " + err.Error())
	}

	return nil
}

func (s *Schema) GetResources(destination string, downloader common.EntryDownloader) error {
	tempResourcesDir, err := ioutil.TempDir("", "opencontrol-resources")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempResourcesDir)

	// Get Certifications
	for _, certification := range s.Dependencies.Certifications {
		err := downloader.DownloadEntry(certification, destination)
		if err != nil {
			return err
		}
	}

	// Get Standards
	for _, standard := range s.Dependencies.Standards {
		err := downloader.DownloadEntry(standard, destination)
		if err != nil {
			return err
		}
	}

	// Get Systems
	for _, system := range s.Dependencies.Systems {
		err := downloader.DownloadEntry(system, destination)
		if err != nil {
			return err
		}
	}

	return nil
}
