package resources_test

import (
	. "github.com/opencontrol/compliance-masonry-go/config/common/resources"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/opencontrol/compliance-masonry-go/tools/mapset"
	"github.com/opencontrol/compliance-masonry-go/config/common"
	fsmocks "github.com/opencontrol/compliance-masonry-go/tools/fs/mocks"
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vektra/errors"
)

var _ = Describe("ResourceGetter", func() {

	Describe("GetLocalResources", func() {
		var (
			resMap mapset.MapSet
		)
		DescribeTable("GetSchemaVersion", func(recursively bool, initMap bool, resources []string, mkdirsError, copyError, copyAllError, expectedError error) {
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
})
