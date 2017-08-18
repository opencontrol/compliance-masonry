package export_test

import (
	. "github.com/opencontrol/compliance-masonry/commands/export"

	"errors"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
)

var _ = Describe("Export", func() {
	Describe("Verify Export functions", func() {
		var (
			workingDir           string
			jsonFormat           OutputFormat
			yamlFormat           OutputFormat
			standardKeySeparator string
			customKeySeparator   string
		)
		BeforeEach(func() {
			workingDir, _ = os.Getwd()
			jsonFormat, _ = ToOutputFormat("json")
			yamlFormat, _ = ToOutputFormat("yaml")
			standardKeySeparator = ":"
			customKeySeparator = ".."
		})
		Describe("bad inputs", func() {
			Context("When no arguments are specified", func() {
				It("should return an error", func() {
					config := Config{}
					err := Export(config)
					assert.Equal(GinkgoT(), []error{errors.New("Error: Missing Certification Argument")}, err)
				})
			})
		})
		Describe("standard processing", func() {
			Context("JSON Export", func() {
				It("should return no error", func() {
					config := Config{
						Certification:   "LATO",
						OpencontrolDir:  filepath.Join(workingDir, "..", "..", "fixtures", "opencontrol_fixtures_complete"),
						DestinationFile: "/tmp/output.json",
						OutputFormat:    jsonFormat,
					}
					err := Export(config)
					assert.Nil(GinkgoT(), err)
				})
			})
			Context("YAML Export", func() {
				It("should return no error", func() {
					config := Config{
						Certification:   "LATO",
						OpencontrolDir:  filepath.Join(workingDir, "..", "..", "fixtures", "opencontrol_fixtures_complete"),
						DestinationFile: "/tmp/output.yaml",
						OutputFormat:    yamlFormat,
					}
					err := Export(config)
					assert.Nil(GinkgoT(), err)
				})
			})
			Context("JSON Export with flattening", func() {
				It("should return no error", func() {
					config := Config{
						Certification:   "LATO",
						OpencontrolDir:  filepath.Join(workingDir, "..", "..", "fixtures", "opencontrol_fixtures_complete"),
						DestinationFile: "/tmp/output.json",
						OutputFormat:    jsonFormat,
						Flatten:         true,
						KeySeparator:    standardKeySeparator,
					}
					err := Export(config)
					assert.Nil(GinkgoT(), err)
				})
			})
			Context("JSON Export with flattening and key inference", func() {
				It("should return no error", func() {
					config := Config{
						Certification:   "LATO",
						OpencontrolDir:  filepath.Join(workingDir, "..", "..", "fixtures", "opencontrol_fixtures_complete"),
						DestinationFile: "/tmp/output.json",
						OutputFormat:    jsonFormat,
						Flatten:         true,
						InferKeys:       true,
						KeySeparator:    standardKeySeparator,
					}
					err := Export(config)
					assert.Nil(GinkgoT(), err)
				})
			})
			Context("JSON Export with flattening and key inference; docxtemplater support", func() {
				It("should return no error", func() {
					config := Config{
						Certification:   "LATO",
						OpencontrolDir:  filepath.Join(workingDir, "..", "..", "fixtures", "opencontrol_fixtures_complete"),
						DestinationFile: "/tmp/output.json",
						OutputFormat:    jsonFormat,
						Flatten:         true,
						InferKeys:       true,
						Docxtemplater:   true,
						KeySeparator:    standardKeySeparator,
					}
					err := Export(config)
					assert.Nil(GinkgoT(), err)
				})
			})
			Context("JSON Export with flattening and key inference; docxtemplater support and custom key separator", func() {
				It("should return no error", func() {
					config := Config{
						Certification:   "LATO",
						OpencontrolDir:  filepath.Join(workingDir, "..", "..", "fixtures", "opencontrol_fixtures_complete"),
						DestinationFile: "/tmp/output.json",
						OutputFormat:    jsonFormat,
						Flatten:         true,
						InferKeys:       true,
						Docxtemplater:   true,
						KeySeparator:    customKeySeparator,
					}
					err := Export(config)
					assert.Nil(GinkgoT(), err)
				})
			})
		})
	})
})
