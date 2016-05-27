package gitbook

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

func (openControl *OpenControlGitBook) exportControl(control *ControlGitbook) (string, string) {
	key := replaceParentheses(fmt.Sprintf("%s-%s", control.standardKey, control.controlKey))
	text := fmt.Sprintf("#%s\n##%s\n", key, control.Name)
	openControl.Justifications.GetAndApply(control.standardKey, control.controlKey, func(selectJustifications models.Verifications) {
		// In the case that no information was found period for the standard and control
		if len(selectJustifications) == 0 {
			errorText := fmt.Sprintf(constants.WarningUnknownStandardAndControlf, control.standardKey, control.controlKey)
			text = fmt.Sprintf("%s%s\n", text, errorText)
			return
		}
		for _, justification := range selectJustifications {
			openControl.Components.GetAndApply(justification.ComponentKey, func(component *models.Component) {
				text = fmt.Sprintf("%s\n#### %s\n", text, component.Name)
				if len(justification.SatisfiesData.Narrative) == 0 {
					text = fmt.Sprintf("%s%s\n", text, constants.WarningNoInformationAvailable)
					return
				}
				for _, narrative := range justification.SatisfiesData.Narrative {
					if narrative.Key != "" {
						text = fmt.Sprintf("%s\n##### %s\n", text, narrative.Key)
					}
					text = fmt.Sprintf("%s%s\n", text, narrative.Text)
				}
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
