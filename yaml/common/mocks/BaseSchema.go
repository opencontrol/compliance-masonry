package mocks

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
