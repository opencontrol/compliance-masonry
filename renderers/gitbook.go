package renderers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry-go/models"
)

type OpenControlGitBook struct {
	*models.OpenControl
	exportPath string
}

type ComponentGitbook struct {
	*models.Component
	exportPath string
}

func exportLink(text string, location string) string {
	return fmt.Sprintf("* [%s](%s)  \n", text, location)
}

func createDirectory(directory string) string {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.MkdirAll(directory, 0700)
	}
	return directory
}

func (openControl *OpenControlGitBook) exportStandards() string {
	var oldFamily, newFamily, familyReadMe, standardLink, readme string
	standardsExportPath := createDirectory(filepath.Join(openControl.exportPath, "standards"))
	readme = "## Standards  \n"
	for standardKey, certStandard := range openControl.Certification.Standards {
		standard := openControl.Standards.Get(standardKey)
		for controlKey := range certStandard.Controls {
			newFamily = standard.Controls[controlKey].Family
			if newFamily != oldFamily {
				if familyReadMe != "" && standardLink != "" {
					familyExportPath := filepath.Join(standardsExportPath, standardLink)
					ioutil.WriteFile(familyExportPath, []byte(familyReadMe), 0700)
				}
				familyReadMe = fmt.Sprintf("### %s  \n", newFamily)
				standardLink = filepath.Join(standardKey + "-" + newFamily + ".md")
				readme += exportLink(standardKey, standardLink)
			}
			controlLink := filepath.Join("standards", standardKey+"-"+controlKey+".md")
			if familyReadMe != "" && standardLink != "" {
				familyReadMe += exportLink(controlKey, standardLink)
			}
			readme += "\t" + exportLink(controlKey, controlLink)
			oldFamily = newFamily
		}
	}
	return readme
}

func (component *ComponentGitbook) exportComponent() string {
	var readme string
	componentPath := component.Key + ".md"
	readme += "\t" + exportLink(component.Name, filepath.Join("components", componentPath))
	return readme
}

func (openControl *OpenControlGitBook) exportComponents() string {
	readme := "## Components  \n"
	componentsExportPath := createDirectory(filepath.Join(openControl.exportPath, "components"))
	for _, component := range openControl.Components.GetAll() {
		componentsGitBook := ComponentGitbook{component, componentsExportPath}
		readme += componentsGitBook.exportComponent()
	}
	return readme
}

func (openControl *OpenControlGitBook) BuildReadMe() {
	var readme string
	readme += openControl.exportStandards()
	readme += openControl.exportComponents()
	ioutil.WriteFile(filepath.Join(openControl.exportPath, "SUMMARY.md"), []byte(readme), 0700)
	ioutil.WriteFile(filepath.Join(openControl.exportPath, "README.md"), []byte(readme), 0700)
}

func BuildGitbook(opencontrolDir string, certificationPath string, exportPath string) {
	openControl := OpenControlGitBook{
		models.LoadData(opencontrolDir, certificationPath),
		exportPath,
	}
	createDirectory(exportPath)

	openControl.BuildReadMe()
}
