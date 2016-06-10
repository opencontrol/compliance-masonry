package docx

import (
	"fmt"
	"text/template"

	"github.com/opencontrol/doc-template"
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"gopkg.in/fatih/set.v0"
	"github.com/opencontrol/compliance-masonry/tools/constants"
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

//BuildDocx exports a Docx ssp based on a template
func (config *Config) BuildDocx() error {
	openControl := OpenControlDocx{models.LoadData(config.OpencontrolDir, "")}
	docTemplate, err := docTemp.GetTemplate(config.TemplatePath)
	if err != nil {
		return err
	}
	funcMap := template.FuncMap{
		"getAllControlSections": openControl.FormatAllNarratives,
		"getControlSection":     openControl.FormatNarrative,
		"getParameter":          openControl.FormatParameter,
		"getResponsibleRole":    openControl.FormatResponsibleRoles,
	}
	docTemplate.AddFunctions(funcMap)
	docTemplate.Parse()
	docTemplate.Execute(config.ExportPath, nil)
	return err
}

type controlInfoType int

const (
	// placeholder for the default case.
	noneInfo controlInfoType = iota
	// represents the request for all the narratives for a control.
	allControlInfo
	// represents the request for specific narrative(s) for a control.
	controlInfo
	// represents the request for specific parameter(s) for a control.
	parameterInfo
	// represents the request for the responsible role for a control..
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

// getNarrativeSection will just print the narrative section text. No need to print the section header since it was specified.
func getNarrativeSection(text string, justification models.Verification, component base.Component, specifiedSections *set.Set) (string) {
	// Add the component name.
	text = fmt.Sprintf("%s%s\n", text, component.GetName())

	// Use generic []base.Section handler.
	return getSpecificGenericSections(justification.SatisfiesData.GetNarratives(), text, specifiedSections)
}

// getAllNarrativeSection will print both the section header and the section text for all narrative sections.
func getAllNarrativeSections(text string, justification models.Verification, component base.Component) (string) {
	// Add the component name.
	text = fmt.Sprintf("%s%s\n", text, component.GetName())
	for _, section := range justification.SatisfiesData.GetNarratives() {
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
	}
	return text

}

// getSpecificGenericSections can be used by both narrative and parameter since they both implement base.Section
func getSpecificGenericSections(sections []base.Section, text string, specifiedSections *set.Set) (string) {
	// In the case that the user does not provide any sections.
	if specifiedSections.Size() == 0 {
		return fmt.Sprintf("%s%s\n", text, constants.WarningNoInformationAvailable)
	}
	for _, section := range sections {
		// If we only want certain section(s)...

		// If section header exists, let's print it's corresponding text and not the header itself.
		if specifiedSections.Has(section.GetKey()) {
			text = fmt.Sprintf("%s%s\n", text, section.GetText())
			specifiedSections.Remove(section.GetKey())
		}
	}
	// In the case that we do not have the section, print warning that information was not found.
	if specifiedSections.Size() != 0 {
		text = fmt.Sprintf("%s%s\n", text, constants.WarningNoInformationAvailable)
	}
	return text
}

// getParameterInfo will just print the parameter section text. No need to print the section header since it was specified.
func getParameterInfo(text string, justification models.Verification, component base.Component, specifiedSections *set.Set) (string) {
	// Add the component name.
	text = fmt.Sprintf("%s%s\n", text, component.GetName())

	// Use generic []base.Section handler.
	return getSpecificGenericSections(justification.SatisfiesData.GetParameters(), text, specifiedSections)
}

// getResponsibleRoleInfo will just print the responsible role if it exists.
func getResponsibleRoleInfo(text string, component base.Component) (string) {
	// Add the component name.
	text = fmt.Sprintf("%s%s: ", text, component.GetName())
	// Print out the component name and the responsible for that component.
	if component.GetResponsibleRole() != "" {
		return fmt.Sprintf("%s%s\n", text, component.GetResponsibleRole())
	}
	// Else, print warning indicating there was no info.
	return fmt.Sprintf("%s%s\n", text, constants.WarningNoInformationAvailable)
}

// getComponentText is for information that will need to dig into the justifications.
func (openControl *OpenControlDocx) getComponentText(infoType controlInfoType, standardKey string, controlKey string, sectionKeys ...string) string {
	var text string
	sectionSet := createSectionsSet(sectionKeys...)
	openControl.Justifications.GetAndApply(standardKey, controlKey, func(selectJustifications models.Verifications) {
		// In the case that no information was found period for the standard and control
		if len(selectJustifications) == 0 {
			text = fmt.Sprintf("No information found for the combination of standard %s and control %s", standardKey, controlKey)
			return
		}

		for _, justification := range selectJustifications {
			openControl.Components.GetAndApply(justification.ComponentKey, func(component base.Component) {
				// Get the Component Text
				switch(infoType) {
				case allControlInfo:
					text = fmt.Sprintf("%s%s", text, getAllNarrativeSections(text, justification, component))
				case controlInfo:
					text = fmt.Sprintf("%s%s", text, getNarrativeSection(text, justification, component, sectionSet))
				case parameterInfo:
					text = fmt.Sprintf("%s%s", text, getParameterInfo(text, justification, component, sectionSet))
				case responsibleRoleInfo:
					text = fmt.Sprintf("%s%s", text, getResponsibleRoleInfo(text, component))
				}
			})
		}
	})

	return text
}

// FormatResponsibleRoles fills in the responsible role for each component for a given standard and control.
func (openControl *OpenControlDocx) FormatResponsibleRoles(standardKey string, controlKey string) string {
	return openControl.getComponentText(responsibleRoleInfo, standardKey, controlKey, "")
}

// FormatParameter fills in the parameter for a given parameter, standard and control.
func (openControl *OpenControlDocx) FormatParameter(standardKey string, controlKey string, sectionKeys ...string) string {
	return openControl.getComponentText(parameterInfo, standardKey, controlKey, sectionKeys...)
}

// FormatAllNarratives returns a control formatted for docx with all the narratives
func (openControl *OpenControlDocx) FormatAllNarratives(standardKey string, controlKey string) string {
	return openControl.getComponentText(allControlInfo, standardKey, controlKey, "")
}

// FormatNarrative returns a control formatted for docx with only the specified narrative section(s)
func (openControl *OpenControlDocx) FormatNarrative(standardKey string, controlKey string, sectionKeys ...string) string {
	return openControl.getComponentText(controlInfo, standardKey, controlKey, sectionKeys...)
}