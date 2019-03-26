/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package implementationstatus_test

import (
	. "github.com/opencontrol/compliance-masonry/pkg/cli/implementationstatus"

	"errors"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
)

var _ = Describe("Implementationstatus", func() {
	Describe("Finding implementation_status", func() {
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
					i, err := FindImplementationStatus(config, "partial")
					assert.Equal(GinkgoT(), []error{errors.New("Error: Missing Certification Argument")}, err)
					assert.Equal(GinkgoT(), 0, len(i.SatisfiesList))
				})
			})
			Context("When bad / no folder location is given", func() {
				It("should return an empty slice and an error", func() {
					config := Config{Certification: "LATO"}
					i, err := FindImplementationStatus(config, "partial")
					assert.Equal(GinkgoT(), []error{errors.New("Error: `certifications` directory does exist")}, err)
					assert.Equal(GinkgoT(), 0, len(i.SatisfiesList))
				})
			})
		})
		Context("When we search for an implementation_status", func() {
			It("should find at least one component in our test data", func() {
				config := Config{
					OpencontrolDir: filepath.Join(workingDir, "..", "..", "..", "test", "fixtures", "opencontrol_fixtures"),
					Certification:  "LATO",
				}
				i, err := FindImplementationStatus(config, "partial")
				assert.Nil(GinkgoT(), err)
				assert.NotEqual(GinkgoT(), 0, len(i.ComponentList))
			})
		})
		Context("When we search for the 'partial' implementation_status", func() {
			It("should find more than one in our test data", func() {
				config := Config{
					OpencontrolDir: filepath.Join(workingDir, "..", "..", "..", "test", "fixtures", "opencontrol_fixtures"),
					Certification:  "LATO",
				}
				i, err := FindImplementationStatus(config, "partial")
				assert.Nil(GinkgoT(), err)
				// I have no idea why assert.Greater is not working, but this will work.
				assert.NotEqual(GinkgoT(), 1, len(i.SatisfiesList))
				assert.NotEqual(GinkgoT(), 0, len(i.SatisfiesList))
			})
		})
		Context("When we search for the 'planned' implementation_status", func() {
			It("should only find one in our test data", func() {
				config := Config{
					OpencontrolDir: filepath.Join(workingDir, "..", "..", "..", "test", "fixtures", "opencontrol_fixtures"),
					Certification:  "LATO",
				}
				i, err := FindImplementationStatus(config, "planned")
				assert.Nil(GinkgoT(), err)
				assert.Equal(GinkgoT(), 1, len(i.SatisfiesList))
			})
		})
	})
})
