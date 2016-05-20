package docx_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry/docx"
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/doc-template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Docx", func() {
	DescribeTable("FormatControl", func(standard string, control string, expectedData string) {
		openControl := docx.OpenControlDocx{
			OpenControl: models.LoadData(filepath.Join("..", "fixtures", "opencontrol_fixtures"), ""),
		}
		actualData := openControl.FormatControl(standard, control)
		assert.Equal(GinkgoT(), expectedData, actualData)
	},
		Entry("openControl.FormatControl(NIST-800-53@CM-2)", "NIST-800-53", "CM-2", "Amazon Elastic Compute Cloud\na:\nJustification in narrative form A for CM-2\nb:\nJustification in narrative form B for CM-2\n"),
		Entry("openControl.FormatControl(PCI-DSS-MAY-2015@2.1)", "PCI-DSS-MAY-2015", "2.1", "Amazon Elastic Compute Cloud\nJustification in narrative form for 2.1\n"),
		Entry("openControl.FormatControl(BogusControl@Nothing)", "BogusControl", "Nothing", ""),
	)

	DescribeTable("FormatControlSection", func(standard string, control string, section string, expectedData string) {
		openControl := docx.OpenControlDocx{
			OpenControl: models.LoadData(filepath.Join("..", "fixtures", "opencontrol_fixtures"), ""),
		}
		actualData := openControl.FormatControlSection(standard, control, section)
		assert.Equal(GinkgoT(), expectedData, actualData)
	},
		Entry(`openControl.FormatControlSection(NIST-800-53,CM-2,a) - Regular case that should return one section from a loaded component.`,
			"NIST-800-53", "CM-2", "a", "Amazon Elastic Compute Cloud\nJustification in narrative form A for CM-2\n"),
		Entry(`openControl.FormatControlSection(PCI-DSS-MAY-2015,2.1,X)
			- Regular case that should return not section nor header from a loaded component because the section does not exist`,
			"PCI-DSS-MAY-2015", "2.1", "X", ""),
		Entry("openControl.FormatControlSection(BogusControl,Nothing,'')", "BogusControl", "Nothing", "", ""),
	)

	Describe("BuildDoc Tests", func() {
		It("load the template and build the final docx", func() {
			tempDir, _ := ioutil.TempDir("", "example")
			defer os.RemoveAll(tempDir)
			exportPath := filepath.Join(tempDir, "test_output.docx")
			config := docx.Config{
				OpencontrolDir: filepath.Join("..", "fixtures", "opencontrol_fixtures"),
				TemplatePath:   filepath.Join("..", "fixtures", "template_fixtures", "test.docx"),
				ExportPath:     exportPath,
			}
			err := config.BuildDocx()
			assert.Nil(GinkgoT(), err)
			expectedDoc, _ := docTemp.GetTemplate(filepath.Join("..", "fixtures", "exports_fixtures", "output.docx"))
			actualDoc, err := docTemp.GetTemplate(exportPath)
			assert.NotNil(GinkgoT(), expectedDoc)
			assert.NotNil(GinkgoT(), actualDoc)
			assert.Nil(GinkgoT(), err)
			assert.Equal(GinkgoT(), expectedDoc.Document.GetContent(), actualDoc.Document.GetContent())
		})

	})

})
