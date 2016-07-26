package get

import (
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol/versions/base"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol/parser"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol"
)

// Get will retrieve all of the resources for the schemas and the resources for all the dependent schemas.
func Get(destination string, configData []byte, worker *base.Worker) error {
	// Check the data.
	if configData == nil || len(configData) == 0 {
		return common.ErrNoDataToParse
	}
	// Parse it.
	configSchema, err := opencontrol.Parse(parser.Parser{}, configData)
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
