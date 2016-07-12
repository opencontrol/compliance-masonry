package gitbook

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// BuildComponentsSummaries creates summaries the components for the general summary
func (openControl *OpenControlGitBook) buildComponentsSummaries() string {
	summary := "\n## Components\n"
	for _, component := range openControl.Components.GetAll() {
		summary += exportLink(component.GetName(), filepath.Join("components", component.GetKey()+".md"))
	}
	return summary
}

// BuildStandardsSummaries creates summaries the standards for the general summary
func (openControl *OpenControlGitBook) buildStandardsSummaries() (string, *map[string]string) {
	var oldFamily, newFamily string
	familySummaryMap := make(map[string]string)
	summary := "## Standards\n"

	openControl.Certification.GetSortedData(func(standardKey string, controlKey string) {
		componentLink := replaceParentheses(standardKey + "-" + controlKey + ".md")
		control := openControl.Standards.Get(standardKey).Controls[controlKey]
		controlFamily := control.Family
		controlName := control.Name
		newFamily = standardKey + "-" + controlFamily
		// create control family headings
		if oldFamily != newFamily {
			familySummaryMap[newFamily] = fmt.Sprintf("## %s\n", newFamily)
			summary += exportLink(controlFamily, filepath.Join("standards", newFamily+".md"))
			oldFamily = newFamily
		}
		controlFullName := fmt.Sprintf("%s: %s", controlKey, controlName)
		familySummaryMap[newFamily] += exportLink(controlFullName, componentLink)
		summary += "\t" + exportLink(controlFullName, filepath.Join("standards", componentLink))
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
		openControl.FSUtil.CopyAll(openControl.markdownPath, openControl.exportPath)
	}
}

// buildSummaries creates the general summary
func (openControl *OpenControlGitBook) buildSummaries() error {
	openControl.buildMarkdowns()
	standardsSummary, familySummaryMap := openControl.buildStandardsSummaries()
	componentsSummary := openControl.buildComponentsSummaries()
	openControl.exportFamilyReadMap(familySummaryMap)
	summary := standardsSummary + componentsSummary
	if err := openControl.FSUtil.AppendOrCreate(filepath.Join(openControl.exportPath, "SUMMARY.md"), summary); err != nil {
		return err
	}
	if err := openControl.FSUtil.AppendOrCreate(filepath.Join(openControl.exportPath, "README.md"), summary); err != nil {
		return err
	}
	return nil
}
