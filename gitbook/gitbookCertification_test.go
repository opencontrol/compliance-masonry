package gitbook

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/tools/fs"
)

type exportControlTest struct {
	opencontrolDir    string
	certificationPath string
	standardKey       string
	controlKey        string
	expectedPath      string
	expectedText      string
}

var exportControlTests = []exportControlTest{
	// Check that a control is exported correctly
	{
		filepath.Join("..", "fixtures", "opencontrol_fixtures"),
		filepath.Join("..", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"),
		"NIST-800-53",
		"CM-2",
		"NIST-800-53-CM-2.md",
		"#NIST-800-53-CM-2  \n##Baseline Configuration  \n  \n#### Amazon Elastic Compute Cloud  \nJustification in narrative form  \nCovered By:  \n* [Amazon Elastic Compute Cloud - EC2 Verification 1](../components/EC2.md)  \n",
	},
}

func TestExportControl(t *testing.T) {
	for _, example := range exportControlTests {
		dir, err := ioutil.TempDir("", "example")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(dir)
		openControl := OpenControlGitBook{
			models.LoadData(example.opencontrolDir, example.certificationPath),
			"",
			dir,
			fs.OSUtil{},
		}
		control := openControl.Standards.Get(example.standardKey).Controls[example.controlKey]
		actualPath, actualText := openControl.exportControl(&ControlGitbook{&control, dir, example.standardKey, example.controlKey})
		expectedPath := filepath.Join(dir, example.expectedPath)
		// Verify the expected export path is the same as the actual export path
		if expectedPath != actualPath {
			t.Errorf("Expected %s, Actual: %s", example.expectedPath, actualPath)
		}
		// Verify the expected text is the same as the actual text
		if example.expectedText != strings.Replace(actualText, "\\", "/", -1) {
			t.Errorf("Expected %s, Actual: %s", example.expectedText, actualText)
		}

	}
}
