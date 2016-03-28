package inventory

import (
	"sync"

	"github.com/opencontrol/compliance-masonry-go/models"
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

// InitInventory initialize an Inventory struct
func InitInventory() (*Inventory, error) {
	return &Inventory{models.NewOpenControl()}, nil
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
