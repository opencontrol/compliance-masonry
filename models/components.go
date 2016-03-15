package models

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v2"
)

// GeneralReference struct contains data for the name and path of a
// compliance reference.
type GeneralReference struct {
	Name string `yaml:"name" json:"name"`
	Path string `yaml:"path" json:"path"`
	Type string `yaml:"type" json:"type"`
}

// VerificationReference struct is a general reference that verifies a specific
// control, it can be pointed to in the control documentation.
type VerificationReference struct {
	Key              string `yaml:"key" json:"key"`
	GeneralReference `yaml:",inline"`
}

// CoveredBy struct is the pointing mechanism for for referring to
// VerificationReferences in the documentation.
type CoveredBy struct {
	ComponentKey    string `yaml:"component_key" json:"component_key"`
	VerificationKey string `yaml:"verification_key" json:"verification_key"`
}

// Satisfies struct contains data demonstrating why a specific component meets
// a control
type Satisfies struct {
	ControlKey  string      `yaml:"control_key" json:"control_key"`
	StandardKey string      `yaml:"standard_key" json:"standard_key"`
	Narrative   string      `yaml:"narrative" json:"narrative"`
	CoveredBy   []CoveredBy `yaml:"covered_by" json:"covered_by"`
}

// Component struct is an individual component requiring documentation
type Component struct {
	Name          string                  `yaml:"name" json:"name"`
	Key           string                  `yaml:"key" json:"key"`
	References    []GeneralReference      `yaml:"references" json:"references"`
	Verifications []VerificationReference `yaml:"verifications" json:"verifications"`
	Satisfies     []Satisfies             `yaml:"satisfies" json:"satisfies"`
	SchemaVersion float32                 `yaml:"schema_version" json:"schema_version"`
}

// Components struct is a thread-safe structure mapping for components
type Components struct {
	mapping map[string]*Component
	sync.RWMutex
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
		log.Fatalln("Component: %s exisits!", component.Key)
	}
	return added
}

// GetAll retrieves all the components
func (components *Components) GetAll() map[string]*Component {
	return components.mapping
}

// LoadComponent imports components into a Component struct and adds it to the
// Components map.
func (openControl *OpenControl) LoadComponent(componentDir string) {
	if _, err := os.Stat(filepath.Join(componentDir, "component.yaml")); err == nil {
		var component *Component
		componentData, err := ioutil.ReadFile(filepath.Join(componentDir, "component.yaml"))
		if err != nil {
			log.Println("here", err.Error())
		}
		err = yaml.Unmarshal(componentData, &component)
		if err != nil {
			log.Println(err.Error())
		}
		if component.Key == "" {
			component.Key = getKey(componentDir)
		}
		if openControl.Components.CompareAndAdd(component) {
			openControl.Justifications.LoadMappings(component)
		}
	}
}
