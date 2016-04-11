package mocks

import "github.com/opencontrol/compliance-masonry/config/common"
import "github.com/stretchr/testify/mock"

type EntryDownloader struct {
	mock.Mock
}

// DownloadEntry provides a mock function with given fields: _a0, _a1
func (_m *EntryDownloader) DownloadEntry(_a0 common.Entry, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Entry, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
