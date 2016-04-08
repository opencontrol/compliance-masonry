package docx_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry-go/docx"
	"github.com/opencontrol/compliance-masonry-go/models"
	"github.com/opencontrol/doc-template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Docx", func() {
	DescribeTable("SplitControl", func(standardControl string, expectedStandard string, expectedControl string) {
		actualStandard, actualControl := docx.SplitControl(standardControl)
		assert.Equal(GinkgoT(), actualStandard, expectedStandard)
		assert.Equal(GinkgoT(), actualControl, expectedControl)
	},
		Entry("SplitControl(NIST-800-53@AC-2)", "NIST-800-53@AC-2", "NIST-800-53", "AC-2"),
		Entry("SplitControl(NIST-800-53@AC)", "NIST-800-53@AC", "NIST-800-53", "AC"),
		Entry("SplitControl(PCI@1.1.1)", "PCI@1.1.1", "PCI", "1.1.1"),
		Entry("SplitControl(PCI)", "PCI", "PCI", ""),
	)

	DescribeTable("FormatControl", func(standardControl string, expectedData string) {
		openControl := docx.OpenControlDocx{
			models.LoadData(filepath.Join("..", "fixtures", "opencontrol_fixtures", "")),
		}
		actualData := openControl.FormatControl(standardControl)
		assert.Equal(GinkgoT(), expectedData, actualData)
	},
		Entry("openControl.FormatControl(NIST-800-53@CM-2)", "NIST-800-53@CM-2", "Amazon Elastic Compute Cloud  \nJustification in narrative form  \nCovered By:  \n- EC2 Verification 1 http://VerificationURL.com  \n"),
		Entry("openControl.FormatControl(PCI-DSS-MAY-2015@2.1)", "PCI-DSS-MAY-2015@2.1", "Amazon Elastic Compute Cloud  \nJustification in narrative form  \n"),
		Entry("openControl.FormatControl(BogusControl@Nothing)", "BogusControl@Nothing", ""),
	)

	Describe("BuildDoc Tests", func() {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		exportPath := filepath.Join(tempDir, "test_output.docx")
		config := docx.Config{
			OpencontrolDir: filepath.Join("..", "fixtures", "opencontrol_fixtures"),
			TemplatePath:   filepath.Join("..", "fixtures", "template_fixtures", "test.docx"),
			ExportPath:     exportPath,
		}
		config.BuildDocx()
		expectedDoc, _ := docTemp.GetTemplate(filepath.Join("..", "fixtures", "exports_fixtures", "output.docx"))
		actualDoc, _ := docTemp.GetTemplate(exportPath)
		assert.Equal(GinkgoT(), expectedDoc.Document.GetContent(), actualDoc.Document.GetContent())
	})
})
