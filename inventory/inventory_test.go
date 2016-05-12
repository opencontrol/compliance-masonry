package inventory_test

import (
	. "github.com/opencontrol/compliance-masonry/inventory"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"path/filepath"
)

var _ = Describe("Inventory", func() {
	Describe("Computing Gap Analysis", func() {
		Context("When there are no controls in certification", func(){
			It("should return an empty slice", func() {
				i := Inventory{}
				config := Config{}
				missingControls := i.ComputeGapAnalysis(config)
				assert.Equal(GinkgoT(), 0, len(missingControls))
			})
		})
		Context("When there controls specified in the certification but no controls have been documented", func()) {
			It("should return the full list of controls", func() {
				i := Inventory{}
				config := Config{
					OpencontrolDir: filepath.Join("fixtures", "opencontrol_fixtures"),
					Certification: "LATO",
				}
				missingControls := i.ComputeGapAnalysis(config)
				assert.Equal(GinkgoT(), 0, len(missingControls))
			})
		}
	})
})
