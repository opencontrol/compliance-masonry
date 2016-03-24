package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/opencontrol/compliance-masonry-go/gitbook"
)

type gitbookConfig struct {
	certification  string
	opencontrolDir string
	exportPath     string
	markdownPath   string
}

func (config *gitbookConfig) makeGitbook() []string {
	var messages []string
	if config.certification == "" {
		messages = append(messages, "Error: New Missing Certification Argument")
		messages = append(messages, "Usage: masonry-go docs gitbook FedRAMP-low")
		return messages
	}
	certificationDir := filepath.Join(config.opencontrolDir, "certifications")
	certificationPath := filepath.Join(certificationDir, config.certification+".yaml")
	if _, err := os.Stat(certificationPath); os.IsNotExist(err) {
		files, err := ioutil.ReadDir(certificationDir)
		if err != nil {
			messages = append(messages, "Error: `opencontrols/certifications` directory does exist")
			return messages
		}
		messages = append(messages, fmt.Sprintf("Error: `%s` does not exist\nUse one of the following:", certificationPath))
		for _, file := range files {
			fileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			messages = append(messages, fmt.Sprintf("`compliance-masonry-go docs gitbook %s`", fileName))
		}
		return messages
	}
	if _, err := os.Stat(config.markdownPath); os.IsNotExist(err) {
		markdownPath = ""
		messages = append(messages, "Warning: markdown directory does not exist")
	}
	gitbook.BuildGitbook(config.opencontrolDir, certificationPath, markdownPath, config.exportPath)
	messages = append(messages, "New Gitbook Documentation Created")
	return messages
}
