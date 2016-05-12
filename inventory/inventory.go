package inventory

import (
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
)

type Inventory struct {
	*models.OpenControl
	masterControlList []standardAndControl
	actualSatifiedControls []models.Satisfies
	MissingControlList[]standardAndControl
	ExtraControlList []models.Satisfies
}

type standardAndControl struct {
	control  models.Control
	standard string
}

func(c standardAndControl) String() string {
	return c.standard + "@" + c.control.Name
}

func (c standardAndControl) EqualToSatifiedControl(other models.Satisfies) bool {
	return c.standard == other.StandardKey && c.control.Name == other.ControlKey
}

type Config struct {
	Certification string
	OpencontrolDir string
}

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
		for _, control:= range standard.Controls {
			i.masterControlList = append(i.masterControlList, standardAndControl{standard: certification, control:control})
		}
	}
	// Actual Controls
	for _, components := range i.Components.GetAll() {
		for _, satifiedComponent := range *components.Satisfies {
			i.actualSatifiedControls = append(i.actualSatifiedControls, satifiedComponent)
		}
	}

	// Missing controls
	for _, masterControl := range i.masterControlList {
		found := false
		for _, actualControl := range i.actualSatifiedControls {
			if masterControl.EqualToSatifiedControl(actualControl) {
				found = true
			}
			break
		}
		if !found {
			i.MissingControlList = append(i.MissingControlList, masterControl)
		}
	}

	// Extra controls
	for _, actualControl := range i.actualSatifiedControls {
		found := false
		for _, masterControl := range i.masterControlList {
			if masterControl.EqualToSatifiedControl(actualControl) {
				found = true
			}
			break
		}
		if !found {
			i.ExtraControlList = append(i.ExtraControlList, actualControl)
		}
	}
	return i, nil
}