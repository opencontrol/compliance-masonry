package docx_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/geramirez/doc-template"
	"github.com/opencontrol/compliance-masonry-go/docx"
	"github.com/opencontrol/compliance-masonry-go/models"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Docx", func() {
	table.DescribeTable("SplitControl", func(standardControl string, expectedStandard string, expectedControl string) {
		actualStandard, actualControl := docx.SplitControl(standardControl)
		assert.Equal(GinkgoT(), actualStandard, expectedStandard)
		assert.Equal(GinkgoT(), actualControl, expectedControl)
	},
		table.Entry("SplitControl(NIST-800-53@AC-2)", "NIST-800-53@AC-2", "NIST-800-53", "AC-2"),
		table.Entry("SplitControl(NIST-800-53@AC)", "NIST-800-53@AC", "NIST-800-53", "AC"),
		table.Entry("SplitControl(PCI@1.1.1)", "PCI@1.1.1", "PCI", "1.1.1"),
		table.Entry("SplitControl(PCI)", "PCI", "PCI", ""),
	)

	table.DescribeTable("FormatControl", func(standardControl string, expectedData string) {
		openControl := docx.OpenControlDocx{
			models.LoadData("../fixtures/opencontrol_fixtures/", ""),
		}
		actualData := openControl.FormatControl(standardControl)
		assert.Equal(GinkgoT(), expectedData, actualData)
	},
		table.Entry("openControl.FormatControl(NIST-800-53@CM-2)", "NIST-800-53@CM-2", "Amazon Elastic Compute Cloud  \nJustification in narrative form  \nCovered By:  \n- EC2 Verification 1 http://VerificationURL.com  \n"),
		table.Entry("openControl.FormatControl(PCI-DSS-MAY-2015@2.1)", "PCI-DSS-MAY-2015@2.1", "Amazon Elastic Compute Cloud  \nJustification in narrative form  \n"),
		table.Entry("openControl.FormatControl(BogusControl@Nothing)", "BogusControl@Nothing", ""),
	)

	Describe("BuildDoc Tests", func() {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		exportPath := filepath.Join(tempDir, "test_output.docx")
		config := docx.Config{
			OpencontrolDir: "../fixtures/opencontrol_fixtures/",
			TemplatePath:   "../fixtures/template_fixtures/test.docx",
			ExportPath:     exportPath,
		}
		config.BuildDocx()
		expectedDoc, _ := docTemp.GetTemplate("../fixtures/exports_fixtures/output.docx")
		actualDoc, _ := docTemp.GetTemplate(exportPath)
		assert.Equal(GinkgoT(), expectedDoc.Document.GetContent(), actualDoc.Document.GetContent())
	})
})
