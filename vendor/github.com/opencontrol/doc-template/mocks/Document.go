package mocks

import "github.com/stretchr/testify/mock"

type Document struct {
	mock.Mock
}

// ReadFile provides a mock function with given fields: _a0
func (_m *Document) ReadFile(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateConent provides a mock function with given fields: _a0
func (_m *Document) UpdateConent(_a0 string) {
	_m.Called(_a0)
}

// GetContent provides a mock function with given fields:
func (_m *Document) GetContent() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WriteToFile provides a mock function with given fields: _a0, _a1
func (_m *Document) WriteToFile(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Document) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
