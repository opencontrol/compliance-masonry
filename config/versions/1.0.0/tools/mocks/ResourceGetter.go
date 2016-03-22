package mocks

import "github.com/stretchr/testify/mock"

import "github.com/opencontrol/compliance-masonry-go/config/common"

type ResourceGetter struct {
	mock.Mock
}

// GetLocalResources provides a mock function with given fields: resources, destination, subfolder, recursively
func (_m *ResourceGetter) GetLocalResources(resources []string, destination string, subfolder string, recursively bool) error {
	ret := _m.Called(resources, destination, subfolder, recursively)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string, string, string, bool) error); ok {
		r0 = rf(resources, destination, subfolder, recursively)
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
