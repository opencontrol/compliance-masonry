package schema

import (
	"errors"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/opencontrol/compliance-masonry/config/common"
	"github.com/opencontrol/compliance-masonry/config/common/resources"
	"github.com/opencontrol/compliance-masonry/config/common/resources/mocks"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Schema", func() {
	Describe("Parsing the schmema", func() {
		DescribeTable("parsing the schema with different data",
			func(data []byte, expectedSchema Schema, expectedErrorExists bool, expectedErrorText string) {
				s := Schema{}
				err := s.Parse(data)
				assert.Equal(GinkgoT(), expectedSchema, s)
				if expectedErrorExists {
					assert.Contains(GinkgoT(), err.Error(), ErrMalformedV1_0_0YamlPrefix)
				}
				assert.Equal(GinkgoT(), expectedErrorExists, err != nil)
			},
			Entry("good v1.0.0 data", []byte(`
schema_version: "1.0.0"
name: test
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
components:
  - ./component-1
  - ./component-2
  - ./component-3
certifications:
  - ./cert-1.yaml
standards:
  - ./standard-1.yaml
dependencies:
  certifications:
    - url: github.com/18F/LATO
      revision: master
  systems:
    - url: github.com/18F/cg-complinace
      revision: master
  standards:
    - url: github.com/18F/NIST-800-53
      revision: master
`),
				Schema{
					resourceGetter: resources.VCSAndLocalFSGetter{},
					Base:           common.Base{SchemaVersion: "1.0.0"},
					Name:           "test",
					Meta: Metadata{
						Description: "A system to test parsing",
						Maintainers: []string{
							"test@test.com",
						},
					},
					Components: []string{
						"./component-1",
						"./component-2",
						"./component-3",
					},
					Certifications: []string{
						"./cert-1.yaml",
					},
					Standards: []string{
						"./standard-1.yaml",
					},
					Dependencies: Dependencies{
						Certifications: []common.Entry{
							common.Entry{
								URL:      "github.com/18F/LATO",
								Revision: "master",
							},
						},
						Systems: []common.Entry{
							common.Entry{
								URL:      "github.com/18F/cg-complinace",
								Revision: "master",
							},
						},
						Standards: []common.Entry{
							common.Entry{
								URL:      "github.com/18F/NIST-800-53",
								Revision: "master",
							},
						},
					},
				}, false, ""),
			Entry("malformed yaml (tabbed over)", []byte(`
			schema_version: "1.0.0"
			system_name: test-system
			metadata:
			  description: "A system to test parsing"
			  maintainers:
			    - test@test.com
			components:
			  - ./component-1
			  - ./component-2
			  - ./component-3
			dependencies:
			  certification:
			    url: github.com/18F/LATO
			    revision: master
			  systems:
			    - url: github.com/18F/cg-complinace
			      revision: master
			  standards:
			    - url: github.com/18F/NIST-800-53
			      revision: master
			`), Schema{}, true, ErrMalformedV1_0_0YamlPrefix))
	})

	Describe("Getting resources", func() {
		var (
			getter                                                           *mocks.ResourceGetter
			dependentStandards, dependentCertifications, dependentComponents []common.Entry
			certifications, standards, components                            []string
			worker                                                           *common.ConfigWorker
			dependencies                                                     Dependencies
			destination                                                      = "."
			expectedError                                                    error
			s                                                                Schema
		)
		BeforeEach(func() {
			getter = new(mocks.ResourceGetter)
			worker = new(common.ConfigWorker)
			dependencies = Dependencies{Certifications: dependentCertifications, Systems: dependentComponents, Standards: dependentStandards}
			s = Schema{resourceGetter: getter, Dependencies: dependencies, Certifications: certifications, Standards: standards, Components: components}
		})
		It("should return an error when it's unable to get local certifications", func() {
			expectedError = errors.New("Cert error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, worker, constants.Certifications).Return(expectedError)
		})
		It("should return an error when it's unable to get local standards", func() {
			expectedError = errors.New("Standards error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, worker, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, worker, constants.Standards).Return(expectedError)
		})
		It("should return an error when it's unable to get local components", func() {
			expectedError = errors.New("Components error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, worker, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, worker, constants.Standards).Return(nil)
			getter.On("GetLocalResources", "", components, destination, constants.DefaultComponentsFolder, true, worker, constants.Components).Return(expectedError)
		})
		It("should return an error when it's unable to get remote certifications", func() {
			expectedError = errors.New("Remote cert error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, worker, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, worker, constants.Standards).Return(nil)
			getter.On("GetLocalResources", "", components, destination, constants.DefaultComponentsFolder, true, worker, constants.Components).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, worker, dependentCertifications).Return(expectedError)
		})
		It("should return an error when it's unable to get remote standards", func() {
			expectedError = errors.New("Remote standards error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, worker, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, worker, constants.Standards).Return(nil)
			getter.On("GetLocalResources", "", components, destination, constants.DefaultComponentsFolder, true, worker, constants.Components).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, worker, dependentCertifications).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultStandardsFolder, worker, dependentStandards).Return(expectedError)
		})
		It("should return an error when it's unable to get remote components", func() {
			expectedError = errors.New("Remote components error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, worker, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, worker, constants.Standards).Return(nil)
			getter.On("GetLocalResources", "", components, destination, constants.DefaultComponentsFolder, true, worker, constants.Components).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, worker, dependentCertifications).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultStandardsFolder, worker, dependentStandards).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultComponentsFolder, worker, dependentStandards).Return(expectedError)
		})
		It("should return no error when able to get all components", func() {
			expectedError = nil
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, worker, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, worker, constants.Standards).Return(nil)
			getter.On("GetLocalResources", "", components, destination, constants.DefaultComponentsFolder, true, worker, constants.Components).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, worker, dependentCertifications).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultStandardsFolder, worker, dependentStandards).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultComponentsFolder, worker, dependentStandards).Return(nil)
		})
		AfterEach(func() {
			err := s.GetResources("", destination, worker)
			assert.Equal(GinkgoT(), expectedError, err)
			getter.AssertExpectations(GinkgoT())
		})
	})

	DescribeTable("Testing GetLocalComponents",
		func(data []byte, expectedComponents []string) {
			s := Schema{}
			err := s.Parse(data)
			actualComponents := s.GetLocalComponents()
			log.Println(err)
			log.Println(actualComponents)
			assert.Equal(GinkgoT(), expectedComponents, actualComponents)
		},
		Entry("v1.0.0 data with no local components",
			[]byte(`
schema_version: "1.0.0"
system_name: test-system
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
dependencies:
  certification:
    url: github.com/18F/LATO
    revision: master
  systems:
    - url: github.com/18F/cg-complinace
      revision: master
  standards:
    - url: github.com/18F/NIST-800-53
      revision: master
`),
			nil,
		),
		Entry("v1.0.0 data with no local components",
			[]byte(`
schema_version: "1.0.0"
system_name: test-system
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
components:
  - ./component-1
  - ./component-2
  - ./component-3
dependencies:
  certification:
    url: github.com/18F/LATO
    revision: master
  systems:
    - url: github.com/18F/cg-complinace
      revision: master
  standards:
    - url: github.com/18F/NIST-800-53
      revision: master
`),
			[]string{"./component-1", "./component-2", "./component-3"},
		),
	)

	//TODO NEW TESTS HERE
	DescribeTable("Testing GetRequiredComponents",
		func(data []byte, expectedComponents []string) {
			s := Schema{}
			err := s.Parse(data)
			actualComponents := s.GetRequiredComponents()
			log.Println(err)
			log.Println(actualComponents)
			assert.Equal(GinkgoT(), expectedComponents, actualComponents)
		},
		Entry("v1.0.0 data with no required components",
			[]byte(`
 schema_version: "1.0.0"
 system_name: test-system
 metadata:
   description: "A system to test parsing"
   maintainers:
     - test@test.com
 dependencies:
   certification:
     url: github.com/18F/LATO
     revision: master
   systems:
     - url: github.com/18F/cg-complinace
       revision: master
   standards:
     - url: github.com/18F/NIST-800-53
       revision: master
 `),
			nil,
		),
		Entry("v1.0.0 data with required components",
			[]byte(`
 schema_version: "1.0.0"
 system_name: test-system
 metadata:
   description: "A system to test parsing"
   maintainers:
     - test@test.com
 required_components:
   - component-1
   - component-2
   - component-3
 dependencies:
   certification:
     url: github.com/18F/LATO
     revision: master
   systems:
     - url: github.com/18F/cg-complinace
       revision: master
   standards:
     - url: github.com/18F/NIST-800-53
       revision: master
 `),
			[]string{"component-1", "component-2", "component-3"},
		),
	)
	//TODO NEW TESTS END HERE
})
