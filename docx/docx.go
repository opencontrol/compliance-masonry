package docx

import (
	"strings"

	"github.com/opencontrol/compliance-masonry-go/models"
)

// Config contains data for docx template export configurations
type Config struct {
	OpencontrolDir string
	TemplatePath   string
	ExportPath     string
}

// OpenControlDocx struct is an extension of models.OpenControl that adds a
// template path and export path.
type OpenControlDocx struct {
	*models.OpenControl
	templatePath string
	exportPath   string
}

//BuildDocx exports a Doxc ssp based on a template
func (config *Config) BuildDocx() {
	openControl := OpenControlDocx{
		models.LoadData("../fixtures/opencontrol_fixtures/", ""),
		"",
		"",
	}
	openControl.LoadCertification("")
}

// getControl returns a control formatted for docx
func (opencontrol *OpenControlDocx) formatControl(standardControl string) string {
	var output string
	standardKey, controlKey := splitControl(standardControl)
	opencontrol.Justifications.GetAndApply(standardKey, controlKey, func(selectJustifications models.Verifications) {
		for _, justification := range selectJustifications {
			output += justification.SatisfiesData.Narrative
		}
	})
	return output
}

// splitControl returns a split standard and control given a standard
// and control delimited with `@`
func splitControl(standardControl string) (string, string) {
	var standard, control string
	splitString := strings.Split(standardControl, "@")
	splitStringLen := len(splitString)
	switch {
	case splitStringLen >= 2:
		standard = splitString[0]
		control = splitString[1]

	case splitStringLen == 1:
		standard = splitString[0]
	}
	return standard, control

}
