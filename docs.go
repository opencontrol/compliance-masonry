package main

import (
	"os"

	"github.com/opencontrol/compliance-masonry/docx"
	"github.com/opencontrol/compliance-masonry/gitbook"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
)

func BuildTemplate(config docx.Config) []string {
	var messages []string
	if config.TemplatePath == "" {
		messages = append(messages, "Error: No Template Supplied")
		return messages
	}
	if _, err := os.Stat(config.TemplatePath); os.IsNotExist(err) {
		messages = append(messages, "Error: Template does not exist")
		return messages
	}
	err := config.BuildDocx()
	if err != nil {
		messages = append(messages, err.Error())
	} else {
		messages = append(messages, "New Docx Created")
	}
	return messages
}

func MakeGitbook(config gitbook.Config) []string {
	certificationPath, messages := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if certificationPath == "" {
		return messages
	}
	if _, err := os.Stat(config.MarkdownPath); os.IsNotExist(err) {
		markdownPath = ""
		messages = append(messages, "Warning: markdown directory does not exist")
	}
	config.Certification = certificationPath
	config.BuildGitbook()
	messages = append(messages, "New Gitbook Documentation Created")
	return messages
}
