package get

import (
	"github.com/opencontrol/compliance-masonry/commands/get/resources"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol"
)

// Get will retrieve all of the resources for the schemas and the resources for all the dependent schemas.
func Get(destination string, configData []byte) error {
	// Check the data.
	if configData == nil || len(configData) == 0 {
		return common.ErrNoDataToParse
	}
	// Parse it.
	parser := opencontrol.YAMLParser{}
	configSchema, err := parser.Parse(configData)
	if err != nil {
		return err
	}
	// Get Resources
	getter := resources.NewVCSAndLocalGetter(parser)
	err = resources.GetResources("", destination, configSchema, getter)
	if err != nil {
		return err
	}
	return nil
}
