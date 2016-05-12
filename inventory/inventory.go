package inventory

import (
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
	"fmt"
)

type Inventory struct {
	*models.OpenControl
}

type Config struct {
	Certification string
	OpencontrolDir string
}

func ComputeGapAnalysis(config Config) ([]interface{}, []string) {
	certificationPath, messages := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if certificationPath == "" {
		return nil, messages
	}
	i := Inventory{models.LoadData(config.OpencontrolDir, certificationPath)}
	if i.Certification == nil || i.Components == nil {
		return nil, []string{"Unable to load data in " + config.OpencontrolDir + " for certification " + config.Certification}
	}
	// Master Controls
	for k1, v1 := range i.Certification.Standards {
		for k2, v2 := range v1.Controls {
			fmt.Println(k1 + " --- " + k2 + " --- " + v2.Name)
		}
	}
	// Actual Controls
	for _, v1 := range i.Components.GetAll() {
		for _, v2 := range *v1.Satisfies {
			fmt.Println(" +++ " + v2.ControlKey)
		}
	}
	return nil, nil
}