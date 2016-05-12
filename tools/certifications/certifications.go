package certifications

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func GetCertification(opencontrolDir string, certification string) (string, []string) {
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
