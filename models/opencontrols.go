package models

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v2"
)

type OpenControl struct {
	Components    map[string]*Component
	Standards     map[string]*Standard
	Certification *Certification
}

func getKey(filePath string) string {
	_, key := filepath.Split(filePath)
	return key
}

func NewOpenControl() *OpenControl {
	return &OpenControl{
		Components: make(map[string]*Component),
		Standards:  make(map[string]*Standard),
	}
}

func LoadData(opencontrolDir string, certificationPath string) *OpenControl {
	var wg sync.WaitGroup
	openControl := NewOpenControl()
	wg.Add(3)
	go func() {
		defer wg.Done()
		openControl.LoadComponents(filepath.Join(opencontrolDir, "components"))
		openControl.LoadComponents(".")
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

func (opencontrol *OpenControl) LoadComponents(directory string) {
	componentsDir, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Println(err.Error())
	}
	for _, componentDir := range componentsDir {
		if componentDir.IsDir() {
			componentDir := filepath.Join(directory, componentDir.Name())
			opencontrol.LoadComponent(componentDir)
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
