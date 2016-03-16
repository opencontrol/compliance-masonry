package renderers

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func (openControl *OpenControlGitBook) exportControl(control *ControlGitbook) (string, string) {
	key := fmt.Sprintf("%s-%s", control.standardKey, control.controlKey)
	text := fmt.Sprintf("#%s  \n##%s  \n", key, control.Name)
	justifications := openControl.Justifications.Get(control.standardKey, control.controlKey)
	for _, justification := range justifications {
		component := openControl.Components.Get(justification.Component)
		verification := component.Verifications.Get(justification.Verification)
		if verification.Name != "" {
			text += exportLink(
				fmt.Sprintf("%s - %s", component.Name, verification.Name),
				filepath.Join("..", "components", component.Key+".md"),
			)
		}
	}
	return filepath.Join(control.exportPath, key+".md"), text
}

func (openControl *OpenControlGitBook) exportStandards() {
	standardsExportPath := createDirectory(filepath.Join(openControl.exportPath, "standards"))
	openControl.Certification.GetSortedData(func(standardKey string, controlKey string) {
		control := openControl.Standards.Get(standardKey).Controls[controlKey]
		controlPath, controlText := openControl.exportControl(&ControlGitbook{&control, standardsExportPath, standardKey, controlKey})
		ioutil.WriteFile(controlPath, []byte(controlText), 0700)
	})
}
