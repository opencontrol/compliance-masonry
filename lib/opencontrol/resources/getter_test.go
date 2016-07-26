package resources

import (

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol/versions/base"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	basemocks "github.com/opencontrol/compliance-masonry/lib/opencontrol/versions/base/mocks"
	fsmocks "github.com/opencontrol/compliance-masonry/tools/fs/mocks"
	"github.com/opencontrol/compliance-masonry/tools/mapset"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vektra/errors"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

var _ = Describe("ResourceGetter", func() {

	Describe("GetLocalResources", func() {
		var (
			resMap mapset.MapSet
		)
		table.DescribeTable("", func(recursively bool, initMap bool, resources []string, mkdirsError, copyError, copyAllError, expectedError error) {
			getter := NewVCSAndLocalGetter()
			fsUtil := new(fsmocks.Util)
			fsUtil.On("Mkdirs", mock.AnythingOfType("string")).Return(mkdirsError)
			fsUtil.On("Copy", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(copyError)
			fsUtil.On("CopyAll", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(copyAllError)
			if initMap {
				resMap = mapset.Init()
			}
			worker := new(base.Worker)
			worker.ResourceMap = resMap
			worker.FSUtil = fsUtil
			err := getter.GetLocalResources("", resources, "dest", "subfolder", recursively, worker, constants.Standards)
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
		table.DescribeTable("", func(downloadEntryError, tempDirError, openAndReadFileError, getResourcesError, parseV1_0_0Error, expectedError error) {
			entries := []common.Entry{
				{
					Path: "",
				},
			}
			getter := vcsAndLocalFSGetter{}
			worker := new(base.Worker)
			downloader := new(basemocks.EntryDownloader)
			downloader.On("DownloadEntry", entries[0], mock.AnythingOfType("string")).Return(downloadEntryError)
			getter.Downloader = downloader
			fsUtil := new(fsmocks.Util)
			fsUtil.On("TempDir", "", "opencontrol-resources").Return("sometempdir", tempDirError)
			data := []byte("schema_version: 1.0.0")
			fsUtil.On("OpenAndReadFile", mock.AnythingOfType("string")).Return(data, openAndReadFileError)
			parser := new(basemocks.SchemaParser)
			schema := new(basemocks.OpenControl)
			schema.On("GetResources", mock.AnythingOfType("string"), mock.AnythingOfType("string"), worker).Return(getResourcesError)
			parser.On("ParseV1_0_0", data).Return(schema, parseV1_0_0Error)
			worker.Parser = parser
			worker.FSUtil = fsUtil
			err := getter.GetRemoteResources("dest", "subfolder", worker, entries)
			assert.Equal(GinkgoT(), expectedError, err)

		},
			table.Entry("success", nil, nil, nil, nil, nil, nil),
			table.Entry("fail to get resources", nil, nil, nil, errors.New("error getting resources"), nil, errors.New("error getting resources")),
			table.Entry("fail to parse config", nil, nil, nil, nil, errors.New("error parsing"), errors.New("error parsing")),
			table.Entry("fail to open and read file", nil, nil, errors.New("error reading file"), nil, nil, errors.New("error reading file")),
			table.Entry("fail to download repo", errors.New("error downloading entry"), nil, nil, nil, nil, errors.New("error downloading entry")),
			table.Entry("fail to create temp dir", nil, errors.New("error creating tempdir"), nil, nil, nil, errors.New("error creating tempdir")),
		)
	})
})
