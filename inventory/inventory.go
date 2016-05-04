package inventory

import (
	"sync"

	"github.com/opencontrol/compliance-masonry/config"
	"github.com/opencontrol/compliance-masonry/config/parser"
	"github.com/opencontrol/compliance-masonry/models"
)

// Inventory struct is an extension of models.OpenControl that adds
// an exportPath
type Inventory struct {
	*models.OpenControl
}

// ImplementationMapping struct is a map of component and the implementation of a given control
type ImplementationMapping struct {
	sync.RWMutex
	mapping map[string]string
}

// ControlInfo struct stores information about the control documentation,
// Exists flags if some documentation exists
// Implementations is a map of component and the implementation of the given control
type ControlInfo struct {
	Exists          bool
	Implementations ImplementationMapping
}

// GetLocalComponents uses the opencontrol.yaml file to get a list of local components
func GetLocalComponents(configBytes []byte) ([]string, error) {
	configSchema, err := config.Parse(parser.Parser{}, configBytes)
	if err != nil {
		return nil, err
	}
	return configSchema.GetLocalComponents(), nil
}

// GetRequiredComponents uses the opencontrol.yaml file to get a list of required controls
func GetRequiredComponents(configBytes []byte) ([]string, error) {
	configSchema, err := config.Parse(parser.Parser{}, configBytes)
	if err != nil {
		return nil, err
	}
	return configSchema.GetRequiredComponents(), nil
}

// LoadLocalComponents loads a set of components given a path
func (inventory *Inventory) LoadLocalComponents(componentPaths []string) error {
	for _, componentPath := range componentPaths {
		err := inventory.LoadComponent(componentPath)
		return err
	}
	return nil
}

// InitInventory initialize an Inventory struct
func InitInventory(configBytes []byte) (*Inventory, error) {
	inventory := &Inventory{models.NewOpenControl()}
	components, err := GetLocalComponents(configBytes)
	if err != nil {
		return nil, err
	}
	err = inventory.LoadLocalComponents(components)
	if err != nil {
		return nil, err
	}
	return inventory, nil
}

//InitControlInfo initialize a ControlInfo struct
func InitControlInfo() (*ControlInfo, error) {
	return &ControlInfo{false, ImplementationMapping{mapping: make(map[string]string)}}, nil
}

// GetControlInfo retrives the control info for a particular control
func (inventory *Inventory) GetControlInfo(standard string, control string) *ControlInfo {
	controlInfo, _ := InitControlInfo()
	inventory.Justifications.GetAndApply(standard, control, func(verifications models.Verifications) {
		if verifications.Len() > 0 {
			controlInfo.Exists = true
		}
		for _, verification := range verifications {
			controlInfo.Implementations.mapping[verification.ComponentKey] = verification.SatisfiesData.ImplementationStatus
		}
	})
	return controlInfo
}
