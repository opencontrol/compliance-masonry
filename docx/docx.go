package docx

import (
	"fmt"
	"text/template"

	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/doc-template"

	"gopkg.in/fatih/set.v0"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
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
		"getControlSection":     openControl.FormatControl,
	}
	docTemplate.AddFunctions(funcMap)
	docTemplate.Parse()
	docTemplate.Execute(config.ExportPath, nil)
	return err
}

// createSectionsSet creates a set of section headers to do easy searching from the slice of string section keys
func createSectionsSet(sections ...string) *set.Set {
	sectionsSet := set.New()
	for _, section := range sections {
		sectionsSet.Add(section)
	}
	return sectionsSet
}

// getComponentControlText will get the appropriate control text that is formatted for the word document.
func getComponentControlText(justification models.Verification, component base.Component, specifiedSections *set.Set) string {
	var text string

	// Add the component name.
	text = fmt.Sprintf("%s%s\n", text, component.GetName())

	// Determine if we want to get all of the sections or just one. If we specify exact sections, that means we do
	// not want all and if we do not specify sections, it means we want all sections.
	allSections := specifiedSections.Size() == 0

	// foundText is a placeholder to indicate that we actually found text for the section.
	foundText := false

	// Print out the narrative(s)
	for _, section := range justification.SatisfiesData.GetNarratives() {
		if allSections {
			// If we want to print out all the sections...

			// If section header exists, let's print it. Key could be empty, in that case
			// just print the text for the section.
			if section.GetKey() != "" {
				text = fmt.Sprintf("%s%s:\n", text, section.GetKey())
			}
			text = fmt.Sprintf("%s%s\n", text, section.GetText())

			// Automatically assume foundText is true as long as the length of
			// justification.SatisfiesData.Narrative is > 0, which is implied if we reach here.
			// Also, in case the section in the YAML is explicitly "", we accept empty string here too.
			foundText = true
		} else {
			// If we only want certain section(s)...

			// If section header exists, let's print it's corresponding text and not the header itself.
			if specifiedSections.Has(section.GetKey()) {
				text = fmt.Sprintf("%s%s\n", text, section.GetText())
				foundText = true
			}
		}
	}

	if !foundText {
		text = fmt.Sprintf("%s%s\n", text, constants.WarningNoInformationAvailable)
	}

	return text
}

// FormatControl returns a control formatted for docx
func (openControl *OpenControlDocx) FormatControl(standardKey string, controlKey string, sectionKeys ...string) string {
	sectionSet := createSectionsSet(sectionKeys...)
	var text string
	openControl.Justifications.GetAndApply(standardKey, controlKey, func(selectJustifications models.Verifications) {
		// In the case that no information was found period for the standard and control
		if len(selectJustifications) == 0 {
			text = fmt.Sprintf(constants.WarningUnknownStandardAndControlf, standardKey, controlKey)
			return
		}
		for _, justification := range selectJustifications {
			openControl.Components.GetAndApply(justification.ComponentKey, func(component base.Component) {
				// Get the Component Text
				text = fmt.Sprintf("%s%s", text, getComponentControlText(justification, component, sectionSet))
			})
		}
	})
	return text
}
