package resources

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	resmocks "github.com/opencontrol/compliance-masonry/commands/get/resources/mocks"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/lib/common/mocks"
	parserMocks "github.com/opencontrol/compliance-masonry/lib/opencontrol/mocks"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	fsmocks "github.com/opencontrol/compliance-masonry/tools/fs/mocks"
	"github.com/opencontrol/compliance-masonry/tools/mapset"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vektra/errors"
)

var _ = Describe("ResourceGetter", func() {

	Describe("Getting resources", func() {
		var (
			getter                                                           *resmocks.Getter
			dependentStandards, dependentCertifications, dependentComponents []common.RemoteSource
			certifications, standards, components                            []string
			destination                                                      = "."
			expectedError                                                    error
			s                                                                *mocks.OpenControl
		)
		BeforeEach(func() {
			getter = new(resmocks.Getter)
			s = new(mocks.OpenControl)
			s.On("GetCertifications").Return(certifications)
			s.On("GetStandards").Return(standards)
			s.On("GetComponents").Return(components)
			s.On("GetCertificationsDependencies").Return(dependentCertifications)
			s.On("GetStandardsDependencies").Return(dependentStandards)
			s.On("GetComponentsDependencies").Return(dependentComponents)
		})
		It("should return an error when it's unable to get local certifications", func() {
			expectedError = errors.New("Cert error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, constants.Certifications).Return(expectedError)
		})
		It("should return an error when it's unable to get local standards", func() {
			expectedError = errors.New("Standards error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, constants.Standards).Return(expectedError)
		})
		It("should return an error when it's unable to get local components", func() {
			expectedError = errors.New("Components error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, constants.Standards).Return(nil)
			getter.On("GetLocalResources", "", components, destination, constants.DefaultComponentsFolder, true, constants.Components).Return(expectedError)
		})
		It("should return an error when it's unable to get remote certifications", func() {
			expectedError = errors.New("Remote cert error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, constants.Standards).Return(nil)
			getter.On("GetLocalResources", "", components, destination, constants.DefaultComponentsFolder, true, constants.Components).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, dependentCertifications).Return(expectedError)
		})
		It("should return an error when it's unable to get remote standards", func() {
			expectedError = errors.New("Remote standards error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, constants.Standards).Return(nil)
			getter.On("GetLocalResources", "", components, destination, constants.DefaultComponentsFolder, true, constants.Components).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, dependentCertifications).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultStandardsFolder, dependentStandards).Return(expectedError)
		})
		It("should return an error when it's unable to get remote components", func() {
			expectedError = errors.New("Remote components error")
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, constants.Standards).Return(nil)
			getter.On("GetLocalResources", "", components, destination, constants.DefaultComponentsFolder, true, constants.Components).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, dependentCertifications).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultStandardsFolder, dependentStandards).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultComponentsFolder, dependentStandards).Return(expectedError)
		})
		It("should return no error when able to get all components", func() {
			expectedError = nil
			getter.On("GetLocalResources", "", certifications, destination, constants.DefaultCertificationsFolder, false, constants.Certifications).Return(nil)
			getter.On("GetLocalResources", "", standards, destination, constants.DefaultStandardsFolder, false, constants.Standards).Return(nil)
			getter.On("GetLocalResources", "", components, destination, constants.DefaultComponentsFolder, true, constants.Components).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder, dependentCertifications).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultStandardsFolder, dependentStandards).Return(nil)
			getter.On("GetRemoteResources", destination, constants.DefaultComponentsFolder, dependentStandards).Return(nil)
		})
		AfterEach(func() {
			err := GetResources("", destination, s, getter)
			assert.Equal(GinkgoT(), expectedError, err)
			getter.AssertExpectations(GinkgoT())
		})
	})

	Describe("GetLocalResources", func() {
		table.DescribeTable("", func(recursively bool, initMap bool, resources []string, mkdirsError, copyError, copyAllError, expectedError error) {
			parser := new(parserMocks.SchemaParser)
			getter := vcsAndLocalFSGetter{}
			fsUtil := new(fsmocks.Util)
			fsUtil.On("Mkdirs", mock.AnythingOfType("string")).Return(mkdirsError)
			fsUtil.On("Copy", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(copyError)
			fsUtil.On("CopyAll", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(copyAllError)
			getter.Parser = parser
			getter.FSUtil = fsUtil
			getter.ResourceMap = mapset.Init()
			err := getter.GetLocalResources("", resources, "dest", "subfolder", recursively, constants.Standards)
			assert.Equal(GinkgoT(), expectedError, err)
		},
			table.Entry("Bad input to reserve", false, true, []string{""}, nil, nil, nil, mapset.ErrEmptyInput),
			table.Entry("Successful recursive copy", true, true, []string{"res"}, nil, nil, nil, nil),
			table.Entry("Successful single copy", false, true, []string{"res"}, nil, nil, nil, nil),
			table.Entry("Failure of single copy", false, true, []string{"res"}, nil, errors.New("single copy fail"), nil, errors.New("single copy fail")),
			table.Entry("Mkdirs", false, true, []string{"res"}, errors.New("mkdirs error"), nil, nil, errors.New("mkdirs error")),
		)
	})
	Describe("GetRemoteResources", func() {
		table.DescribeTable("", func(downloadEntryError, tempDirError, openAndReadFileError, getResourcesError, parserError, expectedError error) {
			// Setup remoteSource mock
			remoteSource := new(mocks.RemoteSource)
			remoteSource.On("GetURL").Return("")
			remoteSource.On("GetConfigFile").Return("")
			entries := []common.RemoteSource{remoteSource}
			// Setup getter
			getter := vcsAndLocalFSGetter{ResourceMap: mapset.Init()}

			// Override the fsutil with a mock
			fsUtil := new(fsmocks.Util)
			fsUtil.On("TempDir", "", "opencontrol-resources").Return("sometempdir", tempDirError)
			data := []byte("schema_version: 1.0.0")
			fsUtil.On("OpenAndReadFile", mock.AnythingOfType("string")).Return(data, openAndReadFileError)
			getter.FSUtil = fsUtil

			downloader := new(resmocks.Downloader)
			downloader.On("DownloadRepo", remoteSource, mock.AnythingOfType("string")).Return(downloadEntryError)
			getter.Downloader = downloader

			schema := new(mocks.OpenControl)

			parser := new(parserMocks.SchemaParser)
			parser.On("Parse", mock.Anything).Return(schema, parserError)
			getter.Parser = parser

			err := getter.GetRemoteResources("dest", "subfolder", entries)
			assert.Equal(GinkgoT(), expectedError, err)

		},
			//table.Entry("success", nil, nil, nil, nil, nil, nil),
			//table.Entry("fail to get resources", nil, nil, nil, errors.New("error getting resources"), nil, errors.New("error getting resources")),
			table.Entry("fail to parse config", nil, nil, nil, nil, errors.New("error parsing"), errors.New("error parsing")),
			table.Entry("fail to open and read file", nil, nil, errors.New("error reading file"), nil, nil, errors.New("error reading file")),
			table.Entry("fail to download repo", errors.New("error downloading entry"), nil, nil, nil, nil, errors.New("error downloading entry")),
			table.Entry("fail to create temp dir", nil, errors.New("error creating tempdir"), nil, nil, nil, errors.New("error creating tempdir")),
		)
	})
})
