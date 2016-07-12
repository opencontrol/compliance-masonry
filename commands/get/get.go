package get

import (
	"github.com/opencontrol/compliance-masonry/config"
	"github.com/opencontrol/compliance-masonry/config/common"
	"github.com/opencontrol/compliance-masonry/config/parser"
)

// Get will retrieve all of the resources for the schemas and the resources for all the dependent schemas.
func Get(destination string, configData []byte, worker *common.ConfigWorker) error {
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
	err = configSchema.GetResources("", destination, worker)
	if err != nil {
		return err
	}
	return nil
}
