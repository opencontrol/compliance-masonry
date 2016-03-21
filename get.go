package main

import (
	"github.com/opencontrol/compliance-masonry-go/config"
	"github.com/opencontrol/compliance-masonry-go/config/common"
	"github.com/opencontrol/compliance-masonry-go/config/parser"
)

func Get(destination string, configData []byte) error {
	// Check the data.
	if configData == nil || len(configData) == 0 {
		return config.ErrNoDataToParse
	}
	// Parse it.
	configSchema, err := config.Parse(parser.Parser{}, configData)
	if err != nil {
		return err
	}
	// Get Resources
	worker := common.ConfigWorker{Parser: parser.Parser{}, Downloader: common.VCSEntryDownloader{}}
	err = configSchema.GetResources(destination, worker)
	if err != nil {
		return err
	}
	return nil
}
