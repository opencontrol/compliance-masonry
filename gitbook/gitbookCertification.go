package gitbook

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry/models"
)

func (openControl *OpenControlGitBook) exportControl(control *ControlGitbook) (string, string) {
	key := replaceParentheses(fmt.Sprintf("%s-%s", control.standardKey, control.controlKey))
	text := fmt.Sprintf("#%s\n##%s\n", key, control.Name)
	openControl.Justifications.GetAndApply(control.standardKey, control.controlKey, func(selectJustifications models.Verifications) {
		for _, justification := range selectJustifications {
			openControl.Components.GetAndApply(justification.ComponentKey, func(component *models.Component) {
				text = fmt.Sprintf("%s\n#### %s\n", text, component.Name)
				text = fmt.Sprintf("%s%s\n", text, justification.SatisfiesData.Narrative)
			})
			if len(justification.SatisfiesData.CoveredBy) > 0 {
				text += "Covered By:\n"
			}
			for _, coveredBy := range justification.SatisfiesData.CoveredBy {
				componentKey := coveredBy.ComponentKey
				if componentKey == "" {
					componentKey = justification.ComponentKey
				}
				openControl.Components.GetAndApply(componentKey, func(component *models.Component) {
					if component != nil {
						verification := component.Verifications.Get(coveredBy.VerificationKey)
						text += exportLink(
							fmt.Sprintf("%s - %s", component.Name, verification.Name),
							filepath.Join("..", "components", component.Key+".md"),
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
