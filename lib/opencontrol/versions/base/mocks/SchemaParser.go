package mocks


import (
	"github.com/stretchr/testify/mock"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol/versions/base"
)

type SchemaParser struct {
	mock.Mock
}

// ParseV1_0_0 provides a mock function with given fields: data
func (_m *SchemaParser) ParseV1_0_0(data []byte) (base.OpenControl, error) {
	ret := _m.Called(data)

	var r0 base.OpenControl
	if rf, ok := ret.Get(0).(func([]byte) base.OpenControl); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(base.OpenControl)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
