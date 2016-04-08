package gitbook

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"testing"

	"github.com/opencontrol/compliance-masonry-go/models"
	"github.com/opencontrol/compliance-masonry-go/tools/fs"
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
	// Check that the component summary is correctly exported
	{
		filepath.Join("..", "fixtures", "opencontrol_fixtures"),
		filepath.Join("..", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"),
		"",
		filepath.Join("..", "fixtures", "exports_fixtures", "gitbook_exports", "components_readme.md"),
	},
}

func TestBuildComponentsSummaries(t *testing.T) {
	for _, example := range buildComponentsSummariesTests {
		openControl := OpenControlGitBook{
			models.LoadData(example.opencontrolDir, example.certificationPath),
			"",
			example.exportPath,
			fs.OSUtil{},
		}
		actualSummary := openControl.buildComponentsSummaries()
		data, err := ioutil.ReadFile(example.expectedSummary)
		if err != nil {
			log.Fatal(err)
		}
		expectedSummary := string(data)
		// Check that the actual and expected summaries are similar
		if strings.Replace(actualSummary, "\\", "/", -1) != expectedSummary {
			t.Errorf("Expected: `%s`, Actual: `%s`", expectedSummary, actualSummary)
		}
	}
}

var buildStandardsSummariesTests = []buildStandardsSummariesTest{
	// Check that a standards summary is correctly exported
	{
		filepath.Join("..", "fixtures", "opencontrol_fixtures"),
		filepath.Join("..", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"),
		"",
		filepath.Join("..", "fixtures", "exports_fixtures", "gitbook_exports", "standards_readme.md"),
		filepath.Join("..", "fixtures", "exports_fixtures", "gitbook_exports", "standards"),
	},
}

func TestBuildStandardsSummaries(t *testing.T) {
	for _, example := range buildStandardsSummariesTests {
		openControl := OpenControlGitBook{
			models.LoadData(example.opencontrolDir, example.certificationPath),
			"",
			example.exportPath,
			fs.OSUtil{},
		}
		actualSummary, familySummaryMap := openControl.buildStandardsSummaries()
		// Check the summary
		data, err := ioutil.ReadFile(example.expectedSummary)
		if err != nil {
			log.Fatal(err)
		}
		expectedSummary := string(data)
		// Check that the actual and expected summaries are similar
		if strings.Replace(actualSummary, "\\", "/", -1) != expectedSummary {
			t.Errorf("Expected: `%s`, Actual: `%s`", expectedSummary, actualSummary)
		}
		// Check individual pages
		for family, familySummary := range *(familySummaryMap) {
			data, err := ioutil.ReadFile(filepath.Join(example.expectedStandardsSummaries, family+".md"))
			if err != nil {
				log.Fatal(err)
			}
			expectedFamilySummary := string(data)
			// Check that the actual and expected summaries are similar
			if strings.Replace(familySummary, "\\", "/", -1) != expectedFamilySummary {
				t.Errorf("Expected: `%s`, Actual: `%s`", expectedFamilySummary, familySummary)
			}
		}

	}
}
