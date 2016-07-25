package common

import "github.com/opencontrol/compliance-masonry/tools/constants"

// Entry is a generic holder for handling the specific location and revision of a resource.
type Entry struct {
	URL      string `yaml:"url"`
	Revision string `yaml:"revision"`
	Path     string `yaml:"path"`
}

// GetConfigFile is a getter for the config file name. Will return DefaultConfigYaml value if none has been set.
func (e Entry) GetConfigFile() string {
	if e.Path == "" {
		return constants.DefaultConfigYaml
	}
	return e.Path
}
