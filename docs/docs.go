package docs

import (
	"os"

	"github.com/opencontrol/compliance-masonry/docx"
	"github.com/opencontrol/compliance-masonry/gitbook"
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

func MakeGitbook(config gitbook.Config) ([]error) {
	var errs []error
	certificationPath, err := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if certificationPath == "" {
		return append(errs, err)
	}
	//errMessages := cli.NewMultiError()
	if _, err := os.Stat(config.MarkdownPath); os.IsNotExist(err) {
		errs = append(errs, errors.New("Warning: markdown directory does not exist"))
		//errMessages.Errors = append(errMessages.Errors, errors.New("Warning: markdown directory does not exist"))
	}
	config.Certification = certificationPath
	if err := config.BuildGitbook(); err != nil {
		//panic(err.Error())
		return append(errs, err...)
		///errMessages.Errors = append(errMessages.Errors, err)
		//panic(errMessages.Error())
	}
	return nil
}
