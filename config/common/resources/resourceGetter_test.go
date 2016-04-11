package resources_test

import (
	. "github.com/opencontrol/compliance-masonry/config/common/resources"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/opencontrol/compliance-masonry/config/common"
	"github.com/opencontrol/compliance-masonry/config/common/mocks"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	fsmocks "github.com/opencontrol/compliance-masonry/tools/fs/mocks"
	"github.com/opencontrol/compliance-masonry/tools/mapset"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vektra/errors"
)

var _ = Describe("ResourceGetter", func() {

	Describe("GetLocalResources", func() {
		var (
			resMap mapset.MapSet
		)
		DescribeTable("", func(recursively bool, initMap bool, resources []string, mkdirsError, copyError, copyAllError, expectedError error) {
			getter := VCSAndLocalFSGetter{}
			fsUtil := new(fsmocks.Util)
			fsUtil.On("Mkdirs", mock.AnythingOfType("string")).Return(mkdirsError)
			fsUtil.On("Copy", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(copyError)
			fsUtil.On("CopyAll", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(copyAllError)
			if initMap {
				resMap = mapset.Init()
			}
			worker := new(common.ConfigWorker)
			worker.ResourceMap = resMap
			worker.FSUtil = fsUtil
			err := getter.GetLocalResources("", resources, "dest", "subfolder", recursively, worker, constants.Standards)
			assert.Equal(GinkgoT(), expectedError, err)
		},
			Entry("Bad input to reserve", false, true, []string{""}, nil, nil, nil, mapset.ErrEmptyInput),
			Entry("Successful recursive copy", true, true, []string{"res"}, nil, nil, nil, nil),
			Entry("Successful single copy", false, true, []string{"res"}, nil, nil, nil, nil),
			Entry("Failure of single copy", false, true, []string{"res"}, nil, errors.New("single copy fail"), nil, errors.New("single copy fail")),
			Entry("Mkdirs", false, true, []string{"res"}, errors.New("mkdirs error"), nil, nil, errors.New("mkdirs error")),
		)
	})
	Describe("GetRemoteResources", func() {
		DescribeTable("", func(downloadEntryError, tempDirError, openAndReadFileError, getResourcesError, parseV1_0_0Error, expectedError error) {
			entries := []common.Entry{
				{
					Path: "",
				},
			}
			getter := VCSAndLocalFSGetter{}
			worker := new(common.ConfigWorker)
			downloader := new(mocks.EntryDownloader)
			downloader.On("DownloadEntry", entries[0], mock.AnythingOfType("string")).Return(downloadEntryError)
			worker.Downloader = downloader
			fsUtil := new(fsmocks.Util)
			fsUtil.On("TempDir", "", "opencontrol-resources").Return("sometempdir", tempDirError)
			data := []byte("schema_version: 1.0.0")
			fsUtil.On("OpenAndReadFile", mock.AnythingOfType("string")).Return(data, openAndReadFileError)
			parser := new(mocks.SchemaParser)
			schema := new(mocks.BaseSchema)
			schema.On("GetResources", mock.AnythingOfType("string"), mock.AnythingOfType("string"), worker).Return(getResourcesError)
			parser.On("ParseV1_0_0", data).Return(schema, parseV1_0_0Error)
			worker.Parser = parser
			worker.FSUtil = fsUtil
			err := getter.GetRemoteResources("dest", "subfolder", worker, entries)
			assert.Equal(GinkgoT(), expectedError, err)

		},
			Entry("success", nil, nil, nil, nil, nil, nil),
			Entry("fail to get resources", nil, nil, nil, errors.New("error getting resources"), nil, errors.New("error getting resources")),
			Entry("fail to parse config", nil, nil, nil, nil, errors.New("error parsing"), errors.New("error parsing")),
			Entry("fail to open and read file", nil, nil, errors.New("error reading file"), nil, nil, errors.New("error reading file")),
			Entry("fail to download repo", errors.New("error downloading entry"), nil, nil, nil, nil, errors.New("error downloading entry")),
			Entry("fail to create temp dir", nil, errors.New("error creating tempdir"), nil, nil, nil, errors.New("error creating tempdir")),
		)
	})
})
