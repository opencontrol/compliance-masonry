package models

import (
	"io/ioutil"
	"log"
	"os"
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
		log.Println("here",err.Error())
	}
	err = yaml.Unmarshal(componentData, &component)
	if err != nil {
		log.Println(err.Error())
	}
	if component.Key == "" {
		component.Key = getKey(componentDir)
	}
	system.Components[component.Key] = component
}

func (system *System) LoadComponents(systemDir string) {
	componentsDir, err := ioutil.ReadDir(systemDir)
	if err != nil {
		log.Println(err.Error())
	}
	for _, componentDir := range componentsDir {
		if componentDir.IsDir() {
			componentDir := filepath.Join(systemDir, componentDir.Name())
			if _, err := os.Stat(filepath.Join(componentDir, "component.yaml")); err == nil {
					system.LoadComponent(componentDir)
			}
		}
	}
}
