package schema

import (
	"errors"
	"github.com/opencontrol/compliance-masonry-go/config/common"
	"github.com/opencontrol/compliance-masonry-go/config/versions/1.0.0/tools"
	"github.com/opencontrol/compliance-masonry-go/config/versions/1.0.0/tools/mocks"
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	var parseTests = []struct {
		data                []byte
		expectedSchema      Schema
		expectedErrorExists bool
		expectedErrorText   string
	}{
		{
			data: []byte(`
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
			expectedSchema: Schema{
				resourceGetter: tools.VCSAndLocalFSGetter{},
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
			},
			expectedErrorExists: false,
		},
		{
			// Malformed yaml. Tabbed over.
			data: []byte(`
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
			expectedSchema:      Schema{},
			expectedErrorExists: true,
			expectedErrorText:   ErrMalformedV1_0_0YamlPrefix,
		},
	}
	for _, test := range parseTests {
		s := Schema{}
		err := s.Parse(test.data)
		assert.Equal(t, test.expectedSchema, s)
		if test.expectedErrorExists {
			assert.Contains(t, err.Error(), ErrMalformedV1_0_0YamlPrefix)
		}
		assert.Equal(t, test.expectedErrorExists, err != nil)
	}
}

func TestGetResources(t *testing.T) {
	dependentStandards := []common.Entry{}
	dependentCertifications := []common.Entry{}
	dependentComponents := []common.Entry{}
	dependencies := Dependencies{Certifications: dependentCertifications, Systems: dependentComponents, Standards: dependentStandards}
	certifications := []string{}
	standards := []string{}
	components := []string{}
	worker := new(common.ConfigWorker)
	destination := "."

	// Local Cert error
	getter := new(mocks.ResourceGetter)
	expectedError := errors.New("Cert error")
	getter.On("GetLocalResources", certifications, destination, constants.DefaultCertificationsFolder, false).Return(expectedError)

	s := Schema{resourceGetter: getter, Dependencies: dependencies, Certifications: certifications, Standards: standards, Components: components}
	err := s.GetResources(destination, worker)
	assert.Equal(t, expectedError, err)
	getter.AssertExpectations(t)

	// Local standards error
	getter = new(mocks.ResourceGetter)
	expectedError = errors.New("Standards error")
	getter.On("GetLocalResources", certifications, destination, constants.DefaultCertificationsFolder, false).Return(nil)
	getter.On("GetLocalResources", standards, destination, constants.DefaultStandardsFolder, false).Return(expectedError)

	s = Schema{resourceGetter: getter, Dependencies: dependencies, Certifications: certifications, Standards: standards, Components: components}
	err = s.GetResources(destination, worker)
	assert.Equal(t, expectedError, err)
	getter.AssertExpectations(t)

	// Local components error
	getter = new(mocks.ResourceGetter)
	expectedError = errors.New("Components error")
	getter.On("GetLocalResources", certifications, destination, constants.DefaultCertificationsFolder, false).Return(nil)
	getter.On("GetLocalResources", standards, destination, constants.DefaultStandardsFolder, false).Return(nil)
	getter.On("GetLocalResources", components, destination, constants.DefaultComponentsFolder, true).Return(expectedError)

	s = Schema{resourceGetter: getter, Dependencies: dependencies, Certifications: certifications, Standards: standards, Components: components}
	err = s.GetResources(destination, worker)
	assert.Equal(t, expectedError, err)
	getter.AssertExpectations(t)

	// Remote Certifications error
	getter = new(mocks.ResourceGetter)
	expectedError = errors.New("Remote cert error")
	getter.On("GetLocalResources", certifications, destination, constants.DefaultCertificationsFolder, false).Return(nil)
	getter.On("GetLocalResources", standards, destination, constants.DefaultStandardsFolder, false).Return(nil)
	getter.On("GetLocalResources", components, destination, constants.DefaultComponentsFolder, true).Return(nil)
	getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, worker, dependentCertifications).Return(expectedError)

	s = Schema{resourceGetter: getter, Dependencies: dependencies, Certifications: certifications, Standards: standards, Components: components}
	err = s.GetResources(destination, worker)
	assert.Equal(t, expectedError, err)
	getter.AssertExpectations(t)

	// Remote standards error
	getter = new(mocks.ResourceGetter)
	expectedError = errors.New("Remote standards error")
	getter.On("GetLocalResources", certifications, destination, constants.DefaultCertificationsFolder, false).Return(nil)
	getter.On("GetLocalResources", standards, destination, constants.DefaultStandardsFolder, false).Return(nil)
	getter.On("GetLocalResources", components, destination, constants.DefaultComponentsFolder, true).Return(nil)
	getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, worker, dependentCertifications).Return(nil)
	getter.On("GetRemoteResources", destination, constants.DefaultStandardsFolder, worker, dependentStandards).Return(expectedError)

	s = Schema{resourceGetter: getter, Dependencies: dependencies, Certifications: certifications, Standards: standards, Components: components}
	err = s.GetResources(destination, worker)
	assert.Equal(t, expectedError, err)
	getter.AssertExpectations(t)

	// Remote standards error
	getter = new(mocks.ResourceGetter)
	expectedError = errors.New("Remote components error")
	getter.On("GetLocalResources", certifications, destination, constants.DefaultCertificationsFolder, false).Return(nil)
	getter.On("GetLocalResources", standards, destination, constants.DefaultStandardsFolder, false).Return(nil)
	getter.On("GetLocalResources", components, destination, constants.DefaultComponentsFolder, true).Return(nil)
	getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, worker, dependentCertifications).Return(nil)
	getter.On("GetRemoteResources", destination, constants.DefaultStandardsFolder, worker, dependentStandards).Return(nil)
	getter.On("GetRemoteResources", destination, constants.DefaultComponentsFolder, worker, dependentStandards).Return(expectedError)

	s = Schema{resourceGetter: getter, Dependencies: dependencies, Certifications: certifications, Standards: standards, Components: components}
	err = s.GetResources(destination, worker)
	assert.Equal(t, expectedError, err)
	getter.AssertExpectations(t)

	// no error
	getter = new(mocks.ResourceGetter)
	expectedError = nil
	getter.On("GetLocalResources", certifications, destination, constants.DefaultCertificationsFolder, false).Return(nil)
	getter.On("GetLocalResources", standards, destination, constants.DefaultStandardsFolder, false).Return(nil)
	getter.On("GetLocalResources", components, destination, constants.DefaultComponentsFolder, true).Return(nil)
	getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, worker, dependentCertifications).Return(nil)
	getter.On("GetRemoteResources", destination, constants.DefaultStandardsFolder, worker, dependentStandards).Return(nil)
	getter.On("GetRemoteResources", destination, constants.DefaultComponentsFolder, worker, dependentStandards).Return(nil)

	s = Schema{resourceGetter: getter, Dependencies: dependencies, Certifications: certifications, Standards: standards, Components: components}
	err = s.GetResources(destination, worker)
	assert.Equal(t, expectedError, err)
	getter.AssertExpectations(t)
}
