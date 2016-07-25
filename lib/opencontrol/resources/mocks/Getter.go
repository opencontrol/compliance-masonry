package mocks

import (
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol/versions/base"
	"github.com/stretchr/testify/mock"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

type Getter struct {
	mock.Mock
}

// GetLocalResources provides a mock function with given fields: source, resources, destination, subfolder, recursively, worker, resourceType
func (_m *Getter) GetLocalResources(source string, resources []string, destination string, subfolder string, recursively bool, worker *base.Worker, resourceType constants.ResourceType) error {
	ret := _m.Called(source, resources, destination, subfolder, recursively, worker, resourceType)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string, string, string, bool, *base.Worker, constants.ResourceType) error); ok {
		r0 = rf(source, resources, destination, subfolder, recursively, worker, resourceType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRemoteResources provides a mock function with given fields: destination, subfolder, worker, entries
func (_m *Getter) GetRemoteResources(destination string, subfolder string, worker *base.Worker, entries []common.Entry) error {
	ret := _m.Called(destination, subfolder, worker, entries)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *base.Worker, []common.Entry) error); ok {
		r0 = rf(destination, subfolder, worker, entries)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
