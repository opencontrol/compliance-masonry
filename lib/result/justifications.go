package result

import (
	"github.com/opencontrol/compliance-masonry/lib/common"
	"sync"
)

// Justifications struct contains the mapping that links controls to specific components
type Justifications struct {
	mapping map[string]map[string]common.Verifications
	sync.RWMutex
}

// NewJustifications creates a new justification
func NewJustifications() *Justifications {
	return &Justifications{mapping: make(map[string]map[string]common.Verifications)}
}

// Add methods adds a new mapping to the justification while locking
func (justifications *Justifications) Add(standardKey string, controlKey string, componentKey string, satisfies common.Satisfies) {
	justifications.Lock()
	newVerification := common.Verification{componentKey, satisfies}
	_, standardKeyExists := justifications.mapping[standardKey]
	if !standardKeyExists {
		justifications.mapping[standardKey] = make(map[string]common.Verifications)
	}
	_, controlKeyExists := justifications.mapping[standardKey][controlKey]
	if !controlKeyExists {
		justifications.mapping[standardKey][controlKey] = common.Verifications{}
	}
	justifications.mapping[standardKey][controlKey] = append(
		justifications.mapping[standardKey][controlKey], newVerification,
	)
	justifications.Unlock()
}

// LoadMappings loads a set of mappings from a component
func (justifications *Justifications) LoadMappings(component common.Component) {
	for _, satisfies := range component.GetAllSatisfies() {
		justifications.Add(satisfies.GetStandardKey(), satisfies.GetControlKey(), component.GetKey(), satisfies)
	}
}

// Get retrieves justifications for a specific standard and control
func (justifications *Justifications) Get(standardKey string, controlKey string) common.Verifications {
	_, standardKeyExists := justifications.mapping[standardKey]
	if !standardKeyExists {
		return nil
	}
	controlJustifications, controlKeyExists := justifications.mapping[standardKey][controlKey]
	if !controlKeyExists {
		return nil
	}
	return controlJustifications
}
