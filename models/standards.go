package models

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

// Control struct stores data on a specific security requirement
type Control struct {
	Family string `yaml:"family" json:"family"`
	Name   string `yaml:"name" json:"name"`
}

// Standard struct is a collection of security requirements
type Standard struct {
	Key      string             `yaml:"name" json:"name"`
	Controls map[string]Control `yaml:",inline"`
}

// Standards struct is a thread save mapping of Standards
type Standards struct {
	mapping map[string]*Standard
	sync.RWMutex
}

// NewStandards creates an instance of Components struct
func NewStandards() *Standards {
	return &Standards{mapping: make(map[string]*Standard)}
}

// Add adds a standard to the standards mapping
func (standards *Standards) Add(standard *Standard) {
	standards.Lock()
	standards.mapping[standard.Key] = standard
	standards.Unlock()
}

// Get retrieves a standard
func (standards *Standards) Get(key string) *Standard {
	standards.Lock()
	defer standards.Unlock()
	return standards.mapping[key]
}

// GetAll retrieves all the standards
func (standards *Standards) GetAll() map[string]*Standard {
	return standards.mapping
}

// LoadStandard imports a standard into the Standard struct and adds it to the
// main object.
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
	openControl.Standards.Add(&standard)
}
