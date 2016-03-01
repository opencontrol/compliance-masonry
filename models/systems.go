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

func (system *System) LoadComponent(componentDir string) {
	var component *Component
	componentData, err := ioutil.ReadFile(filepath.Join(componentDir, "component.yaml"))
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(componentData, &component)
	if err != nil {
		log.Println(err.Error())
	}
	if component.Key == "" {
		component.Key = getKey(componentDir)
	}
	system.LoadComponents(componentDir)
	system.Components[component.Key] = component
}

func (system *System) LoadComponents(systemDir string) {
	componentsDir, err := ioutil.ReadDir(systemDir)
	if err != nil {
		log.Println(err.Error())
	}
	for _, componentDir := range componentsDir {
		if componentDir.IsDir() {
			system.LoadComponent(filepath.Join(systemDir, componentDir.Name()))
		}
	}
}
