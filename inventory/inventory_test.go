package inventory_test

import (
	. "github.com/opencontrol/compliance-masonry/inventory"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
)

var _ = Describe("Inventory", func() {
	Describe("Computing Gap Analysis", func() {
		var (
			workingDir string
		)
		BeforeEach(func() {
			workingDir, _ = os.Getwd()
		})
		Describe("bad inputs", func() {
			Context("When no certification is specified", func() {
				It("should return an empty slice and an error", func() {
					config := Config{}
					i, err := ComputeGapAnalysis(config)
					assert.Equal(GinkgoT(), []string{"Error: Missing Certification Argument"}, err)
					assert.Equal(GinkgoT(), 0, len(i.MissingControlList))
				})
			})
			Context("When bad / no folder location is given", func() {
				It("should return an empty slice and an error", func() {
					config := Config{Certification: "LATO"}
					i, err := ComputeGapAnalysis(config)
					assert.Equal(GinkgoT(), []string{"Error: `certifications` directory does exist"}, err)
					assert.Equal(GinkgoT(), 0, len(i.MissingControlList))
				})
			})
		})
		Context("When there controls specified in the certification but no controls have been documented", func() {
			It("should return the full list of controls", func() {
				config := Config{
					OpencontrolDir: filepath.Join(workingDir, "..", "fixtures", "opencontrol_fixtures"),
					Certification:  "LATO",
				}
				i, err := ComputeGapAnalysis(config)
				assert.Nil(GinkgoT(), err)
				assert.Equal(GinkgoT(), 2, len(i.MissingControlList))
			})
		})
		Context("When there controls specified in the certification and we have documented them", func() {
			It("should return no missing controls", func() {
				config := Config{
					OpencontrolDir: filepath.Join(workingDir, "..", "fixtures", "opencontrol_fixtures_complete"),
					Certification:  "LATO",
				}
				i, err := ComputeGapAnalysis(config)
				assert.Nil(GinkgoT(), err)
				assert.Equal(GinkgoT(), 0, len(i.MissingControlList))
			})
		})
	})
})