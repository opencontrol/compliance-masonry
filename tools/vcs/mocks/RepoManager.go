package mocks

import "github.com/stretchr/testify/mock"

type RepoManager struct {
	mock.Mock
}

// Clone provides a mock function with given fields: url, revision, dir
func (_m *RepoManager) Clone(url string, revision string, dir string) error {
	ret := _m.Called(url, revision, dir)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(url, revision, dir)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
