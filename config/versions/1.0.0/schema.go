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

func (s *Schema) GetResources(destination string, worker *common.ConfigWorker) error {
	// Local
	// Get Certifications
	log.Println("Retrieving certifications")
	err := s.getLocalResources(s.Certifications, destination, constants.DefaultCertificationsFolder, false)
	if err != nil {
		return err
	}
	// Get Standards
	log.Println("Retrieving standards")
	err = s.getLocalResources(s.Standards, destination, constants.DefaultStandardsFolder, false)
	if err != nil {
		return err
	}
	// Get Components
	log.Println("Retrieving components")
	err = s.getLocalResources(s.Components, destination, constants.DefaultComponentsFolder, true)
	if err != nil {
		return err
	}

	// Remote
	// Get Certifications
	log.Println("Retrieving dependent certifications")
	err = s.getRemoteResources(destination, constants.DefaultCertificationsFolder, worker, s.Dependencies.Certifications)
	if err != nil {
		return err
	}
	// Get Standards
	log.Println("Retrieving dependent standards")
	err = s.getRemoteResources(destination, constants.DefaultStandardsFolder, worker, s.Dependencies.Standards)
	if err != nil {
		return err
	}
	// Get Components
	log.Println("Retrieving dependent components")
	err = s.getRemoteResources(destination, constants.DefaultComponentsFolder, worker, s.Dependencies.Systems)
	if err != nil {
		return err
	}

	return nil
}

func (s *Schema) getLocalResources(resources []string, destination string, subfolder string, recursively bool) error {
	for _, resource := range resources {
		var err error
		if recursively {
			err = ufs.CopyAll(resource,
				filepath.Join(destination, subfolder, filepath.Base(resource)), nil)
		} else {
			err = ufs.CopyFile(resource,
				filepath.Join(destination, subfolder, filepath.Base(resource)))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Schema) getRemoteResources(destination string, subfolder string, worker *common.ConfigWorker, entries []common.Entry) error {
	tempResourcesDir, err := ioutil.TempDir("", "opencontrol-resources")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempResourcesDir)
	for _, entry := range entries{
		tempPath := filepath.Join(tempResourcesDir, subfolder, filepath.Base(entry.URL))
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
		err = schema.GetResources(destination, worker)
		if err != nil {
			return err
		}
	}
	return nil
}
