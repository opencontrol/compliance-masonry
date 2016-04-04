package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/opencontrol/compliance-masonry-go/docx"
	"github.com/opencontrol/compliance-masonry-go/gitbook"
)

func getCertification(opencontrolDir string, certification string) (string, []string) {
	var (
		certificationPath string
		messages          []string
	)
	if certification == "" {
		messages = append(messages, "Error: Missing Certification Argument")
		return "", messages
	}
	certificationDir := filepath.Join(opencontrolDir, "certifications")
	certificationPath = filepath.Join(certificationDir, certification+".yaml")
	if _, err := os.Stat(certificationPath); os.IsNotExist(err) {
		files, err := ioutil.ReadDir(certificationDir)
		if err != nil {
			messages = append(messages, "Error: `opencontrols/certifications` directory does exist")
			return "", messages
		}
		messages = append(messages, fmt.Sprintf("Error: `%s` does not exist\nUse one of the following:", certificationPath))
		for _, file := range files {
			fileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			messages = append(messages, fmt.Sprintf("`%s`", fileName))
		}
		return "", messages
	}
	return certificationPath, messages
}

func buildTemplate(config *docx.Config) []string {
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

func makeGitbook(config *gitbook.Config) []string {
	certificationPath, messages := getCertification(config.OpencontrolDir, config.Certification)
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
