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

type SystemGitbook struct {
	*models.System
	exportPath string
}

type ComponentGitbook struct {
	*models.Component
	exportPath string
	systemKey  string
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

func (component *ComponentGitbook) exportComponent() (string, string) {
	var readme, systemReadme string
	componentPath := component.systemKey + "-" + component.Key + ".md"
	systemReadme += exportLink(component.Name, componentPath)
	readme += "\t" + exportLink(component.Name, filepath.Join("systems", componentPath))
	return readme, systemReadme
}

func (system *SystemGitbook) exportSystem() string {
	var readme, systemReadme string
	systemLink := filepath.Join("systems", system.Key+".md")
	systemReadme = fmt.Sprintf("# %s  \n", system.Name)
	readme += exportLink(system.Name, systemLink)
	for _, component := range system.Components {
		componentGitbook := ComponentGitbook{component, system.exportPath, system.Key}
		readmeUpdate, systemReadmeUpdate := componentGitbook.exportComponent()
		readme += readmeUpdate
		systemReadme += systemReadmeUpdate
	}
	ioutil.WriteFile(filepath.Join(system.exportPath, system.Key+".md"), []byte(systemReadme), 0700)
	return readme
}

func (openControl *OpenControlGitBook) exportSystems() string {
	readme := "## Systems  \n"
	systemsExportPath := filepath.Join(openControl.exportPath, "systems")
	if _, err := os.Stat(systemsExportPath); os.IsNotExist(err) {
		os.MkdirAll(systemsExportPath, 0700)
	}
	for _, system := range openControl.Systems {
		systemGitBook := SystemGitbook{system, systemsExportPath}
		readme += systemGitBook.exportSystem()
	}
	return readme
}

func (openControl *OpenControlGitBook) BuildReadMe() {
	var readme string
	readme += openControl.exportStandards()
	readme += openControl.exportSystems()
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
