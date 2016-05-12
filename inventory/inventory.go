package inventory

import (
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
)

// Inventory maintains the inventory of all the controls within a given workspace.
type Inventory struct {
	*models.OpenControl
	masterControlList      []standardAndControl
	actualSatifiedControls []models.Satisfies
	MissingControlList     []standardAndControl
}

type standardAndControl struct {
	control  string
	standard string
}

func (c standardAndControl) String() string {
	return c.standard + "@" + c.control
}

func (c standardAndControl) EqualToSatisfiedControl(other models.Satisfies) bool {
	return c.standard == other.StandardKey && c.control == other.ControlKey
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
	i := Inventory{OpenControl: models.LoadData(config.OpencontrolDir, certificationPath)}
	if i.Certification == nil || i.Components == nil {
		return Inventory{}, []string{"Unable to load data in " + config.OpencontrolDir + " for certification " + config.Certification}
	}
	// Master Controls
	for certification, standard := range i.Certification.Standards {
		for control, _ := range standard.Controls {
			i.masterControlList = append(i.masterControlList, standardAndControl{standard: certification, control: control})
		}
	}
	// Actual Controls
	for _, components := range i.Components.GetAll() {
		for _, satisfiedComponent := range *components.Satisfies {
			i.actualSatifiedControls = append(i.actualSatifiedControls, satisfiedComponent)
		}
	}

	// Missing controls
	for _, masterControl := range i.masterControlList {
		found := false
		for _, actualControl := range i.actualSatifiedControls {
			if masterControl.EqualToSatisfiedControl(actualControl) {
				found = true
				break
			}
		}
		if !found {
			i.MissingControlList = append(i.MissingControlList, masterControl)
		}
	}
	return i, nil
}
