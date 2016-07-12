package docs

import (
	"os"

	"github.com/opencontrol/compliance-masonry/commands/docs/docx"
	"github.com/opencontrol/compliance-masonry/commands/docs/gitbook"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
	"errors"
)

func BuildTemplate(config docx.Config) error {
	if config.TemplatePath == "" {
		return errors.New("Error: No Template Supplied")
	}
	if _, err := os.Stat(config.TemplatePath); os.IsNotExist(err) {
		return errors.New("Error: Template does not exist")
	}
	err := config.BuildDocx()
	if err != nil {
		return err
	}
	return nil
}

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
