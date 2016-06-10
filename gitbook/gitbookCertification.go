package gitbook

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"github.com/opencontrol/compliance-masonry/models/common"
)

func (openControl *OpenControlGitBook) getResponsibleRole(text string, component base.Component) string {
	if component.GetResponsibleRole() != "" {
		text = fmt.Sprintf("%s\n##### Responsible Role: %s\n", text, component.GetResponsibleRole())
	}
	return text
}

func (openControl *OpenControlGitBook) getNarratives(narratives []base.Section, text string, control *ControlGitbook) string {
	if len(narratives) == 0 {
		return fmt.Sprintf("%sNo narrative found for the combination of standard %s and control %s\n", text, control.standardKey, control.controlKey)
	}

	for _, narrative := range narratives {
		text = openControl.getNarrative(narrative, text)
	}
	return text
}

func (openControl *OpenControlGitBook) getNarrative(narrative base.Section, text string) string {
	if narrative.GetKey() != "" {
		text = fmt.Sprintf("%s\n##### %s\n", text, narrative.GetKey())
	}
	text = fmt.Sprintf("%s%s\n", text, narrative.GetText())
	return text
}

func (openControl *OpenControlGitBook) getCoveredBy(text string, justification models.Verification) string {
	if len(justification.SatisfiesData.GetCoveredBy()) > 0 {
		text += "Covered By:\n"
	}
	for _, coveredBy := range justification.SatisfiesData.GetCoveredBy() {
		// In case the component key is missing, get it from the justification.
		componentKey := coveredBy.ComponentKey
		if componentKey == "" {
			componentKey = justification.ComponentKey
		}
		openControl.Components.GetAndApply(componentKey, func(component base.Component) {
			text = openControl.getCoveredByVerification(text, component, coveredBy)
		})
	}
	return text
}

func (openControl *OpenControlGitBook) getCoveredByVerification(text string, component base.Component, coveredBy common.CoveredBy) string {
	if component != nil {
		verification := component.GetVerifications().Get(coveredBy.VerificationKey)
		text += exportLink(
			fmt.Sprintf("%s - %s", component.GetName(), verification.Name),
			filepath.Join("..", "components", component.GetKey()+".md"),
		)
	}
	return text
}



func (openControl *OpenControlGitBook) exportControl(control *ControlGitbook) (string, string) {
	key := replaceParentheses(fmt.Sprintf("%s-%s", control.standardKey, control.controlKey))
	text := fmt.Sprintf("#%s\n##%s\n", key, control.Name)
	openControl.Justifications.GetAndApply(control.standardKey, control.controlKey, func(selectJustifications models.Verifications) {
		// In the case that no information was found period for the standard and control
		if len(selectJustifications) == 0 {
			errorText := fmt.Sprintf("No information found for the combination of standard %s and control %s", control.standardKey, control.controlKey)
			text = fmt.Sprintf("%s%s\n", text, errorText)
			return
		}
		for _, justification := range selectJustifications {
			openControl.Components.GetAndApply(justification.ComponentKey, func(component base.Component) {
				text = fmt.Sprintf("%s\n#### %s\n", text, component.GetName())

				text = openControl.getResponsibleRole(text, component)

				text = openControl.getNarratives(justification.SatisfiesData.GetNarratives(), text, control)
			})
			text = openControl.getCoveredBy(text, justification)
		}
	})
	return filepath.Join(control.exportPath, key+".md"), text
}

func (openControl *OpenControlGitBook) exportStandards() {
	standardsExportPath := filepath.Join(openControl.exportPath, "standards")
	openControl.FSUtil.Mkdirs(standardsExportPath)
	openControl.Certification.GetSortedData(func(standardKey string, controlKey string) {
		control := openControl.Standards.Get(standardKey).Controls[controlKey]
		controlPath, controlText := openControl.exportControl(&ControlGitbook{&control, standardsExportPath, standardKey, controlKey})
		ioutil.WriteFile(controlPath, []byte(controlText), 0700)
	})
}
