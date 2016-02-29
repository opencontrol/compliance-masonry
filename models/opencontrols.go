package models

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v2"
)

type OpenControl struct {
	Systems       map[string]*System
	Standards     map[string]*Standard
	Certification *Certification
}

func getKey(filePath string) string {
	_, key := filepath.Split(filePath)
	return key
}

func NewOpenControl() *OpenControl {
	return &OpenControl{
		Systems:   make(map[string]*System),
		Standards: make(map[string]*Standard),
	}
}

func LoadData(opencontrolDir string, certificationPath string) *OpenControl {
	var wg sync.WaitGroup
	openControl := NewOpenControl()
	wg.Add(3)
	go func() {
		defer wg.Done()
		openControl.LoadSystems(filepath.Join(opencontrolDir, "components"))
	}()
	go func() {
		defer wg.Done()
		openControl.LoadStandards(filepath.Join(opencontrolDir, "standards"))

	}()
	go func() {
		defer wg.Done()
		openControl.LoadCertification(certificationPath)
	}()
	wg.Wait()
	return openControl
}

func (openControl *OpenControl) LoadSystem(systemDir string) {
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

func (openControl *OpenControl) LoadSystems(opencontrolsDir string) {
	systemsDirs, err := ioutil.ReadDir(opencontrolsDir)
	if err != nil {
		log.Println(err.Error())
	}
	for _, systemDir := range systemsDirs {
		if systemDir.IsDir() {
			openControl.LoadSystem(filepath.Join(opencontrolsDir, systemDir.Name()))
		}
	}
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

func (openControl *OpenControl) LoadStandards(standardsDir string) {
	standardsFiles, err := ioutil.ReadDir(standardsDir)
	if err != nil {
		log.Println(err.Error())
	}
	for _, standardFile := range standardsFiles {
		openControl.LoadStandard(filepath.Join(standardsDir, standardFile.Name()))
	}
}

func (openControl *OpenControl) LoadCertification(certificationFile string) {
	var certification Certification
	certificationData, err := ioutil.ReadFile(certificationFile)
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(certificationData, &certification)
	if err != nil {
		log.Println(err.Error())
	}
	openControl.Certification = &certification
}
