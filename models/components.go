package models

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v2"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
)

// Components struct is a thread-safe structure mapping for components
type Components struct {
	mapping map[string]base.Component
	sync.RWMutex
}

// Component struct is an individual component requiring documentation
// Schema info: https://github.com/opencontrol/schemas#component-yaml
type Component struct {
	Name          string                  `yaml:"name" json:"name"`
	Key           string                  `yaml:"key" json:"key"`
	References    common.GeneralReferences      `yaml:"references" json:"references"`
	Verifications common.VerificationReferences `yaml:"verifications" json:"verifications"`
	Satisfies     []Satisfies          `yaml:"satisfies" json:"satisfies"`
	SchemaVersion float32                 `yaml:"schema_version" json:"schema_version"`
}

func (c Component) GetName() string {
	return c.Name
}

func (c Component) GetKey() string {
	return c.Key
}

func (c *Component) SetKey(key string) {
	c.Key = key
}

func (c Component) GetVerifications() *common.VerificationReferences {
	return &c.Verifications
}

func (c Component) GetReferences() *common.GeneralReferences {
	return &c.References
}

func (c Component) GetAllSatisfies() []base.Satisfies {
	// Have to do manual conversion from this Component's Satisfies to the interface base.Satisfies.
	baseSatisfies := make([]base.Satisfies, len(c.Satisfies))
	for idx, value := range c.Satisfies {
		baseSatisfies[idx] = value
	}
	return baseSatisfies
}

func (c Component) GetVersion() float32 {
	return c.SchemaVersion
}

// Satisfies struct contains data demonstrating why a specific component meets
// a control
// This struct is a one-to-one mapping of a `satisfies` item in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type Satisfies struct {
	ControlKey  string        `yaml:"control_key" json:"control_key"`
	StandardKey string        `yaml:"standard_key" json:"standard_key"`
	Narrative   string        `yaml:"narrative" json:"narrative"`
	CoveredBy   common.CoveredByList `yaml:"covered_by" json:"covered_by"`
}

func (s Satisfies) GetControlKey() string {
	return s.ControlKey
}

func (s Satisfies) GetStandardKey() string {
	return s.StandardKey
}

func (s Satisfies) GetNarrative() string {
	return s.Narrative
}

func (s Satisfies) GetCoveredBy() common.CoveredByList {
	return s.CoveredBy
}

// NewComponents creates an instance of Components struct
func NewComponents() *Components {
	return &Components{mapping: make(map[string]base.Component)}
}

// Add adds a new component to the component map
func (components *Components) Add(component base.Component) {
	components.Lock()
	components.mapping[component.GetKey()] = component
	components.Unlock()
}

// Get retrieves a new component from the component map
func (components *Components) Get(key string) base.Component {
	components.Lock()
	defer components.Unlock()
	return components.mapping[key]
}

// GetAndApply get a component and apply the callback function inside while locking
// components
func (components *Components) GetAndApply(key string, callback func(component base.Component)) {
	components.Lock()
	callback(components.mapping[key])
	components.Unlock()
}

// CompareAndAdd compares to see if the component exists in the map. If not, it adds the component.
// This function is thread-safe.
func (components *Components) CompareAndAdd(component base.Component) bool {
	components.Lock()
	defer components.Unlock()
	added := false
	if _, exists := components.mapping[component.GetKey()]; !exists {
		components.mapping[component.GetKey()] = component
		added = true
	} else {
		log.Fatalf("Component: %s exisits!\n", component.GetKey())
	}
	return added
}

// GetAll retrieves all the components
func (components *Components) GetAll() map[string]base.Component {
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