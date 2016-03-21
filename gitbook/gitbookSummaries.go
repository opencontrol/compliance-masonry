package gitbook

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry-go/helpers"
)

// BuildComponentsSummaries creates summaries the components for the general summary
func (openControl *OpenControlGitBook) buildComponentsSummaries() string {
	summary := "  \n## Components  \n"
	for _, component := range openControl.Components.GetAll() {
		summary += exportLink(component.Name, filepath.Join("components", component.Key+".md"))
	}
	return summary
}

// BuildStandardsSummaries creates summaries the standards for the general summary
func (openControl *OpenControlGitBook) buildStandardsSummaries() (string, *map[string]string) {
	var oldFamily, newFamily string
	familySummaryMap := make(map[string]string)
	summary := "## Standards  \n"

	openControl.Certification.GetSortedData(func(standardKey string, controlKey string) {
		componentLink := replaceParentheses(standardKey + "-" + controlKey + ".md")
		controlFamily := openControl.Standards.Get(standardKey).Controls[controlKey].Family
		newFamily = standardKey + "-" + controlFamily
		if oldFamily != newFamily {
			familySummaryMap[newFamily] = fmt.Sprintf("## %s  \n", newFamily)
			summary += exportLink(controlKey, filepath.Join("standards", newFamily+".md"))
			oldFamily = newFamily
		}
		familySummaryMap[newFamily] += exportLink(controlKey, componentLink)
		summary += "\t" + exportLink(controlKey, filepath.Join("standards", componentLink))
	})
	return summary, &familySummaryMap
}

func (openControl *OpenControlGitBook) exportFamilyReadMap(familySummaryMap *map[string]string) {
	for family, familySummary := range *(familySummaryMap) {
		ioutil.WriteFile(filepath.Join(openControl.exportPath, "standards", family+".md"), []byte(familySummary), 0700)
	}
}

func (openControl *OpenControlGitBook) buildMarkdowns() {
	if openControl.markdownPath != "" {
		helpers.CopyDir(openControl.markdownPath, openControl.exportPath)
	}
}

// buildSummaries creates the general summary
func (openControl *OpenControlGitBook) buildSummaries() {
	openControl.buildMarkdowns()
	standardsSummary, familySummaryMap := openControl.buildStandardsSummaries()
	componentsSummary := openControl.buildComponentsSummaries()
	openControl.exportFamilyReadMap(familySummaryMap)
	summary := standardsSummary + componentsSummary
	go helpers.AppendOrCreate(filepath.Join(openControl.exportPath, "SUMMARY.md"), summary)
	go helpers.AppendOrCreate(filepath.Join(openControl.exportPath, "README.md"), summary)
}
