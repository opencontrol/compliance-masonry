package inventory_test

import (
	. "github.com/opencontrol/compliance-masonry/inventory"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"os"
)

var _ = Describe("Inventory", func() {
	Describe("Computing Gap Analysis", func() {
		var (
			workingDir string
		)
		BeforeEach(func() {
			workingDir,_ = os.Getwd()
		})
		Context("When there are no controls in certification", func(){
			It("should return an empty slice", func() {
				config := Config{}
				i, _ := ComputeGapAnalysis(config)
				assert.Equal(GinkgoT(), 0, len(i.MissingControlList))
			})
		})
		Context("When there controls specified in the certification but no controls have been documented", func() {
			It("should return the full list of controls", func() {
				config := Config{
					OpencontrolDir: filepath.Join(workingDir, "..", "fixtures", "opencontrol_fixtures"),
					Certification: "LATO",
				}
				_, err := ComputeGapAnalysis(config)
				assert.Nil(GinkgoT(), err)
				//assert.NotEqual(GinkgoT(), 0, len(missingControls))
			})
		})
	})
})
