package renderers

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
)

func (component *ComponentGitbook) exportComponent() (string, string) {
	var text string
	componentPath := component.Key + ".md"
	text = fmt.Sprintf("# %s  \n", component.Name)
	// Sort Components and Verifications
	if component.References != nil {
		if component.References.Len() > 0 {
			sort.Sort(component.References)
			text += "## References  \n"
			for _, reference := range *(component.References) {
				text += exportLink(reference.Name, reference.Path)
			}
		}
	}
	if component.Verifications != nil {
		if component.Verifications.Len() > 0 {
			sort.Sort(component.Verifications)
			text += "## Verifications  \n"
			for _, reference := range *(component.Verifications) {
				text += exportLink(reference.Name, reference.Path)
			}
		}
	}
	return filepath.Join(component.exportPath, componentPath), text
}

func (openControl *OpenControlGitBook) exportComponents() {
	componentsExportPath := createDirectory(filepath.Join(openControl.exportPath, "components"))
	for _, component := range openControl.Components.GetAll() {
		componentsGitBook := ComponentGitbook{component, componentsExportPath}
		componentPath, componentText := componentsGitBook.exportComponent()
		ioutil.WriteFile(componentPath, []byte(componentText), 0700)
	}
}
