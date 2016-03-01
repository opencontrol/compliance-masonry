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

func (openControl *OpenControl) LoadSystem(systemDir string) {
	if _, err := os.Stat(filepath.Join(systemDir, "system.yaml")); err == nil {
		system := NewSystem()
		systemData, err := ioutil.ReadFile(filepath.Join(systemDir, "system.yaml"))
		if err != nil {
			log.Println(err.Error())
		}
		err = yaml.Unmarshal(systemData, &system)
		if err != nil {
			log.Println(err.Error())
		}
		if system.Key == "" {
			system.Key = getKey(systemDir)
		}
		system.LoadComponents(systemDir)
		openControl.Systems[system.Key] = system
	}
}

func (system *System) LoadComponents(systemDir string) {
	componentsDir, err := ioutil.ReadDir(systemDir)
	if err != nil {
		log.Println(err.Error())
	}
	for _, componentDir := range componentsDir {
		if componentDir.IsDir() {
			componentDir := filepath.Join(systemDir, componentDir.Name())
			system.LoadComponent(componentDir)
		}
	}
}
