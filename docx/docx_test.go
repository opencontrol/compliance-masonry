package docx

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/doc-template"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
	"fmt"
	v3 "github.com/opencontrol/compliance-masonry/models/components/versions/3_0_0"
)

var _ = Describe("Docx", func() {
	DescribeTable("FormatControl", func(standard string, control string, expectedData string, sectionKeys ...string) {
		openControl := OpenControlDocx{
			OpenControl: models.LoadData(filepath.Join("..", "fixtures", "opencontrol_fixtures"), ""),
		}
		actualData := openControl.FormatControl(standard, control, sectionKeys...)
		assert.Equal(GinkgoT(), expectedData, actualData)
	},
		// Get All Control Data
		Entry("openControl.FormatControl(NIST-800-53@CM-2)", "NIST-800-53", "CM-2", "Amazon Elastic Compute Cloud\na:\nJustification in narrative form A for CM-2\nb:\nJustification in narrative form B for CM-2\n"),
		Entry("openControl.FormatControl(PCI-DSS-MAY-2015@2.1)", "PCI-DSS-MAY-2015", "2.1", "Amazon Elastic Compute Cloud\nJustification in narrative form for 2.1\n"),
		Entry("openControl.FormatControl(PCI-DSS-MAY-2015@1.1.1) - Missing Control", "PCI-DSS-MAY-2015", "1.1.1", "No information found for the combination of standard PCI-DSS-MAY-2015 and control 1.1.1"),
		Entry("openControl.FormatControl(BogusStandard@NothingControl)", "BogusStandard", "NothingControl", "No information found for the combination of standard BogusStandard and control NothingControl"),
		// Get Specific Control Data
		Entry(`openControl.FormatControl(NIST-800-53,CM-2,a) - Regular case that should return one section from a loaded component.`,
			"NIST-800-53", "CM-2", "Amazon Elastic Compute Cloud\nJustification in narrative form A for CM-2\n", "a"),
		Entry(`openControl.FormatControl(PCI-DSS-MAY-2015,2.1,X)
			- Regular case that should return no section nor header from a loaded component
			 because the section does not exist. Instead provide a warning that nothing exists`,
			"PCI-DSS-MAY-2015", "2.1", "Amazon Elastic Compute Cloud\nNo information available for component\n", "X"),
		Entry("openControl.FormatControl(BogusStandard,NothingControl,'')", "BogusStandard", "NothingControl", fmt.Sprintf("No information found for the combination of standard %s and control %s", "BogusStandard", "NothingControl")),
	)

	DescribeTable("FormatParameter", func(standard string, control string, expectedData string, sectionKeys ...string) {
		openControl := OpenControlDocx{
			OpenControl: models.LoadData(filepath.Join("..", "fixtures", "opencontrol_fixtures"), ""),
		}
		actualData := openControl.FormatParameter(standard, control, sectionKeys...)
		assert.Equal(GinkgoT(), expectedData, actualData)
	},
		Entry("openControl.FormatParameter(NIST-800-53@CM-2) - Not specifying a parameter should say no available info", "NIST-800-53", "CM-2", "Amazon Elastic Compute Cloud\nNo information available for component\n"),
		Entry("openControl.FormatParameter(NIST-800-53@CM-2,a) - Not specifying a parameter should say no available info", "NIST-800-53", "CM-2", "Amazon Elastic Compute Cloud\nNo information available for component\n", "a"),

		Entry(`openControl.FormatParameter(PCI-DSS-MAY-2015,1.1,a) - Regular case`,
			"PCI-DSS-MAY-2015", "1.1", "Amazon Elastic Compute Cloud\nParameter A for 1.1\n", "a"),
		Entry("openControl.FormatParameter(BogusStandard,NothingControl,'')", "BogusStandard", "NothingControl", "No information found for the combination of standard BogusStandard and control NothingControl"),
	)

	DescribeTable("FormatResponsibleRole", func(standard string, control string, expectedData string, sectionKeys ...string) {
		openControl := OpenControlDocx{
			OpenControl: models.LoadData(filepath.Join("..", "fixtures", "opencontrol_fixtures"), ""),
		}
		actualData := openControl.FormatResponsibleRoles(standard, control)
		assert.Equal(GinkgoT(), expectedData, actualData)
	},
		Entry("openControl.FormatResponsibleRole(NIST-800-53@CM-2) - Regular case", "NIST-800-53", "CM-2", "Amazon Elastic Compute Cloud: AWS Staff\n"),
		Entry("openControl.FormatResponsibleRole(PCI-DSS-MAY-2015@1.1.1) - Missing Control", "PCI-DSS-MAY-2015", "1.1.1", "No information found for the combination of standard PCI-DSS-MAY-2015 and control 1.1.1"),
		Entry("openControl.FormatResponsibleRole(BogusStandard@NothingControl)", "BogusStandard", "NothingControl", "No information found for the combination of standard BogusStandard and control NothingControl"),
	)
	// For tests that there are no fixtures for.
	Describe("Misc. FormatResponsibleRole Tests", func() {
		It("should return just the component name when trying to get the responsible role and the component does not have one defined.", func() {
			c := v3.Component{Name:"Component Name"}
			text, found := getResponsibleRoleInfo("", &c)
			assert.Equal(GinkgoT(), "Component Name: ", text)
			assert.Equal(GinkgoT(), false, found)

		})
	})

	Describe("BuildDoc Tests", func() {
		It("loads the template and build the final docx", func() {
			tempDir, _ := ioutil.TempDir("", "example")
			defer os.RemoveAll(tempDir)
			exportPath := filepath.Join(tempDir, "test_output.docx")
			config := Config{
				OpencontrolDir: filepath.Join("..", "fixtures", "opencontrol_fixtures"),
				TemplatePath:   filepath.Join("..", "fixtures", "template_fixtures", "test.docx"),
				ExportPath:     exportPath,
			}
			err := config.BuildDocx()
			assert.Nil(GinkgoT(), err)
			expectedDoc, _ := docTemp.GetTemplate(filepath.Join("..", "fixtures", "exports_fixtures", "output.docx"))
			actualDoc, err := docTemp.GetTemplate(exportPath)
			assert.Nil(GinkgoT(), err)
			assert.Equal(GinkgoT(), expectedDoc.Document.GetContent(), actualDoc.Document.GetContent())
		})
	})
})
