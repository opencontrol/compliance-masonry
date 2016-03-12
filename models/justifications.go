package models

import "sync"

// Justifications struct contains the mapping that links controls to specific components
type Justifications struct {
	mapping map[string]map[string][]string
	sync.RWMutex
}

// NewJustifications creates a new justification
func NewJustifications() *Justifications {
	return &Justifications{mapping: make(map[string]map[string][]string)}
}

// Add methods adds a new mapping to the justification while locking
func (justifications *Justifications) Add(standardKey string, controlKey string, componentKey string) {
	justifications.Lock()
	_, standardKeyExists := justifications.mapping[standardKey]
	if !standardKeyExists {
		justifications.mapping[standardKey] = make(map[string][]string)
	}

	_, controlKeyExists := justifications.mapping[standardKey][controlKey]
	if !controlKeyExists {
		justifications.mapping[standardKey][controlKey] = []string{}
	}
	justifications.mapping[standardKey][controlKey] = append(
		justifications.mapping[standardKey][controlKey], componentKey,
	)
	justifications.Unlock()
}

// LoadMappings loads a set of mappings from a component
func (justifications *Justifications) LoadMappings(component *Component) {
	for _, satsifies := range component.Satisfies {
		justifications.Add(satsifies.StandardKey, satsifies.ControlKey, component.Key)
	}
}
