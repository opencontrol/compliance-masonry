package renderers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/geramirez/masonry-go/models"
)

type OpenControlGitBook struct {
	*models.OpenControl
	exportPath string
}

func exportLink(text string, location string) string {
	return fmt.Sprintf("* [%s](%s)  \n", text, location)
}

func (openControl *OpenControlGitBook) exportStandards() string {
	var oldFamily, newFamily string
	readme := "## Standards  \n"
	for standardKey, certStandard := range openControl.Certification.Standards {
		standard := openControl.Standards[standardKey]
		for controlKey := range certStandard.Controls {
			newFamily = standard.Controls[controlKey].Family
			if newFamily != oldFamily {
				standardLink := filepath.Join("standards", standardKey+"-"+newFamily+".md")
				readme += exportLink(standardKey, standardLink)
			}
			controlLink := filepath.Join("standards", standardKey+"-"+controlKey+".md")
			readme += "\t" + exportLink(controlKey, controlLink)
			oldFamily = newFamily
		}
	}
	return readme
}

func (openControl *OpenControlGitBook) exportSystemsReadMe() string {
	var systemReadme string
	readme := "## Systems  \n"
	systemsExportPath := filepath.Join(openControl.exportPath, "systems")
	if _, err := os.Stat(systemsExportPath); os.IsNotExist(err) {
		os.MkdirAll(systemsExportPath, 0700)
	}
	for _, system := range openControl.Systems {
		systemLink := filepath.Join("systems", system.Key+".md")
		systemReadme = fmt.Sprintf("# %s  \n", system.Name)
		readme += exportLink(system.Name, systemLink)
		for _, component := range system.Components {
			componentPath := system.Key + "-" + component.Key + ".md"
			systemReadme += exportLink(component.Name, componentPath)
			readme += "\t" + exportLink(component.Name, filepath.Join("systems", componentPath))
		}
		ioutil.WriteFile(filepath.Join(systemsExportPath, system.Key+".md"), []byte(systemReadme), 0700)
	}
	return readme

}

func (openControl *OpenControlGitBook) BuildReadMe() {
	var readme string
	readme += openControl.exportStandards()
	readme += openControl.exportSystemsReadMe()
	ioutil.WriteFile(filepath.Join(openControl.exportPath, "SUMMARY.md"), []byte(readme), 0700)
	ioutil.WriteFile(filepath.Join(openControl.exportPath, "README.md"), []byte(readme), 0700)
}

func BuildGitbook(opencontrolDir string, certificationPath string, exportPath string) {
	openControl := OpenControlGitBook{
		models.LoadData(opencontrolDir, certificationPath),
		exportPath,
	}
	if _, err := os.Stat(exportPath); os.IsNotExist(err) {
		os.MkdirAll(exportPath, 0700)
	}

	openControl.BuildReadMe()
}
