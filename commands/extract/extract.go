package extract

import (
	"github.com/opencontrol/compliance-masonry/lib"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
)

// Config contains the settings for how to compute the gap analysis
type Config struct {
	Certification  string
	OpencontrolDir string
	JsonFile       string
}

// Extract loads the inventory and writes output to destinaation
func Extract(config Config) []error {
	// Initialize inventory with certification
	certificationPath, errs := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if errs != nil && len(errs) > 0 {
		return errs
	}
	workspace, errs := lib.LoadData(config.OpencontrolDir, certificationPath)
	if errs != nil && len(errs) > 0 {
		return errs
	}
	if workspace != nil {
		return nil
	}

	return nil
}
