package mocks

import "github.com/stretchr/testify/mock"

import "github.com/opencontrol/compliance-masonry/config/common"
import "github.com/opencontrol/compliance-masonry/tools/constants"

type ResourceGetter struct {
	mock.Mock
}

// GetLocalResources provides a mock function with given fields: source, resources, destination, subfolder, recursively, worker, resourceType
func (_m *ResourceGetter) GetLocalResources(source string, resources []string, destination string, subfolder string, recursively bool, worker *common.ConfigWorker, resourceType constants.ResourceType) error {
	ret := _m.Called(source, resources, destination, subfolder, recursively, worker, resourceType)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, string, string, bool, *common.ConfigWorker, constants.ResourceType) error); ok {
		r0 = rf(source, resources, destination, subfolder, recursively, worker, resourceType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRemoteResources provides a mock function with given fields: destination, subfolder, worker, entries
func (_m *ResourceGetter) GetRemoteResources(destination string, subfolder string, worker *common.ConfigWorker, entries []common.Entry) error {
	ret := _m.Called(destination, subfolder, worker, entries)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *common.ConfigWorker, []common.Entry) error); ok {
		r0 = rf(destination, subfolder, worker, entries)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
