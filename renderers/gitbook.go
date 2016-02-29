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

func (openControl *OpenControlGitBook) exportStandardsReadMe() string {
	var oldFamily, newFamily string
	readme := "## Standards  \n"
	for standardKey, certStandard := range openControl.Certification.Standards {
		standard := openControl.Standards[standardKey]
		for controlKey := range certStandard.Controls {
			newFamily = standard.Controls[controlKey].Family
			if newFamily != oldFamily {
				readme += exportLink(
					standardKey,
					filepath.Join("standards", standardKey+"-"+newFamily+".md"),
				)
			}
			readme += "\t" + exportLink(
				controlKey,
				filepath.Join("standards", standardKey+"-"+controlKey+".md"),
			)
			oldFamily = newFamily
		}
	}
	return readme

}

func (openControl *OpenControlGitBook) exportSystemsReadMe() string {
	readme := "## Systems  \n"
	for _, system := range openControl.Systems {
		readme += exportLink(
			system.Name,
			filepath.Join("standards", system.Key+".md"),
		)
		for _, component := range system.Components {
			readme += "\t" + exportLink(
				component.Name,
				filepath.Join("standards", system.Key+"-"+component.Key+".md"),
			)
		}
	}
	return readme

}

func (openControl *OpenControlGitBook) BuildReadMe() {
	var readme string
	readme += openControl.exportStandardsReadMe()
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
