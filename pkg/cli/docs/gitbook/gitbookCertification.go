/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package gitbook

import (
	"fmt"

	"path/filepath"

	"github.com/opencontrol/compliance-masonry/internal/constants"
	"github.com/opencontrol/compliance-masonry/internal/utils"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
)

func (openControl *OpenControlGitBook) getResponsibleRole(text string, component common.Component) string {
	if component.GetResponsibleRole() != "" {
		text = fmt.Sprintf("%s\n##### Responsible Role: %s\n", text, component.GetResponsibleRole())
	}
	return text
}

func (openControl *OpenControlGitBook) getNarratives(narratives []common.Section, text string, control *ControlGitbook) string {
	if len(narratives) == 0 {
		return fmt.Sprintf("%s\nNo narrative found for the combination of standard %s and control %s\n", text, control.standardKey, control.controlKey)
	}

	for _, narrative := range narratives {
		text = openControl.getNarrative(narrative, text)
	}
	return text
}

func (openControl *OpenControlGitBook) getNarrative(narrative common.Section, text string) string {
	if narrative.GetKey() != "" {
		text = fmt.Sprintf("%s\n##### %s\n", text, narrative.GetKey())
	}
	text = fmt.Sprintf("%s%s\n", text, narrative.GetText())
	return text
}

func (openControl *OpenControlGitBook) getParameters(text string, parameters []common.Section) string {
	if len(parameters) > 0 {
		text = fmt.Sprintf("%s\n##### Parameters:\n", text)
	}
	for _, parameter := range parameters {
		text = openControl.getParameter(text, parameter)
	}
	return text
}

func (openControl *OpenControlGitBook) getParameter(text string, parameter common.Section) string {
	text = fmt.Sprintf("%s\n###### %s\n", text, parameter.GetKey())
	text = fmt.Sprintf("%s%s\n", text, parameter.GetText())
	return text
}

func (openControl *OpenControlGitBook) getCoveredBy(text string, justification common.Verification) string {
	if len(justification.SatisfiesData.GetCoveredBy()) > 0 {
		text += "Covered By:\n"
	}
	for _, coveredBy := range justification.SatisfiesData.GetCoveredBy() {
		// In case the component key is missing, get it from the justification.
		componentKey := coveredBy.ComponentKey
		if componentKey == "" {
			componentKey = justification.ComponentKey
		}
		component, found := openControl.GetComponent(componentKey)
		if !found {
			continue
		}
		text = openControl.getCoveredByVerification(text, component, coveredBy)
	}
	return text
}

func (openControl *OpenControlGitBook) getCoveredByVerification(text string, component common.Component, coveredBy common.CoveredBy) string {
	if component != nil {
		verification := component.GetVerifications().Get(coveredBy.VerificationKey)
		text += exportLink(
			fmt.Sprintf("%s - %s", component.GetName(), verification.Name),
			filepath.Join("..", "components", component.GetKey()+".md"),
		)
	}
	return text
}

func (openControl *OpenControlGitBook) getControlOrigin(text string, controlOrigin string) string {
	if controlOrigin != "" {
		text = fmt.Sprintf("%s\n##### Control Origin: %s\n", text, controlOrigin)
	}
	return text
}

func (openControl *OpenControlGitBook) exportControl(control *ControlGitbook) (string, string) {
	key := masonryutil.FileNameHandler(fmt.Sprintf("%s-%s", control.standardKey, control.controlKey))
	text := fmt.Sprintf("# %s\n## %s\n", key, control.GetName())
	if len(control.GetDescription()) > 0 {
		text += "#### Description\n"
		text += control.GetDescription()
	}
	selectJustifications := openControl.GetAllVerificationsWith(control.standardKey, control.controlKey)
	// In the case that no information was found period for the standard and control
	if len(selectJustifications) == 0 {
		errorText := fmt.Sprintf("No information found for the combination of standard %s and control %s", control.standardKey, control.controlKey)
		text = fmt.Sprintf("%s\n%s\n", text, errorText)
	}
	for _, justification := range selectJustifications {
		component, found := openControl.GetComponent(justification.ComponentKey)
		if !found {
			continue
		}
		text = fmt.Sprintf("%s\n#### %s\n", text, component.GetName())

		text = openControl.getResponsibleRole(text, component)

		text = openControl.getParameters(text, justification.SatisfiesData.GetParameters())

		text = openControl.getControlOrigin(text, justification.SatisfiesData.GetControlOrigin())

		text = openControl.getNarratives(justification.SatisfiesData.GetNarratives(), text, control)
		text = openControl.getCoveredBy(text, justification)
	}
	return filepath.Join(control.exportPath, key+".md"), text
}

func (openControl *OpenControlGitBook) exportStandards() {
	standardsExportPath := filepath.Join(openControl.exportPath, "standards")
	openControl.FSUtil.Mkdirs(standardsExportPath)
	standardKeys := openControl.GetCertification().GetSortedStandards()
	for _, standardKey := range standardKeys {
		standard, found := openControl.GetStandard(standardKey)
		if !found {
			continue
		}
		controlKeys := standard.GetSortedControls()
		for _, controlKey := range controlKeys {
			control := standard.GetControl(controlKey)
			controlPath, controlText := openControl.exportControl(&ControlGitbook{control, standardsExportPath, standardKey, controlKey})
			masonryutil.FileWriter(controlPath, []byte(controlText), constants.FileReadWrite)
		}
	}
}
