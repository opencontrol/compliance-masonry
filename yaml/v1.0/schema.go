package v1_0

import (
	"gopkg.in/yaml.v2"
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
)

type Schema struct {
	common.Base `yaml:",inline"`
	Meta Metadata `yaml:"metadata"`
	SystemName string `yaml:"system_name"`
	Components []string `yaml:",flow"`
	Dependencies Dependencies `yaml:"dependencies"`
}

type Dependencies struct {
	Certification Entry `yaml:"certification"`
	Systems []Entry `yaml:",flow"`
	Standards []Entry `yaml:",flow"`
}

type Metadata struct {
	Description string `yaml:"description"`
	Maintainers []string `yaml:",flow"`
}

type Entry struct {
	Protocol string `yaml:"protocol"`
	URL string `yaml:"url"`
	Revision string `yaml:"revision"`
}

func (s *Schema) Parse(data []byte) error {
	err := yaml.Unmarshal(data, s)
	if err != nil {
		return err
	}

	return nil
}
