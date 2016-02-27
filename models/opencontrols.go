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

func getKey(file_path string) string {
	_, key := filepath.Split(file_path)
	return key
}

func NewOpenControl() *OpenControl {
	return &OpenControl{
		Systems:   make(map[string]*System),
		Standards: make(map[string]*Standard),
	}
}

func LoadData(opencontrol_dir string, certification_path string) *OpenControl {
	var wg sync.WaitGroup
	openControl := NewOpenControl()
	wg.Add(3)
	go func() {
		defer wg.Done()
		openControl.LoadSystems(filepath.Join(opencontrol_dir, "components"))
	}()
	go func() {
		defer wg.Done()
		openControl.LoadStandards(filepath.Join(opencontrol_dir, "standards"))

	}()
	go func() {
		defer wg.Done()
		openControl.LoadCertification(certification_path)
	}()
	wg.Wait()
	return openControl
}

func (openControl *OpenControl) LoadSystem(system_dir string) {
	system := NewSystem()
	system_data, err := ioutil.ReadFile(filepath.Join(system_dir, "system.yaml"))
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(system_data, &system)
	if err != nil {
		log.Println(err.Error())
	}
	if system.Key == "" {
		system.Key = getKey(system_dir)
	}
	system.LoadComponents(system_dir)
	openControl.Systems[system.Key] = system
}

func (openControl *OpenControl) LoadSystems(opencontrols_dir string) {
	systems_dirs, err := ioutil.ReadDir(opencontrols_dir)
	if err != nil {
		log.Println(err.Error())
	}
	for _, system_dir := range systems_dirs {
		if system_dir.IsDir() {
			openControl.LoadSystem(filepath.Join(opencontrols_dir, system_dir.Name()))
		}
	}
}

func (openControl *OpenControl) LoadStandard(standard_file string) {
	var standard Standard
	standard_data, err := ioutil.ReadFile(standard_file)
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(standard_data, &standard)
	if err != nil {
		log.Println(err.Error())
	}
	openControl.Standards[standard.Key] = &standard
}

func (openControl *OpenControl) LoadStandards(standards_dir string) {
	standards_files, err := ioutil.ReadDir(standards_dir)
	if err != nil {
		log.Println(err.Error())
	}
	for _, standard_file := range standards_files {
		openControl.LoadStandard(filepath.Join(standards_dir, standard_file.Name()))
	}
}

func (openControl *OpenControl) LoadCertification(certification_file string) {
	var certification Certification
	certification_data, err := ioutil.ReadFile(certification_file)
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(certification_data, &certification)
	if err != nil {
		log.Println(err.Error())
	}
	openControl.Certification = &certification
}
