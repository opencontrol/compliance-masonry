package renderers

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// BuildComponentsSummaries creates summaries the components for the general readme
func (openControl *OpenControlGitBook) BuildComponentsSummaries() string {
	readme := "  \n## Components  \n"
	for _, component := range openControl.Components.GetAll() {
		readme += "\t" + exportLink(component.Name, filepath.Join("components", component.Key+".md"))
	}
	return readme
}

// BuildStandardsSummaries creates summaries the standards for the general readme
func (openControl *OpenControlGitBook) BuildStandardsSummaries() (string, *map[string]string) {
	var oldFamily, newFamily string
	familyReadMeMap := make(map[string]string)
	readme := "## Standards  \n"
	openControl.Certification.GetSortedData(func(standardKey string, controlKey string) {
		componentLink := standardKey + "-" + controlKey + ".md"
		controlFamily := openControl.Standards.Get(standardKey).Controls[controlKey].Family
		newFamily = standardKey + "-" + controlFamily
		if oldFamily != newFamily {
			familyReadMeMap[newFamily] = fmt.Sprintf("## %s  \n", newFamily)
			readme += exportLink(controlKey, filepath.Join("standards", newFamily+".md"))
			oldFamily = newFamily
		}
		familyReadMeMap[newFamily] += exportLink(controlKey, componentLink)
		readme += "\t" + exportLink(controlKey, filepath.Join("standards", componentLink))
	})
	return readme, &familyReadMeMap
}

func (openControl *OpenControlGitBook) exportFamilyReadMap(familyReadMeMap *map[string]string) {
	for family, familyReadMe := range *(familyReadMeMap) {
		ioutil.WriteFile(filepath.Join(openControl.exportPath, "standards", family+".md"), []byte(familyReadMe), 0700)
	}
}

// BuildSummaries creates the general readme
func (openControl *OpenControlGitBook) BuildSummaries() {
	standardsReadMe, familyReadMeMap := openControl.BuildStandardsSummaries()
	componentsReadMe := openControl.BuildComponentsSummaries()
	openControl.exportFamilyReadMap(familyReadMeMap)
	readMe := standardsReadMe + componentsReadMe
	go ioutil.WriteFile(filepath.Join(openControl.exportPath, "SUMMARY.md"), []byte(readMe), 0700)
	go ioutil.WriteFile(filepath.Join(openControl.exportPath, "README.md"), []byte(readMe), 0700)
}
