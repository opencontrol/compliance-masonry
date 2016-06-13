package certifications

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"github.com/codegangsta/cli"
	"github.com/vektra/errors"
)

func GetCertification(opencontrolDir string, certification string) (string, error) {
	certificationPath := ""
	errMessages := cli.NewMultiError()
	if certification == "" {
		return "", errors.New("Error: Missing Certification Argument")
	}
	certificationDir := filepath.Join(opencontrolDir, "certifications")
	certificationPath = filepath.Join(certificationDir, certification+".yaml")
	if _, err := os.Stat(certificationPath); os.IsNotExist(err) {
		files, err := ioutil.ReadDir(certificationDir)
		if err != nil {
			errMessages.Errors = append(errMessages.Errors, errors.New("Error: `"+certificationDir+"` directory does exist"))
			return "", errMessages
		}
		errMessage := fmt.Sprintf("Error: `%s` does not exist\nUse one of the following:", certificationPath)
		for _, file := range files {
			fileName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			errMessage = fmt.Sprintf("%s\n`%s`", errMessages, fileName)
		}
		errMessages.Errors = append(errMessages.Errors, errors.New(errMessage))
		return "", errMessages
	}
	return certificationPath, errMessages
}
