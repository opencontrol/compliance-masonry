package mocks

import "github.com/stretchr/testify/mock"

type Util struct {
	mock.Mock
}

// OpenAndReadFile provides a mock function with given fields: file
func (_m *Util) OpenAndReadFile(file string) ([]byte, error) {
	ret := _m.Called(file)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyAll provides a mock function with given fields: source, destination
func (_m *Util) CopyAll(source string, destination string) error {
	ret := _m.Called(source, destination)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(source, destination)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Copy provides a mock function with given fields: source, destination
func (_m *Util) Copy(source string, destination string) error {
	ret := _m.Called(source, destination)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(source, destination)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TempDir provides a mock function with given fields: dir, prefix
func (_m *Util) TempDir(dir string, prefix string) (string, error) {
	ret := _m.Called(dir, prefix)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(dir, prefix)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(dir, prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Mkdirs provides a mock function with given fields: dir
func (_m *Util) Mkdirs(dir string) error {
	ret := _m.Called(dir)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(dir)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppendOrCreate provides a mock function with given fields: filePath, text
func (_m *Util) AppendOrCreate(filePath string, text string) error {
	ret := _m.Called(filePath, text)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(filePath, text)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
