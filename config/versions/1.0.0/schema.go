package schema

import (
	"errors"
	"github.com/go-utils/ufs"
	"github.com/opencontrol/compliance-masonry-go/config"
	"github.com/opencontrol/compliance-masonry-go/config/common"
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
	"github.com/opencontrol/compliance-masonry-go/tools/fs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

func (s *Schema) GetResources(destination string, worker common.ConfigWorker) error {
	// Local
	// Get Certifications
	for _, certification := range s.Certifications {
		err := ufs.CopyFile(certification,
			filepath.Join(destination, constants.DefaultCertificationsFolder, filepath.Base(certification)))
		if err != nil {
			return err
		}
	}

	// Get Standards
	for _, standard := range s.Standards {
		err := ufs.CopyFile(standard,
			filepath.Join(destination, constants.DefaultStandardsFolder, filepath.Base(standard)))
		if err != nil {
			return err
		}
	}

	// Get Components
	for _, component := range s.Components {
		err := ufs.CopyAll(component,
			filepath.Join(destination, constants.DefaultComponentsFolder, filepath.Base(component)), nil)
		if err != nil {
			return err
		}
	}
	// Remote
	tempResourcesDir, err := ioutil.TempDir("", "opencontrol-resources")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempResourcesDir)

	// Get Certifications
	log.Println("Retrieving dependent certifications")
	for _, certification := range s.Dependencies.Certifications {
		tempPath := filepath.Join(tempResourcesDir, constants.DefaultCertificationsFolder, filepath.Base(certification.URL))
		s.getResource(destination, tempPath, worker, certification)
	}

	// Get Standards
	log.Println("Retrieving dependent standards")
	for _, standard := range s.Dependencies.Standards {
		tempPath := filepath.Join(tempResourcesDir, constants.DefaultStandardsFolder, filepath.Base(standard.URL))
		s.getResource(destination, tempPath, worker, standard)
	}

	// Get Systems
	log.Println("Retrieving dependent systems")
	for _, system := range s.Dependencies.Systems {
		tempPath := filepath.Join(tempResourcesDir, constants.DefaultComponentsFolder, filepath.Base(system.URL))
		s.getResource(destination, tempPath, worker, system)
	}

	return nil
}

func (s *Schema) getResource(destination string, tempResourcesDir string, worker common.ConfigWorker, entry common.Entry) error {
	tempPath := filepath.Join(tempResourcesDir, constants.DefaultCertificationsFolder, filepath.Base(entry.URL))
	// Clone repo
	log.Printf("Attempting to clone %v into %s\n", entry, tempPath)
	err := worker.Downloader.DownloadEntry(entry, tempPath)
	if err != nil {
		return err
	}
	// Parse
	configBytes, err := fs.OpenAndReadFile(filepath.Join(tempPath, entry.GetConfigFile()))
	if err != nil {
		return err
	}
	schema, err := config.Parse(worker.Parser, configBytes)
	if err != nil {
		return err
	}
	return schema.GetResources(destination, worker)
}
