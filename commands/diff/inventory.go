package diff

import (
	"github.com/opencontrol/compliance-masonry/lib"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
	"github.com/opencontrol/compliance-masonry/lib/components/versions/base"
	"fmt"
)

// Inventory maintains the inventory of all the controls within a given workspace.
type Inventory struct {
	lib.Workspace
	masterControlList       map[string]lib.Control
	actualSatisfiedControls map[string]base.Satisfies
	MissingControlList      map[string]lib.Control
}

// retrieveMasterControlsList will gather the list of controls needed for a given certification.
func (i *Inventory) retrieveMasterControlsList() {
	for standardKey, standard := range i.GetCertification().Standards {
		for controlKey, control := range standard.Controls {
			key := standardAndControlString(standardKey, controlKey)
			if _, exists := i.masterControlList[key]; !exists {
				i.masterControlList[key] = control
			}
		}
	}
}

// findDocumentedControls will find the list of all documented controls found within the workspace.
func (i *Inventory) findDocumentedControls() {
	for _, component := range i.GetAllComponents() {
		for _, satisfiedControl := range component.GetAllSatisfies() {
			key := standardAndControlString(satisfiedControl.GetStandardKey(), satisfiedControl.GetControlKey())
			if _, exists := i.actualSatisfiedControls[key]; !exists {
				i.actualSatisfiedControls[key] = satisfiedControl
			}
		}
	}
}

// calculateNonDocumentedControls will compute the diff between the master list of controls and the documented controls.
func (i *Inventory) calculateNonDocumentedControls() {
	for standardAndControlKey, control := range i.masterControlList {
		if _, exists := i.actualSatisfiedControls[standardAndControlKey]; !exists {
			i.MissingControlList[standardAndControlKey] = control
		}
	}
}

// standardAndControlString makes a string from the standard and the control.
// This is helpful for functions that want to create unique keys consistently.
func standardAndControlString(standard string, control string) string {
	return standard + "@" + control
}

// Config contains the settings for how to compute the gap analysis
type Config struct {
	Certification  string
	OpencontrolDir string
}

// ComputeGapAnalysis will compute the gap analysis and return the inventory of the controls for the
// opencontrol workspace if successful. Otherwise, it will return a list of error messages.
// TODO: fix the error return to return of type error. This was used because existing code returned that type
// TODO: e.g. GetCertification
func ComputeGapAnalysis(config Config) (Inventory, []error) {
	// Initialize inventory with certification
	certificationPath, errs := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if certificationPath == "" {
		return Inventory{}, errs
	}
	workspace, _ := lib.LoadData(config.OpencontrolDir, certificationPath)
	i := Inventory{
		Workspace:          workspace,
		masterControlList:       make(map[string]lib.Control),
		actualSatisfiedControls: make(map[string]base.Satisfies),
		MissingControlList:      make(map[string]lib.Control),
	}
	if i.GetCertification() == nil || len(i.GetAllComponents()) == 0 {
		return Inventory{}, []error{fmt.Errorf("Unable to load data in %s for certification %s", config.OpencontrolDir, config.Certification)}
	}

	// Gather list of all controls for certification
	i.retrieveMasterControlsList()
	// Find the documented controls.
	i.findDocumentedControls()
	// Calculate the Missing controls / Non documented
	i.calculateNonDocumentedControls()

	return i, nil
}
