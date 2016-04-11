package models

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v2"
)

// Components struct is a thread-safe structure mapping for components
type Components struct {
	mapping map[string]*Component
	sync.RWMutex
}

// Component struct is an individual component requiring documentation
// Schema info: https://github.com/opencontrol/schemas#component-yaml
type Component struct {
	Name          string                  `yaml:"name" json:"name"`
	Key           string                  `yaml:"key" json:"key"`
	References    *GeneralReferences      `yaml:"references" json:"references"`
	Verifications *VerificationReferences `yaml:"verifications" json:"verifications"`
	Satisfies     *SatisfiesList          `yaml:"satisfies" json:"satisfies"`
	SchemaVersion float32                 `yaml:"schema_version" json:"schema_version"`
}

// SatisfiesList is a list of Satisfies
type SatisfiesList []Satisfies

// Satisfies struct contains data demonstrating why a specific component meets
// a control
// This struct is a one-to-one mapping of a `satisfies` item in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type Satisfies struct {
	ControlKey  string        `yaml:"control_key" json:"control_key"`
	StandardKey string        `yaml:"standard_key" json:"standard_key"`
	Narrative   string        `yaml:"narrative" json:"narrative"`
	CoveredBy   CoveredByList `yaml:"covered_by" json:"covered_by"`
}

// NewComponents creates an instance of Components struct
func NewComponents() *Components {
	return &Components{mapping: make(map[string]*Component)}
}

// Add adds a new component to the component map
func (components *Components) Add(component *Component) {
	components.Lock()
	components.mapping[component.Key] = component
	components.Unlock()
}

// Get retrieves a new component from the component map
func (components *Components) Get(key string) *Component {
	components.Lock()
	defer components.Unlock()
	return components.mapping[key]
}

// GetAndApply get a component and apply the callback function inside while locking
// components
func (components *Components) GetAndApply(key string, callback func(component *Component)) {
	components.Lock()
	callback(components.mapping[key])
	components.Unlock()
}

// CompareAndAdd compares to see if the component exists in the map. If not, it adds the component.
// This function is thread-safe.
func (components *Components) CompareAndAdd(component *Component) bool {
	components.Lock()
	defer components.Unlock()
	added := false
	if _, exists := components.mapping[component.Key]; !exists {
		components.mapping[component.Key] = component
		added = true
	} else {
		log.Fatalf("Component: %s exisits!\n", component.Key)
	}
	return added
}

// GetAll retrieves all the components
func (components *Components) GetAll() map[string]*Component {
	return components.mapping
}

// LoadComponent imports components into a Component struct and adds it to the
// Components map.
func (openControl *OpenControl) LoadComponent(componentDir string) error {
	_, err := os.Stat(filepath.Join(componentDir, "component.yaml"))
	if err != nil {
		return ErrComponentFileDNE
	}
	var component *Component
	componentData, err := ioutil.ReadFile(filepath.Join(componentDir, "component.yaml"))
	if err != nil {
		return ErrReadFile
	}
	err = yaml.Unmarshal(componentData, &component)
	if err != nil {
		return ErrControlSchema
	}
	if component.Key == "" {
		component.Key = getKey(componentDir)
	}
	if openControl.Components.CompareAndAdd(component) {
		openControl.Justifications.LoadMappings(component)
	}
	return nil
}

// Len retruns the length of a SatisfiesList struct
func (slice SatisfiesList) Len() int {
	return len(slice)
}
