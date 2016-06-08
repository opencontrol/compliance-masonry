package docx_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry/docx"
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/doc-template"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Docx", func() {
	DescribeTable("FormatControl", func(standard string, control string, expectedData string, sectionKeys ...string) {
		openControl := docx.OpenControlDocx{
			OpenControl: models.LoadData(filepath.Join("..", "fixtures", "opencontrol_fixtures"), ""),
		}
		actualData := openControl.FormatControl(standard, control, sectionKeys...)
		assert.Equal(GinkgoT(), expectedData, actualData)
	},
		// Get All Control Data
		Entry("openControl.FormatControl(NIST-800-53@CM-2)", "NIST-800-53", "CM-2", "Amazon Elastic Compute Cloud\na:\nJustification in narrative form A for CM-2\nb:\nJustification in narrative form B for CM-2\n"),
		Entry("openControl.FormatControl(PCI-DSS-MAY-2015@2.1)", "PCI-DSS-MAY-2015", "2.1", "Amazon Elastic Compute Cloud\nJustification in narrative form for 2.1\n"),
		Entry("openControl.FormatControl(PCI-DSS-MAY-2015@1.1.1)", "PCI-DSS-MAY-2015", "1.1.1", "Amazon Elastic Compute Cloud\n"+constants.WarningNoInformationAvailable+"\n"),
		Entry("openControl.FormatControl(BogusStandard@NothingControl)", "BogusStandard", "NothingControl", fmt.Sprintf(constants.WarningUnknownStandardAndControlf, "BogusStandard", "NothingControl")),
		// Get Specific Control Data
		Entry(`openControl.FormatControl(NIST-800-53,CM-2,a) - Regular case that should return one section from a loaded component.`,
			"NIST-800-53", "CM-2", "Amazon Elastic Compute Cloud\nJustification in narrative form A for CM-2\n", "a"),
		Entry(`openControl.FormatControl(PCI-DSS-MAY-2015,2.1,X)
			- Regular case that should return not section nor header from a loaded component
			 because the section does not exist. Instead provide a warning that nothing exists`,
			"PCI-DSS-MAY-2015", "2.1", "Amazon Elastic Compute Cloud\n"+constants.WarningNoInformationAvailable+"\n", "X"),
		Entry("openControl.FormatControl(BogusStandard,NothingControl,'')", "BogusStandard", "NothingControl", fmt.Sprintf(constants.WarningUnknownStandardAndControlf, "BogusStandard", "NothingControl")),
	)

	Describe("BuildDoc Tests", func() {
		It("loads the template and build the final docx", func() {
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
