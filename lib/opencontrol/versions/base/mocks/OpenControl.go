package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol/versions/base"
)

type OpenControl struct {
	mock.Mock
}

// Parse provides a mock function with given fields: data
func (_m *OpenControl) Parse(data []byte) error {
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
func (_m *OpenControl) GetSchemaVersion() string {
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
func (_m *OpenControl) GetResources(_a0 string, _a1 string, _a2 *base.Worker) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *base.Worker) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
