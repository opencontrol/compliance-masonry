package main

import (
	"github.com/opencontrol/compliance-masonry-go/yaml"
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
	"github.com/opencontrol/compliance-masonry-go/yaml/parser"
)

const (
	DefaultDestination = "opencontrols"
	DefaultConfigYaml  = "opencontrol.yaml"
)

func Get(destination string, configData []byte) error {
	// Check the data.
	if configData == nil || len(configData) == 0 {
		return yaml.ErrNoDataToParse
	}
	// Parse it.
	configSchema, err := yaml.Parse(parser.Parser{}, configData)
	if err != nil {
		return err
	}
	// Get Resources
	err = configSchema.GetResources(destination, common.VCSEntryDownloader{})
	if err != nil {
		return err
	}
	return nil
}
