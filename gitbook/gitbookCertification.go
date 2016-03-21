package gitbook

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry-go/helpers"
	"github.com/opencontrol/compliance-masonry-go/models"
)

func (openControl *OpenControlGitBook) exportControl(control *ControlGitbook) (string, string) {
	key := replaceParentheses(fmt.Sprintf("%s-%s", control.standardKey, control.controlKey))
	text := fmt.Sprintf("#%s  \n##%s  \n", key, control.Name)

	openControl.Justifications.GetAndApply(control.standardKey, control.controlKey, func(justifications models.Verifications) {
		for _, justification := range justifications {
			openControl.Components.GetAndApply(justification.Component, func(component *models.Component) {
				if component != nil {
					verification := component.Verifications.Get(justification.Verification)
					if verification.Name != "" {
						text += exportLink(
							fmt.Sprintf("%s - %s", component.Name, verification.Name),
							filepath.Join("..", "components", component.Key+".md"),
						)
					}
				}
			})
		}
	})
	return filepath.Join(control.exportPath, key+".md"), text
}

func (openControl *OpenControlGitBook) exportStandards() {
	standardsExportPath := helpers.CreateDirectory(filepath.Join(openControl.exportPath, "standards"))
	openControl.Certification.GetSortedData(func(standardKey string, controlKey string) {
		control := openControl.Standards.Get(standardKey).Controls[controlKey]
		controlPath, controlText := openControl.exportControl(&ControlGitbook{&control, standardsExportPath, standardKey, controlKey})
		ioutil.WriteFile(controlPath, []byte(controlText), 0700)
	})
}
