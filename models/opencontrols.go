package models

import (
	"io/ioutil"
	"log"
	"os"
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
		openControl.LoadSystem(".")
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

func (openControl *OpenControl) LoadSystems(opencontrolsDir string) {
	systemsDirs, err := ioutil.ReadDir(opencontrolsDir)
	if err != nil {
		log.Println(err.Error())
	}
	for _, systemDir := range systemsDirs {
		if systemDir.IsDir() {
			if _, err := os.Stat(filepath.Join(opencontrolsDir, "system.yaml")); err == nil {
				openControl.LoadSystem(filepath.Join(opencontrolsDir, systemDir.Name()))
			}
		}
	}
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
