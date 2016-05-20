package docx

import (
	"fmt"
	"text/template"

	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/doc-template"
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
	funcMap := template.FuncMap{
		"getAllControlSections": openControl.FormatControl,
		"getControlSection":     openControl.FormatControlSection,
	}
	docTemplate.AddFunctions(funcMap)
	docTemplate.Parse()
	docTemplate.Execute(config.ExportPath, nil)
	return err
}

func (openControl *OpenControlDocx) FormatControlSection(standardKey string, controlKey string, sectionKey string) string {
	var text string
	openControl.Justifications.GetAndApply(standardKey, controlKey, func(selectJustifications models.Verifications) {
		for _, justification := range selectJustifications {
			openControl.Components.GetAndApply(justification.ComponentKey, func(component *models.Component) {

				// Print out the narrative(s)
				found := false
				var narrativeText string
				for _, section := range justification.SatisfiesData.Narrative {
					// If section header exists, let's print it's corresponding text and not the header itself.
					if section.Key == sectionKey {
						narrativeText = fmt.Sprintf("%s%s\n", narrativeText, section.Text)
						found = true
					}
				}
				// If we actually had a section, let's get the component and add the narrative
				if found {
					// Print out the component name.
					text = fmt.Sprintf("%s%s\n%s", text, component.Name, narrativeText)
				}
			})
		}
	})
	return text
}

// FormatControl returns a control formatted for docx
func (openControl *OpenControlDocx) FormatControl(standardKey string, controlKey string) string {
	var text string
	openControl.Justifications.GetAndApply(standardKey, controlKey, func(selectJustifications models.Verifications) {
		for _, justification := range selectJustifications {
			openControl.Components.GetAndApply(justification.ComponentKey, func(component *models.Component) {
				// Print out the component name.
				text = fmt.Sprintf("%s%s\n", text, component.Name)
				// Print out the narrative(s)
				for _, section := range justification.SatisfiesData.Narrative {
					// If section header exists, let's print it. Key could be empty, in that case
					// just print the text for the section.
					if section.Key != "" {
						text = fmt.Sprintf("%s%s:\n", text, section.Key)
					}
					text = fmt.Sprintf("%s%s\n", text, section.Text)
				}
			})
		}
	})
	return text
}
