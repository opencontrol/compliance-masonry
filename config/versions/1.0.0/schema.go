package schema

import (
	"errors"
	"github.com/opencontrol/compliance-masonry/config/common"
	"github.com/opencontrol/compliance-masonry/config/common/resources"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"gopkg.in/yaml.v2"
	"log"
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
	resourceGetter resources.ResourceGetter
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
	s.resourceGetter = resources.VCSAndLocalFSGetter{}

	return nil
}

// GetResources will download all the resources that are specified by the v1.0.0 of the schema first by copying the
// local resources then downloading the remote ones and letting their respective schema version handle
// how to get their resources.
func (s *Schema) GetResources(source string, destination string, worker *common.ConfigWorker) error {
	// Local
	// Get Certifications
	log.Println("Retrieving certifications")
	err := s.resourceGetter.GetLocalResources(source, s.Certifications, destination, constants.DefaultCertificationsFolder, false, worker, constants.Certifications)
	if err != nil {
		return err
	}
	// Get Standards
	log.Println("Retrieving standards")
	err = s.resourceGetter.GetLocalResources(source, s.Standards, destination, constants.DefaultStandardsFolder, false, worker, constants.Standards)
	if err != nil {
		return err
	}
	// Get Components
	log.Println("Retrieving components")
	err = s.resourceGetter.GetLocalResources(source, s.Components, destination, constants.DefaultComponentsFolder, true, worker, constants.Components)
	if err != nil {
		return err
	}

	// Remote
	// Get Certifications
	log.Println("Retrieving dependent certifications")
	err = s.resourceGetter.GetRemoteResources(destination, constants.DefaultCertificationsFolder, worker, s.Dependencies.Certifications)
	if err != nil {
		return err
	}
	// Get Standards
	log.Println("Retrieving dependent standards")
	err = s.resourceGetter.GetRemoteResources(destination, constants.DefaultStandardsFolder, worker, s.Dependencies.Standards)
	if err != nil {
		return err
	}
	// Get Components
	log.Println("Retrieving dependent components")
	err = s.resourceGetter.GetRemoteResources(destination, constants.DefaultComponentsFolder, worker, s.Dependencies.Systems)
	if err != nil {
		return err
	}

	return nil
}
