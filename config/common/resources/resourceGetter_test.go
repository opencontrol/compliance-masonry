package resources

import (
	"github.com/opencontrol/compliance-masonry-go/config/common"
	"github.com/opencontrol/compliance-masonry-go/config/common/mocks"
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
	fsmocks "github.com/opencontrol/compliance-masonry-go/tools/fs/mocks"
	"github.com/opencontrol/compliance-masonry-go/tools/mapset"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vektra/errors"
	"testing"
)

func TestGetLocalResources(t *testing.T) {
	getter := VCSAndLocalFSGetter{}
	resources := []string{
		"",
	}
	// Bad Input to Reserve
	resMap := mapset.Init()
	worker := new(common.ConfigWorker)
	worker.ResourceMap = resMap
	err := getter.GetLocalResources("", resources, "dest", "subfolder", false, worker, constants.Standards)
	assert.NotNil(t, err)

	// Try Recursively copy success
	resources = []string{
		"res",
	}
	fsUtil := new(fsmocks.Util)
	fsUtil.On("Mkdirs", mock.AnythingOfType("string")).Return(nil)
	fsUtil.On("Copy", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
	worker.FSUtil = fsUtil
	err = getter.GetLocalResources("", resources, "dest", "subfolder", false, worker, constants.Standards)
	assert.Nil(t, err)

	// Try single copy success.
	resMap = mapset.Init()
	worker.ResourceMap = resMap
	fsUtil = new(fsmocks.Util)
	fsUtil.On("Mkdirs", mock.AnythingOfType("string")).Return(nil)
	fsUtil.On("CopyAll", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)
	worker.FSUtil = fsUtil
	err = getter.GetLocalResources("", resources, "dest", "subfolder", true, worker, constants.Standards)
	assert.Nil(t, err)

	// Try single copy fail.
	resMap = mapset.Init()
	worker.ResourceMap = resMap
	fsUtil = new(fsmocks.Util)
	fsUtil.On("Mkdirs", mock.AnythingOfType("string")).Return(nil)
	expectedError := errors.New("single copy error")
	fsUtil.On("CopyAll", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(expectedError)
	worker.FSUtil = fsUtil
	err = getter.GetLocalResources("", resources, "dest", "subfolder", true, worker, constants.Standards)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)

	// Try mkdirs failure
	resMap = mapset.Init()
	worker.ResourceMap = resMap
	resources = []string{
		"res",
	}
	expectedError = errors.New("mkdirs error")
	fsUtil = new(fsmocks.Util)
	fsUtil.On("Mkdirs", mock.AnythingOfType("string")).Return(expectedError)
	worker.FSUtil = fsUtil
	err = getter.GetLocalResources("", resources, "dest", "subfolder", false, worker, constants.Standards)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}

func TestGetRemoteResources(t *testing.T) {
	// Success case.
	getter := VCSAndLocalFSGetter{}
	worker := new(common.ConfigWorker)
	downloader := new(mocks.EntryDownloader)
	entries := []common.Entry{
		{
			Path: "",
			URL:  "",
		},
	}
	downloader.On("DownloadEntry", entries[0], mock.AnythingOfType("string")).Return(nil)
	worker.Downloader = downloader
	fsUtil := new(fsmocks.Util)
	fsUtil.On("TempDir", "", "opencontrol-resources").Return("sometempdir", nil)
	data := []byte("schema_version: 1.0.0")
	fsUtil.On("OpenAndReadFile", mock.AnythingOfType("string")).Return(data, nil)
	parser := new(mocks.SchemaParser)
	schema := new(mocks.BaseSchema)
	schema.On("GetResources", mock.AnythingOfType("string"), mock.AnythingOfType("string"), worker).Return(nil)
	parser.On("ParseV1_0_0", data).Return(schema, nil)
	worker.Parser = parser
	worker.FSUtil = fsUtil
	err := getter.GetRemoteResources("dest", "subfolder", worker, entries)
	assert.Nil(t, err)

	// Fail to GetResources
	getter = VCSAndLocalFSGetter{}
	worker = new(common.ConfigWorker)
	downloader = new(mocks.EntryDownloader)
	downloader.On("DownloadEntry", entries[0], mock.AnythingOfType("string")).Return(nil)
	worker.Downloader = downloader
	fsUtil = new(fsmocks.Util)
	fsUtil.On("TempDir", "", "opencontrol-resources").Return("sometempdir", nil)
	data = []byte("schema_version: 1.0.0")
	fsUtil.On("OpenAndReadFile", mock.AnythingOfType("string")).Return(data, nil)
	parser = new(mocks.SchemaParser)
	schema = new(mocks.BaseSchema)
	expectedError := errors.New("error getting resources")
	schema.On("GetResources", mock.AnythingOfType("string"), mock.AnythingOfType("string"), worker).Return(expectedError)
	parser.On("ParseV1_0_0", data).Return(schema, nil)
	worker.Parser = parser
	worker.FSUtil = fsUtil
	err = getter.GetRemoteResources("dest", "subfolder", worker, entries)
	assert.NotNil(t, err)
	assert.Equal(t, err, expectedError)

	// Fail to Parse config
	getter = VCSAndLocalFSGetter{}
	worker = new(common.ConfigWorker)
	downloader = new(mocks.EntryDownloader)
	downloader.On("DownloadEntry", entries[0], mock.AnythingOfType("string")).Return(nil)
	worker.Downloader = downloader
	fsUtil = new(fsmocks.Util)
	fsUtil.On("TempDir", "", "opencontrol-resources").Return("sometempdir", nil)
	data = []byte("schema_version: 1.0.0")
	fsUtil.On("OpenAndReadFile", mock.AnythingOfType("string")).Return(data, nil)
	parser = new(mocks.SchemaParser)
	expectedError = errors.New("error parsing")
	parser.On("ParseV1_0_0", data).Return(schema, expectedError)
	worker.Parser = parser
	worker.FSUtil = fsUtil
	err = getter.GetRemoteResources("dest", "subfolder", worker, entries)
	assert.NotNil(t, err)
	assert.Equal(t, err, expectedError)

	// Fail to open file
	getter = VCSAndLocalFSGetter{}
	worker = new(common.ConfigWorker)
	downloader = new(mocks.EntryDownloader)
	downloader.On("DownloadEntry", entries[0], mock.AnythingOfType("string")).Return(nil)
	worker.Downloader = downloader
	fsUtil = new(fsmocks.Util)
	fsUtil.On("TempDir", "", "opencontrol-resources").Return("sometempdir", nil)
	data = []byte("schema_version: 1.0.0")
	expectedError = errors.New("error reading file")
	fsUtil.On("OpenAndReadFile", mock.AnythingOfType("string")).Return(data, expectedError)
	worker.FSUtil = fsUtil
	err = getter.GetRemoteResources("dest", "subfolder", worker, entries)
	assert.NotNil(t, err)
	assert.Equal(t, err, expectedError)

	// Fail to download repo
	getter = VCSAndLocalFSGetter{}
	worker = new(common.ConfigWorker)
	downloader = new(mocks.EntryDownloader)
	fsUtil = new(fsmocks.Util)
	fsUtil.On("TempDir", "", "opencontrol-resources").Return("sometempdir", nil)
	worker.FSUtil = fsUtil
	expectedError = errors.New("error downloading entry")
	downloader.On("DownloadEntry", entries[0], mock.AnythingOfType("string")).Return(expectedError)
	worker.Downloader = downloader
	err = getter.GetRemoteResources("dest", "subfolder", worker, entries)
	assert.NotNil(t, err)
	assert.Equal(t, err, expectedError)

	// Fail to create temp dir
	getter = VCSAndLocalFSGetter{}
	worker = new(common.ConfigWorker)
	downloader = new(mocks.EntryDownloader)
	expectedError = errors.New("error creating tempdir")
	fsUtil = new(fsmocks.Util)
	fsUtil.On("TempDir", "", "opencontrol-resources").Return("sometempdir", expectedError)
	worker.FSUtil = fsUtil
	err = getter.GetRemoteResources("dest", "subfolder", worker, entries)
	assert.NotNil(t, err)
	assert.Equal(t, err, expectedError)
}
