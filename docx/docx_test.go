package docx

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/geramirez/doc-template"
	"github.com/opencontrol/compliance-masonry-go/models"
	"github.com/stretchr/testify/assert"
)

var splitControlTests = []struct {
	standardControl  string
	expectedControl  string
	expectedStandard string
}{
	{standardControl: "NIST-800-53@AC-2", expectedStandard: "NIST-800-53", expectedControl: "AC-2"},
	{standardControl: "NIST-800-53@AC", expectedStandard: "NIST-800-53", expectedControl: "AC"},
	{standardControl: "PCI@1.1.1", expectedStandard: "PCI", expectedControl: "1.1.1"},
	{standardControl: "PCI", expectedStandard: "PCI", expectedControl: ""},
}

func TestSplitControl(t *testing.T) {
	for _, example := range splitControlTests {
		actualStandard, actualControl := splitControl(example.standardControl)
		assert.Equal(t, example.expectedStandard, actualStandard)
		assert.Equal(t, example.expectedControl, actualControl)
	}
}

var formatControlTests = []struct {
	standardControl string
	expectedData    string
}{
	{standardControl: "NIST-800-53@CM-2", expectedData: "Justification in narrative form"},
	{standardControl: "PCI-DSS-MAY-2015@2.1", expectedData: "Justification in narrative form"},
}

func TestFormatControl(t *testing.T) {
	openControl := OpenControlDocx{
		models.LoadData("../fixtures/opencontrol_fixtures/", ""),
	}
	for _, example := range formatControlTests {
		actualData := openControl.formatControl(example.standardControl)
		assert.Equal(t, example.expectedData, actualData)
	}
}

func TestBuildDocx(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(tempDir)
	exportPath := filepath.Join(tempDir, "test_output.docx")
	config := Config{
		OpencontrolDir: "../fixtures/opencontrol_fixtures/",
		TemplatePath:   "../fixtures/opencontrol_fixtures/test.docx",
		ExportPath:     exportPath,
	}
	config.BuildDocx()
	expectedDoc, _ := docTemp.GetTemplate("../fixtures/exports_fixtures/output.docx")
	actualDoc, _ := docTemp.GetTemplate(exportPath)
	assert.Equal(t, expectedDoc.Document.GetContent(), actualDoc.Document.GetContent())

}
