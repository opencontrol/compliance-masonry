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

	// create the Standards sub-heading for the SUMMARY.md
	summary := "## Standards\n"

	// Go through all the standards for the certification.
	standardKeys := openControl.Certification.GetSortedStandards()
	for _, standardKey := range standardKeys {
		// Find all the information for a particular standard.
		standard := openControl.Standards.Get(standardKey)
		// Go through all the controls for the
		controlKeys := standard.GetSortedControls()
		for _, controlKey := range controlKeys {
			componentLink := replaceParentheses(standardKey + "-" + controlKey + ".md")
			control := standard.GetControl(controlKey)
			controlFamily := control.GetFamily()
			controlName := control.GetName()
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
		}
	}
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
