package models

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Control struct {
	Family string `yaml:"family" json:"family"`
	Name   string `yaml:"name" json:"name"`
}

type Standard struct {
	Key      string             `yaml:"name" json:"name"`
	Controls map[string]Control `yaml:",inline"`
}

func (openControl *OpenControl) LoadStandard(standardFile string) {
	var standard Standard
	standardData, err := ioutil.ReadFile(standardFile)
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(standardData, &standard)
	if err != nil {
		log.Println(err.Error())
	}
	openControl.Standards[standard.Key] = &standard
}
