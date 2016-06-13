package certifications

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetCertification(opencontrolDir string, certification string) (string, []error) {
	certificationPath := ""
	var errMessages []error
	if certification == "" {
		return "", []error{errors.New("Error: Missing Certification Argument")}
	}
	certificationDir := filepath.Join(opencontrolDir, "certifications")
	certificationPath = filepath.Join(certificationDir, certification+".yaml")
	if _, err := os.Stat(certificationPath); os.IsNotExist(err) {
		files, err := ioutil.ReadDir(certificationDir)
		if err != nil {
			return "", []error{errors.New("Error: `"+certificationDir+"` directory does exist")}
		}
		errMessage := fmt.Sprintf("Error: `%s` does not exist\nUse one of the following:", certificationPath)
		for _, file := range files {
			fileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			errMessage = fmt.Sprintf("%s\n%s", errMessage, fileName)
		}
		errMessages = append(errMessages, errors.New(errMessage))
		return "", errMessages
	}
	return certificationPath, errMessages
}
