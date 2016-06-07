package gitbook

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
)

func (component *ComponentGitbook) exportComponent() (string, string) {
	var text string
	componentPath := component.GetKey() + ".md"
	text = fmt.Sprintf("# %s\n", component.GetName())
	// Sort Components and Verifications
	if component.GetReferences() != nil {
		if component.GetReferences().Len() > 0 {
			sort.Sort(component.GetReferences())
			text += "## References\n"
			for _, reference := range *(component.GetReferences()) {
				text += exportLink(reference.Name, reference.Path)
			}
		}
	}
	if component.GetVerifications() != nil {
		if component.GetVerifications().Len() > 0 {
			sort.Sort(component.GetVerifications())
			text += "## Verifications\n"
			for _, reference := range *(component.GetVerifications()) {
				text += exportLink(reference.Name, reference.Path)
			}
		}
	}
	return filepath.Join(component.exportPath, componentPath), text
}

func (openControl *OpenControlGitBook) exportComponents() {
	componentsExportPath := filepath.Join(openControl.exportPath, "components")
	openControl.FSUtil.Mkdirs(componentsExportPath)
	for _, component := range openControl.Components.GetAll() {
		componentsGitBook := ComponentGitbook{component, componentsExportPath}
		componentPath, componentText := componentsGitBook.exportComponent()
		ioutil.WriteFile(componentPath, []byte(componentText), 0700)
	}
}
