package renderers

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/opencontrol/compliance-masonry-go/models"
)

type buildComponentsSummariesTest struct {
	opencontrolDir    string
	certificationPath string
	exportPath        string
	expectedSummary   string
}

type buildStandardsSummariesTest struct {
	opencontrolDir             string
	certificationPath          string
	exportPath                 string
	expectedSummary            string
	expectedStandardsSummaries string
}

var buildComponentsSummariesTests = []buildComponentsSummariesTest{
	{
		"../fixtures/opencontrol_fixtures/",
		"../fixtures/opencontrol_fixtures/certifications/LATO.yaml",
		"",
		"../fixtures/exports_fixtures/gitbook_exports/components_readme.md",
	},
}

func TestBuildComponentsSummaries(t *testing.T) {
	for _, example := range buildComponentsSummariesTests {
		openControl := OpenControlGitBook{
			models.LoadData(example.opencontrolDir, example.certificationPath),
			example.exportPath,
		}
		actualSummary := openControl.buildComponentsSummaries()
		data, err := ioutil.ReadFile(example.expectedSummary)
		if err != nil {
			log.Fatal(err)
		}
		expectedSummary := string(data)
		if actualSummary != expectedSummary {
			t.Errorf("Expected: `%s`, Actual: `%s`", expectedSummary, actualSummary)
		}
	}
}

var buildStandardsSummariesTests = []buildStandardsSummariesTest{
	{
		"../fixtures/opencontrol_fixtures/",
		"../fixtures/opencontrol_fixtures/certifications/LATO.yaml",
		"",
		"../fixtures/exports_fixtures/gitbook_exports/standards_readme.md",
		"../fixtures/exports_fixtures/gitbook_exports/standards/",
	},
}

func TestBuildStandardsSummaries(t *testing.T) {
	for _, example := range buildStandardsSummariesTests {
		openControl := OpenControlGitBook{
			models.LoadData(example.opencontrolDir, example.certificationPath),
			example.exportPath,
		}
		actualSummary, familySummaryMap := openControl.buildStandardsSummaries()
		// Check the summary
		data, err := ioutil.ReadFile(example.expectedSummary)
		if err != nil {
			log.Fatal(err)
		}
		expectedSummary := string(data)
		if actualSummary != expectedSummary {
			t.Errorf("Expected: `%s`, Actual: `%s`", expectedSummary, actualSummary)
		}
		// Check individual pages
		for family, familySummary := range *(familySummaryMap) {
			data, err := ioutil.ReadFile(filepath.Join(example.expectedStandardsSummaries, family+".md"))
			if err != nil {
				log.Fatal(err)
			}
			expectedFamilySummary := string(data)
			if familySummary != expectedFamilySummary {
				t.Errorf("Expected: `%s`, Actual: `%s`", expectedFamilySummary, familySummary)
			}
		}

	}
}
