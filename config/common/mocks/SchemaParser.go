package mocks

import "github.com/opencontrol/compliance-masonry/config/common"
import "github.com/stretchr/testify/mock"

type SchemaParser struct {
	mock.Mock
}

// ParseV1_0_0 provides a mock function with given fields: data
func (_m *SchemaParser) ParseV1_0_0(data []byte) (common.BaseSchema, error) {
	ret := _m.Called(data)

	var r0 common.BaseSchema
	if rf, ok := ret.Get(0).(func([]byte) common.BaseSchema); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(common.BaseSchema)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
