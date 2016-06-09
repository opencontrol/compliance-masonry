package models

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"github.com/opencontrol/compliance-masonry/models/components/versions"
	"github.com/opencontrol/compliance-masonry/models/components"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/tools/fs"
)

var (
	// ErrReadDir is raised when a directory can not be read
	ErrReadDir = errors.New("Unable to read the directory")
	// ErrReadFile is raised when a file can not be read
	ErrReadFile = errors.New("Unable to read the file")
	// ErrCertificationSchema is raised a certification cannot be parsed
	ErrCertificationSchema = errors.New("Unable to parse certification")
	// ErrStandardSchema is raised a standard cannot be parsed
	ErrStandardSchema = errors.New("Unable to parse standard")
)

// OpenControl struct combines components, standards, and a certification data
// For more information on the opencontrol schema visit: https://github.com/opencontrol/schemas
type OpenControl struct {
	Components     *components.Components
	Standards      *Standards
	Justifications *Justifications
	Certification  *Certification
}

// getKey extracts a component key from the filepath
func getKey(filePath string) string {
	_, key := filepath.Split(filePath)
	return key
}

// NewOpenControl initializes an empty OpenControl struct
func NewOpenControl() *OpenControl {
	return &OpenControl{
		Justifications: NewJustifications(),
		Components:     components.NewComponents(),
		Standards:      NewStandards(),
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
func (openControl *OpenControl) LoadComponents(directory string) error {
	var wg sync.WaitGroup
	componentsDir, err := ioutil.ReadDir(directory)
	if err != nil {
		return ErrReadDir
	}
	for _, componentDir := range componentsDir {
		wg.Add(1)
		go func(componentDir os.FileInfo) {
			if componentDir.IsDir() {
				componentDir := filepath.Join(directory, componentDir.Name())
				openControl.LoadComponent(componentDir)
			}
			wg.Done()
		}(componentDir)
	}
	wg.Wait()
	return nil
}

// LoadStandards loads multiple standards by searching for components in a
// given directory
func (openControl *OpenControl) LoadStandards(standardsDir string) error {
	var wg sync.WaitGroup

	standardsFiles, err := ioutil.ReadDir(standardsDir)
	if err != nil {
		return ErrReadDir
	}
	for _, standardFile := range standardsFiles {
		wg.Add(1)
		go func(standardFile os.FileInfo) {
			openControl.LoadStandard(filepath.Join(standardsDir, standardFile.Name()))
			wg.Done()
		}(standardFile)
	}
	wg.Wait()
	return nil
}


// LoadComponent imports components into a Component struct and adds it to the
// Components map.
func (openControl *OpenControl) LoadComponent(componentDir string) error {
	// Get file system assistance.
	fs := fs.OSUtil{}
	// Read the component file.
	componentData, err := fs.OpenAndReadFile(filepath.Join(componentDir, "component.yaml"))
	if err != nil {
		return errors.New(constants.ErrComponentFileDNE)
	}
	// Parse the component.
	var component base.Component
	component, err = versions.ParseComponent(componentData)
	if err != nil {
		return err
	}
	// Ensure we have a key for the component.
	if component.GetKey() == "" {
		component.SetKey(getKey(componentDir))
	}
	// If the component is new, make sure we load the justifications as well.
	if openControl.Components.CompareAndAdd(component) {
		openControl.Justifications.LoadMappings(component)
	}
	return nil
}
