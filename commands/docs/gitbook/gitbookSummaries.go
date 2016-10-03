package gitbook

import (
	"fmt"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// createSubHeading will create a subheading with the passed in string.
func createSubHeading(text string) string {
	return fmt.Sprintf("## %s\n", text)
}

// fileName contains the base file name without the extension.
type fileName struct {
	name string
}

// withExt is a wrapper to wrap adding a extension and returning the final string.
func (f fileName) withExt(ext string) string {
	return fmt.Sprintf("%s%s", f.name, ext)
}

// createFileName creates a file name from multiple strings. Between each string, there will be a hyphen between them.
func createFileName(fileNameParts ...string) fileName {
	name := strings.Join(fileNameParts, "-")
	return fileName{name: name}
}

// BuildComponentsSummaries creates summaries the components for the general summary
func (openControl *OpenControlGitBook) buildComponentsSummaries() string {
	summary := "\n## Components\n"
	for _, component := range openControl.GetAllComponents() {
		summary += exportLink(component.GetName(),
			filepath.Join("components", createFileName(component.GetKey()).withExt(".md")))
	}
	return summary
}

// buildStandardsSummaries creates summaries the standards for the general summary (via SUMMARY.md) which is the
// returned string.
// In addition, it builds the summary page for the standard-control combination (e.g. NIST-800-53-CM.html for standard
// NIST-800-53 and control family CM). The collection of these summary pages are in the returned map.
func (openControl *OpenControlGitBook) buildStandardsSummaries() (string, *map[string]string) {
	var familyFileName fileName
	familySummaryMap := make(map[string]string)

	// create the Standards sub-heading for the SUMMARY.md
	summary := createSubHeading("Standards")

	// Go through all the standards for the certification.
	standardKeys := openControl.GetCertification().GetSortedStandards()
	for _, standardKey := range standardKeys {
		// Find all the information for a particular standard.
		standard, found := openControl.GetStandard(standardKey)
		if !found {
			continue
		}
		// Go through all the controls for the certification.
		controlKeys := openControl.GetCertification().GetControlKeysFor(standardKey)
		for _, controlKey := range controlKeys {
			var controlSummary string
			controlSummary, familyFileName, familySummaryMap =
				openControl.buildStandardsSummary(standardKey, controlKey, standard, familyFileName,
					familySummaryMap)
			// append summary
			summary += controlSummary
		}
	}
	return summary, &familySummaryMap
}

func (*OpenControlGitBook) buildStandardsSummary(standardKey, controlKey string, standard common.Standard,
	oldFamilyFileName fileName, familySummaryMap map[string]string) (string, fileName, map[string]string) {
	summary := ""
	// format the filename
	controlLink := replaceParentheses(createFileName(standardKey, controlKey).withExt(".md"))

	// get the control.
	control := standard.GetControl(controlKey)

	// get the control family and name
	controlFamily := control.GetFamily()
	controlName := control.GetName()

	// concatenate the standard and control family to create a filename.
	newFamilyFileName := createFileName(standardKey, controlFamily)

	// create control family headings if we are finally on a new heading file.
	if oldFamilyFileName.name != newFamilyFileName.name {
		familySummaryMap[newFamilyFileName.name] = createSubHeading(newFamilyFileName.name)
		summary += exportLink(controlFamily,
			filepath.Join("standards", newFamilyFileName.withExt(".md")))
		oldFamilyFileName = newFamilyFileName
	}

	// Add the control name as a link under the control family header.
	controlFullName := fmt.Sprintf("%s: %s", controlKey, controlName)
	summary += "\t" + exportLink(controlFullName, filepath.Join("standards", controlLink))

	// add the link to the summary page for that particular standard-control combination
	// which will be created later on.
	familySummaryMap[newFamilyFileName.name] += exportLink(controlFullName, controlLink)
	return summary, oldFamilyFileName, familySummaryMap
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
