package gitbook

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
)

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

				if component.GetResponsibleRole() != "" {
					text = fmt.Sprintf("%s\n##### Responsible Role: %s\n", text, component.GetResponsibleRole())
				}

				if len(justification.SatisfiesData.GetNarratives()) == 0 {
					text = fmt.Sprintf("%sNo narrative found for the combination of standard %s and control %s\n", text, control.standardKey, control.controlKey)
					return
				}

				for _, narrative := range justification.SatisfiesData.GetNarratives() {
					if narrative.GetKey() != "" {
						text = fmt.Sprintf("%s\n##### %s\n", text, narrative.GetKey())
					}
					text = fmt.Sprintf("%s%s\n", text, narrative.GetText())
				}
			})
			if len(justification.SatisfiesData.GetCoveredBy()) > 0 {
				text += "Covered By:\n"
			}
			for _, coveredBy := range justification.SatisfiesData.GetCoveredBy() {
				componentKey := coveredBy.ComponentKey
				if componentKey == "" {
					componentKey = justification.ComponentKey
				}
				openControl.Components.GetAndApply(componentKey, func(component base.Component) {
					if component != nil {
						verification := component.GetVerifications().Get(coveredBy.VerificationKey)
						text += exportLink(
							fmt.Sprintf("%s - %s", component.GetName(), verification.Name),
							filepath.Join("..", "components", component.GetKey()+".md"),
						)
					}
				})
			}
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
