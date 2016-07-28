package docs

import (
	"os"

	"github.com/opencontrol/compliance-masonry/commands/docs/gitbook"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
)

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
