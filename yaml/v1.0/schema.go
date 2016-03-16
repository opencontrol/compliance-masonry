package v1_0

import (
	"gopkg.in/yaml.v2"
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
)

type Schema struct {
	common.Base `yaml:",inline"`
	Metadata `yaml:",inline"`
	Entries []Entry `yaml:",flow"`
}

type Metadata struct {
	Name string `yaml:"name"`
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
