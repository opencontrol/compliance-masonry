package docx

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/opencontrol/doc-template"
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
}

//BuildDocx exports a Doxc ssp based on a template
func (config *Config) BuildDocx() error {
	openControl := OpenControlDocx{models.LoadData(config.OpencontrolDir, "")}
	docTemplate, err := docTemp.GetTemplate(config.TemplatePath)
	if err != nil {
		return err
	}
	funcMap := template.FuncMap{"getControl": openControl.FormatControl}
	docTemplate.AddFunctions(funcMap)
	docTemplate.Parse()
	docTemplate.Execute(config.ExportPath, nil)
	return err
}

// FormatControl returns a control formatted for docx
func (openControl *OpenControlDocx) FormatControl(standardControl string) string {
	var text string
	standardKey, controlKey := SplitControl(standardControl)
	openControl.Justifications.GetAndApply(standardKey, controlKey, func(selectJustifications models.Verifications) {
		for _, justification := range selectJustifications {
			openControl.Components.GetAndApply(justification.ComponentKey, func(component *models.Component) {
				text = fmt.Sprintf("%s%s  \n", text, component.Name)
				text = fmt.Sprintf("%s%s  \n", text, justification.SatisfiesData.Narrative)
			})

			if len(justification.SatisfiesData.CoveredBy) > 0 {
				text += "Covered By:  \n"
			}

			for _, coveredBy := range justification.SatisfiesData.CoveredBy {
				componentKey := coveredBy.ComponentKey
				if componentKey == "" {
					componentKey = justification.ComponentKey
				}
				openControl.Components.GetAndApply(componentKey, func(component *models.Component) {
					if component != nil {
						verification := component.Verifications.Get(coveredBy.VerificationKey)
						text += fmt.Sprintf("- %s %s  \n", verification.Name, verification.Path)
					}
				})
			}
		}
	})
	return text
}

// SplitControl returns a split standard and control given a standard
// and control delimited with `@`
func SplitControl(standardControl string) (string, string) {
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
