package renderers

import (
	"testing"

	"github.com/opencontrol/compliance-masonry-go/models"
)

type buildComponentsSummariesTest struct {
	opencontrolDir    string
	certificationPath string
	exportPath        string
	expectedReadMe    string
}

var buildComponentsSummariesTests = []buildComponentsSummariesTest{
	{"../fixtures/opencontrol_fixtures/", "../fixtures/opencontrol_fixtures/certifications/LATO.yaml", "", "  \n## Components  \n\t* [Amazon Elastic Compute Cloud](components/EC2.md)  \n"},
}

func TestBuildComponentsSummaries(t *testing.T) {
	for _, example := range buildComponentsSummariesTests {
		openControl := OpenControlGitBook{
			models.LoadData(example.opencontrolDir, example.certificationPath),
			example.exportPath,
		}
		actualReadMe := openControl.BuildComponentsSummaries()
		if actualReadMe != example.expectedReadMe {
			t.Errorf("Expected: `%s`, Actual: `%s`", example.expectedReadMe, actualReadMe)
		}
	}
}
