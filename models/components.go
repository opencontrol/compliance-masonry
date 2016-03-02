package models

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type GeneralReference struct {
	Name string `yaml:"name" json:"name"`
	Path string `yaml:"path" json:"path"`
	Type string `yaml:"type" json:"type"`
}

type VerificationReference struct {
	Key              string `yaml:"key" json:"key"`
	GeneralReference `yaml:",inline"`
}

type CoveredBy struct {
	ComponentKey    string `yaml:"component_key" json:"component_key"`
	SystemKey       string `yaml:"system_key" json:"system_key"`
	VerificationKey string `yaml:"verification_key" json:"verification_key"`
}

type Satisfies struct {
	ControlKey  string      `yaml:"control_key" json:"control_key"`
	StandardKey string      `yaml:"standard_key" json:"standard_key"`
	Narrative   string      `yaml:"narrative" json:"narrative"`
	CoveredBy   []CoveredBy `yaml:"covered_by" json:"covered_by"`
}

type Component struct {
	Name          string                  `yaml:"name" json:"name"`
	Key           string                  `yaml:"key" json:"key"`
	References    []GeneralReference      `yaml:"references" json:"references"`
	Verifications []VerificationReference `yaml:"verifications" json:"verifications"`
	Satisfies     Satisfies               `yaml:"satsifies" json:"satsifies"`
	SchemaVersion float32                 `yaml:"schema_version" json:"schema_version"`
}

func (system *System) LoadComponent(componentDir string) {
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
		if system.Components[system.Key] != nil {
			log.Fatalln("Component: %s exisits!", component.Key)
		}
		system.Components[component.Key] = component
	}
}
