package renderers

import (
	"testing"

	"github.com/opencontrol/compliance-masonry-go/models"
)

type componentExportTest struct {
	component    models.Component
	expectedPath string
	expectedText string
}

var componentExportTests = []componentExportTest{
	// Check that a component is correctly exported
	{
		models.Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    &models.GeneralReferences{{}},
			Verifications: &models.VerificationReferences{{}},
			Satisfies:     &models.SatisfiesList{{}},
			SchemaVersion: 2.0,
		},
		"EC2.md",
		"# Amazon Elastic Compute Cloud  \n## References  \n* []()  \n## Verifications  \n* []()  \n",
	},
}

func TestExportComponent(t *testing.T) {
	for _, example := range componentExportTests {
		gitbookComponent := ComponentGitbook{&example.component, ""}
		actualPath, actualText := gitbookComponent.exportComponent()
		if example.expectedPath != actualPath {
			t.Errorf("Expected %s, Actual: %s", example.expectedPath, actualPath)
		}
		if example.expectedText != actualText {
			t.Errorf("Expected %s, Actual: %s", example.expectedText, actualText)
		}
	}
}
