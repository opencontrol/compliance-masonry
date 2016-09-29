package resources

import (
	"errors"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	resmocks "github.com/opencontrol/compliance-masonry/commands/get/resources/mocks"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/lib/common/mocks"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol"
	parserMocks "github.com/opencontrol/compliance-masonry/lib/opencontrol/mocks"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/tools/fs"
	fsmocks "github.com/opencontrol/compliance-masonry/tools/fs/mocks"
	"github.com/opencontrol/compliance-masonry/tools/mapset"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("ResourceGetter", func() {
	table.DescribeTable("GetResources",
		func(errs getterErrors) {
			dest := "."
			// get mocks for the getter and opencontrol yaml.
			getter, opencontrol := createMockGetterAndOpenControl(dest, errs)
			// Call GetResources
			err := GetResources("", dest, opencontrol, getter)
			// Make sure that we check the error.
			assert.Equal(GinkgoT(), errs.expectedError, err)
		},
		// Note: each of the variables not specified by the getterErrors init per entry will default to nil.
		// Thus, it's not necessary to be explicit with all of them.
		table.Entry("should return an error when it's unable to get local certifications", getterErrors{
			localCertError: errors.New("Cert error"),
			expectedError:  errors.New("Cert error"),
		}),
		table.Entry("should return an error when it's unable to get local standards", getterErrors{
			localStandardError: errors.New("Standards error"),
			expectedError:      errors.New("Standards error"),
		}),
		table.Entry("should return an error when it's unable to get local components", getterErrors{
			localComponentError: errors.New("Components error"),
			expectedError:       errors.New("Components error"),
		}),
		table.Entry("should return an error when it's unable to get remote certifications", getterErrors{
			remoteCertError: errors.New("Remote cert error"),
			expectedError:   errors.New("Remote cert error"),
		}),
		table.Entry("should return an error when it's unable to get remote standards", getterErrors{
			remoteStandardError: errors.New("Remote standards error"),
			expectedError:       errors.New("Remote standards error"),
		}),
		table.Entry("should return an error when it's unable to get remote components", getterErrors{
			remoteComponentError: errors.New("Remote components error"),
			expectedError:        errors.New("Remote components error"),
		}),
		table.Entry("should return no error when able to get all components", getterErrors{
		// everything is nil.
		}),
	)

	Describe("GetLocalResources", func() {
		table.DescribeTable("", func(recursively bool, initMap bool, resources []string, mkdirsError, copyError, copyAllError, expectedError error) {
			getter := vcsAndLocalFSGetter{}
			getter.Parser = createMockParser(nil)
			getter.FSUtil = createMockFSUtil(nil, nil, mkdirsError, copyError, copyAllError)
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
		table.DescribeTable("", func(downloadEntryError, tempDirError, openAndReadFileError, parserError, expectedError error) {
			// Override remoteSource with a mock.
			remoteSource := createMockRemoteSource()
			entries := []common.RemoteSource{remoteSource}

			// Setup getter
			getter := vcsAndLocalFSGetter{ResourceMap: mapset.Init()}

			// Override the fsutil with a mock.
			getter.FSUtil = createMockFSUtil(tempDirError, openAndReadFileError, nil, nil, nil)

			// Override downloader with a mock.
			getter.Downloader = createMockDownloader(remoteSource, downloadEntryError)

			// Override parser with a mock.
			getter.Parser = createMockParser(parserError)

			err := getter.GetRemoteResources("dest", "subfolder", entries)
			assert.Equal(GinkgoT(), expectedError, err)

		},
			table.Entry("fail to parse config", nil, nil, nil, errors.New("error parsing"), errors.New("error parsing")),
			table.Entry("fail to open and read file", nil, nil, errors.New("error reading file"), nil, errors.New("error reading file")),
			table.Entry("fail to download repo", errors.New("error downloading entry"), nil, nil, nil, errors.New("error downloading entry")),
			table.Entry("fail to create temp dir", nil, errors.New("error creating tempdir"), nil, nil, errors.New("error creating tempdir")),
		)
	})
})

func createMockFSUtil(tempDirError, openAndReadFileError, mkdirsError, copyError, copyAllError error) fs.Util {
	fsUtil := new(fsmocks.Util)
	fsUtil.On("TempDir", "", "opencontrol-resources").Return("sometempdir", tempDirError)
	data := []byte("schema_version: 1.0.0")
	fsUtil.On("OpenAndReadFile", mock.AnythingOfType("string")).Return(data, openAndReadFileError)
	fsUtil.On("Mkdirs", mock.AnythingOfType("string")).Return(mkdirsError)
	fsUtil.On("Copy", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(copyError)
	fsUtil.On("CopyAll", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(copyAllError)
	return fsUtil
}

func createMockDownloader(remoteSource common.RemoteSource, downloadEntryError error) Downloader {
	downloader := new(resmocks.Downloader)
	downloader.On("DownloadRepo", remoteSource, mock.AnythingOfType("string")).Return(downloadEntryError)
	return downloader
}

func createMockParser(parserError error) opencontrol.SchemaParser {
	schema := new(mocks.OpenControl)

	parser := new(parserMocks.SchemaParser)
	parser.On("Parse", mock.Anything).Return(schema, parserError)
	return parser
}

func createMockRemoteSource() common.RemoteSource {
	// Setup remoteSource mock
	remoteSource := new(mocks.RemoteSource)
	remoteSource.On("GetURL").Return("")
	remoteSource.On("GetConfigFile").Return("")
	return remoteSource
}

func createMockGetterAndOpenControl(destination string, errs getterErrors) (Getter, common.OpenControl) {
	// create the common variables used in both the opencontrol mock and getter mock.
	var dependentStandards, dependentCertifications, dependentComponents []common.RemoteSource
	var certifications, standards, components []string

	// Create the opencontrol mock
	opencontrol := new(mocks.OpenControl)
	opencontrol.On("GetCertifications").Return(certifications)
	opencontrol.On("GetStandards").Return(standards)
	opencontrol.On("GetComponents").Return(components)
	opencontrol.On("GetCertificationsDependencies").Return(dependentCertifications)
	opencontrol.On("GetStandardsDependencies").Return(dependentStandards)
	opencontrol.On("GetComponentsDependencies").Return(dependentComponents)

	// Create the getter mock.
	getter := new(resmocks.Getter)
	getter.On("GetLocalResources", "", certifications, destination,
		constants.DefaultCertificationsFolder, false, constants.Certifications).Return(errs.localCertError)
	getter.On("GetLocalResources", "", standards, destination,
		constants.DefaultStandardsFolder, false, constants.Standards).Return(errs.localStandardError)
	getter.On("GetLocalResources", "", components, destination,
		constants.DefaultComponentsFolder, true, constants.Components).Return(errs.localComponentError)
	getter.On("GetRemoteResources", destination, constants.DefaultCertificationsFolder,
		dependentCertifications).Return(errs.remoteCertError)
	getter.On("GetRemoteResources", destination, constants.DefaultStandardsFolder,
		dependentStandards).Return(errs.remoteStandardError)
	getter.On("GetRemoteResources", destination, constants.DefaultComponentsFolder,
		dependentStandards).Return(errs.remoteComponentError)
	return getter, opencontrol
}

type getterErrors struct {
	localCertError, localStandardError, localComponentError    error
	remoteCertError, remoteStandardError, remoteComponentError error
	expectedError                                              error
}
