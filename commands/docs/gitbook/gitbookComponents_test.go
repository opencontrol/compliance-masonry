package gitbook

import (
	"testing"

	v2 "github.com/opencontrol/compliance-masonry/models/components/versions/2_0_0"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/blang/semver"
)

type componentExportTest struct {
	component    v2.Component
	expectedPath string
	expectedText string
}

var componentExportTests = []componentExportTest{
	// Check that a component is correctly exported
	{
		v2.Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    common.GeneralReferences{{}},
			Verifications: common.VerificationReferences{{}},
			Satisfies:     nil,
			SchemaVersion: semver.MustParse("2.0.0"),
		},
		"EC2.md",
		"# Amazon Elastic Compute Cloud\n## References\n* []()\n## Verifications\n* []()\n",
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
