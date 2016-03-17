package schema

import (
	"errors"
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
	"gopkg.in/yaml.v2"
)

const (
	ErrMalformedV1_0_0YamlPrefix = "Unable to parse yaml data"
)

// Schema contains the structs for the v1.0.0 schema
type Schema struct {
	common.Base  `yaml:",inline"`
	Meta         Metadata     `yaml:"metadata"`
	SystemName   string       `yaml:"system_name"`
	Components   []string     `yaml:",flow"`
	Dependencies Dependencies `yaml:"dependencies"`
}

// Dependencies contains all the dependencies for the system
type Dependencies struct {
	Certification Entry   `yaml:"certification"`
	Systems       []Entry `yaml:",flow"`
	Standards     []Entry `yaml:",flow"`
}

// Metadata contains metadata about the system.
type Metadata struct {
	Description string   `yaml:"description"`
	Maintainers []string `yaml:",flow"`
}

// Entry is a generic holder for handling the specific location and revision of a resource.
type Entry struct {
	Protocol string `yaml:"protocol"`
	URL      string `yaml:"url"`
	Revision string `yaml:"revision"`
}

// Parse will parse using it's own schema. In this case the v1.0.0 schema.
func (s *Schema) Parse(data []byte) error {
	err := yaml.Unmarshal(data, s)
	if err != nil {
		return errors.New(ErrMalformedV1_0_0YamlPrefix + " - " + err.Error())
	}

	return nil
}
