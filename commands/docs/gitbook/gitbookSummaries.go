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

// buildStandardsSummaries creates summaries the standards for the general summary (via SUMMARY.md) which is the
// returned string.
// In addition, it builds the summary page for the standard-control combination (e.g. NIST-800-53-CM.html for standard
// NIST-800-53 and control family CM). The collection of these summary pages are in the returned map.
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
		// Go through all the controls for the certification.
		controlKeys := openControl.Certification.GetControlKeysFor(standardKey)
		for _, controlKey := range controlKeys {
			// format the filename
			controlLink := replaceParentheses(standardKey + "-" + controlKey + ".md")

			// get the control.
			control := standard.GetControl(controlKey)

			// get the control family and name
			controlFamily := control.GetFamily()
			controlName := control.GetName()

			// concatenate the standard and control family
			newFamily = standardKey + "-" + controlFamily

			// create control family headings if we are finally on a new heading.
			if oldFamily != newFamily {
				familySummaryMap[newFamily] = fmt.Sprintf("## %s\n", newFamily)
				summary += exportLink(controlFamily, filepath.Join("standards", newFamily+".md"))
				oldFamily = newFamily
			}

			// Add the control name as a link under the control family header.
			controlFullName := fmt.Sprintf("%s: %s", controlKey, controlName)
			summary += "\t" + exportLink(controlFullName, filepath.Join("standards", controlLink))

			// add the link to the summary page for that particular standard-control combination
			// which will be created later on.
			familySummaryMap[newFamily] += exportLink(controlFullName, controlLink)
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
