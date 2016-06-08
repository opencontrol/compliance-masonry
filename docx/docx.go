package docx

import (
	"fmt"
	"text/template"

	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/doc-template"

	"gopkg.in/fatih/set.v0"
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
		"getParameter":          openControl.FormatParameter,
		"getResponsibleRole":    openControl.FormatResponsibleRoles,
	}
	docTemplate.AddFunctions(funcMap)
	docTemplate.Parse()
	docTemplate.Execute(config.ExportPath, nil)
	return err
}

type componentInfoType int

const (
	noneInfo componentInfoType = iota
	controlInfo
	parameterInfo
	responsibleRoleInfo
)

// createSectionsSet creates a set of section headers to do easy searching from the slice of string section keys
func createSectionsSet(sections ...string) *set.Set {
	sectionsSet := set.New()
	for _, section := range sections {
		sectionsSet.Add(section)
	}
	return sectionsSet
}

func getControlInfo(text string, justification models.Verification, component *models.Component, specifiedSections *set.Set) (string, bool) {
	// Add the component name.
	text = fmt.Sprintf("%s%s\n", text, component.Name)

	// foundText is a placeholder to indicate that we actually found text for the section.
	foundText := false

	// Determine if we want to get all of the sections or just one. If we specify exact sections, that means we do
	// not want all and if we do not specify sections, it means we want all sections.
	allSections := specifiedSections.Size() == 0
	// Print out the narrative(s)
	for _, section := range justification.SatisfiesData.Narrative {
		if allSections {
			// If we want to print out all the sections...

			// If section header exists, let's print it. Key could be empty, in that case
			// just print the text for the section.
			if section.Key != "" {
				text = fmt.Sprintf("%s%s:\n", text, section.Key)
			}
			text = fmt.Sprintf("%s%s\n", text, section.Text)

			// Automatically assume foundText is true as long as the length of
			// justification.SatisfiesData.Narrative is > 0, which is implied if we reach here.
			// Also, in case the section in the YAML is explicitly "", we accept empty string here too.
			foundText = true
		} else {
			// If we only want certain section(s)...

			// If section header exists, let's print it's corresponding text and not the header itself.
			if specifiedSections.Has(section.Key) {
				text = fmt.Sprintf("%s%s\n", text, section.Text)
				foundText = true
			}
		}
	}
	return text, foundText
}

func getParameterInfo(text string, justification models.Verification, component *models.Component, specifiedSections *set.Set) (string, bool) {
	// Add the component name.
	text = fmt.Sprintf("%s%s\n", text, component.Name)

	// foundText is a placeholder to indicate that we actually found text for the section.
	foundText := false

	for _, parameter := range justification.SatisfiesData.Parameters {
		// If section header exists, let's print it's corresponding text and not the header itself.
		if specifiedSections.Has(parameter.Key){
			text = fmt.Sprintf("%s%s\n", text, parameter.Text)
			foundText = true
		}
	}
	return text, foundText
}

func getResponsibleRoleInfo(text string, justification models.Verification, component *models.Component, specifiedSections *set.Set) (string, bool) {
	// Add the component name.
	text = fmt.Sprintf("%s%s: ", text, component.Name)
	// Print out the component name and the responsible for that component.
	if justification.SatisfiesData.ResponsibleRole != "" {
		return fmt.Sprintf("%s%s\n", text, justification.SatisfiesData.ResponsibleRole), true
	} else {
		return text, false
	}
}

// getComponentControlText will get the appropriate control text that is formatted for the word document.
func (openControl *OpenControlDocx) getComponentText(infoType componentInfoType, standardKey string, controlKey string, sectionKeys ...string) string {

	var text string
	sectionSet := createSectionsSet(sectionKeys...)

	openControl.Justifications.GetAndApply(standardKey, controlKey, func(selectJustifications models.Verifications) {
		// In the case that no information was found period for the standard and control
		if len(selectJustifications) == 0 {
			text = fmt.Sprintf(constants.WarningUnknownStandardAndControlf, standardKey, controlKey)
			return
		}
		for _, justification := range selectJustifications {
			openControl.Components.GetAndApply(justification.ComponentKey, func(component *models.Component) {
				// Get the Component Text
				var specificText string
				var found bool
				switch(infoType) {
				case controlInfo:
					specificText, found = getControlInfo(text, justification, component, sectionSet)
				case parameterInfo:
					specificText, found = getParameterInfo(text, justification, component, sectionSet)
				case responsibleRoleInfo:
					specificText, found = getResponsibleRoleInfo(text, justification, component, sectionSet)
				}
				if found {
					text = fmt.Sprintf("%s%s", text, specificText)
				} else {
					text = fmt.Sprintf("%s%s\n", text, constants.WarningNoInformationAvailable)
				}
			})
		}
	})

	return text
}

// FormatResponsibleRole fills in the responsible role for each component for a given standard and control.
func (openControl *OpenControlDocx) FormatResponsibleRoles(standardKey string, controlKey string) string {
	return openControl.getComponentText(responsibleRoleInfo, standardKey, controlKey, "")
}

// FormatParameter fills in the parameter for a given parameter, standard and control.
func (openControl *OpenControlDocx) FormatParameter(standardKey string, controlKey string, sectionKeys ...string) string {
	return openControl.getComponentText(parameterInfo, standardKey, controlKey, sectionKeys...)
}

// FormatControl returns a control formatted for docx
func (openControl *OpenControlDocx) FormatControl(standardKey string, controlKey string, sectionKeys ...string) string {
	return openControl.getComponentText(controlInfo, standardKey, controlKey, sectionKeys...)
}
