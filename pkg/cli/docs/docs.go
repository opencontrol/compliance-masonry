package docs

import (
	"os"

	"github.com/opencontrol/compliance-masonry/pkg/cli/docs/gitbook"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
)

// MakeGitbook is the wrapper function that will create a gitbook for the specified certification.
func MakeGitbook(config gitbook.Config) (string, []error) {
	warning := ""
	certificationPath, err := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if certificationPath == "" {
		return warning, err
	}
	if _, err := os.Stat(config.MarkdownPath); os.IsNotExist(err) {
		warning = "Warning: markdown directory does not exist"
	}
	config.Certification = certificationPath
	if err := config.BuildGitbook(); err != nil {
		return warning, err
	}
	return warning, nil
}
