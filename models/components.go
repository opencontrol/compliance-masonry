package models

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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

// CoveredBy struct is the pointing mechanism for for refering to
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
		if openControl.Components[component.Key] != nil {
			log.Fatalln("Component: %s exisits!", component.Key)
		}
		openControl.Justifications.LoadMappings(component)
		openControl.Components[component.Key] = component
	}
}
