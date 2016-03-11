package models

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

// OpenControl struct combines data on of components, standards, and a certification
// for creating and verifying component documentation.
type OpenControl struct {
	Components    map[string]*Component
	Standards     map[string]*Standard
	Certification *Certification
}

// getKey extracts a component key from the filepath
func getKey(filePath string) string {
	_, key := filepath.Split(filePath)
	return key
}

// NewOpenControl Initalizes an empty OpenControl struct
func NewOpenControl() *OpenControl {
	return &OpenControl{
		Components: make(map[string]*Component),
		Standards:  make(map[string]*Standard),
	}
}

// LoadData creates a new instance of OpenControl struct and loads
// the components, standards, and certification data.
func LoadData(openControlDir string, certificationPath string) *OpenControl {
	var wg sync.WaitGroup
	openControl := NewOpenControl()
	wg.Add(3)
	go func() {
		defer wg.Done()
		openControl.LoadComponents(filepath.Join(openControlDir, "components"))
		openControl.LoadComponents(".")
	}()
	go func() {
		defer wg.Done()
		openControl.LoadStandards(filepath.Join(openControlDir, "standards"))

	}()
	go func() {
		defer wg.Done()
		openControl.LoadCertification(certificationPath)
	}()
	wg.Wait()
	return openControl
}

// LoadComponents loads multiple components by searching for components in a
// given directory
func (openControl *OpenControl) LoadComponents(directory string) {
	componentsDir, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println(err.Error())
	}
	for _, componentDir := range componentsDir {
		if componentDir.IsDir() {
			componentDir := filepath.Join(directory, componentDir.Name())
			openControl.LoadComponent(componentDir)
		}
	}
}

// LoadStandards loads multiple standards by searching for components in a
// given directory
func (openControl *OpenControl) LoadStandards(standardsDir string) {
	standardsFiles, err := ioutil.ReadDir(standardsDir)
	if err != nil {
		log.Println(err.Error())
	}
	for _, standardFile := range standardsFiles {
		openControl.LoadStandard(filepath.Join(standardsDir, standardFile.Name()))
	}
}
