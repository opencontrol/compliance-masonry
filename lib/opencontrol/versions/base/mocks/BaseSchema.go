package mocks

import "github.com/opencontrol/compliance-masonry/config/common"
import "github.com/stretchr/testify/mock"

type BaseSchema struct {
	mock.Mock
}

// Parse provides a mock function with given fields: data
func (_m *BaseSchema) Parse(data []byte) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSchemaVersion provides a mock function with given fields:
func (_m *BaseSchema) GetSchemaVersion() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetResources provides a mock function with given fields: _a0, _a1, _a2
func (_m *BaseSchema) GetResources(_a0 string, _a1 string, _a2 *common.ConfigWorker) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *common.ConfigWorker) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
