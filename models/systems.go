package models

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type System struct {
	Name       string `yaml:"name" json:"name"`
	Key        string `yaml:"key" json:"key"`
	Components map[string]*Component
}

func NewSystem() *System {
	return &System{Components: make(map[string]*Component)}
}

func (system *System) LoadComponent(component_dir string) {
	var component *Component
	component_data, err := ioutil.ReadFile(filepath.Join(component_dir, "component.yaml"))
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(component_data, &component)
	if err != nil {
		log.Println(err.Error())
	}
	if system.Key == "" {
		component.Key = getKey(component_dir)
	}
	system.LoadComponents(component_dir)
	system.Components[component.Key] = component
}

func (system *System) LoadComponents(system_dir string) {
	components_dir, err := ioutil.ReadDir(system_dir)
	if err != nil {
		log.Println(err.Error())
	}
	for _, component_dir := range components_dir {
		if component_dir.IsDir() {
			system.LoadComponent(filepath.Join(system_dir, component_dir.Name()))
		}
	}
}
