package inventory

import (
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
)

// Inventory maintains the inventory of all the controls within a given workspace.
type Inventory struct {
	*models.OpenControl
	masterControlList       map[string]models.Control
	actualSatisfiedControls map[string]models.Satisfies
	MissingControlList      map[string]models.Control
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
func ComputeGapAnalysis(config Config) (Inventory, []string) {
	certificationPath, messages := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if certificationPath == "" {
		return Inventory{}, messages
	}
	i := Inventory{
		OpenControl:             models.LoadData(config.OpencontrolDir, certificationPath),
		masterControlList:       make(map[string]models.Control),
		actualSatisfiedControls: make(map[string]models.Satisfies),
		MissingControlList:      make(map[string]models.Control),
	}
	if i.Certification == nil || i.Components == nil {
		return Inventory{}, []string{"Unable to load data in " + config.OpencontrolDir + " for certification " + config.Certification}
	}
	// Master Controls
	for standardKey, standard := range i.Certification.Standards {
		for controlKey, control := range standard.Controls {
			key := standardAndControlString(standardKey, controlKey)
			if _, exists := i.masterControlList[key]; !exists {
				i.masterControlList[key] = control
			}
		}
	}
	// Actual Controls
	for _, components := range i.Components.GetAll() {
		for _, satisfiedComponent := range *components.Satisfies {
			key := standardAndControlString(satisfiedComponent.StandardKey, satisfiedComponent.ControlKey)
			if _, exists := i.actualSatisfiedControls[key]; !exists {
				i.actualSatisfiedControls[key] = satisfiedComponent
			}
		}
	}

	// Missing controls
	for standardAndControlKey, control := range i.masterControlList {
		if _, exists := i.actualSatisfiedControls[standardAndControlKey]; !exists {
			i.MissingControlList[standardAndControlKey] = control
		}
	}
	return i, nil
}
