/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package resources

import (
	"errors"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	commonMocks "github.com/opencontrol/compliance-masonry/pkg/lib/common/mocks"
	"github.com/opencontrol/compliance-masonry/tools/vcs/mocks"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Downloader", func() {
	Describe("Constructing a new VCEntrySDownloader", func() {
		It("should return a downloader of type VCSEntryDownloader", func() {
			downloader := NewVCSDownloader()
			assert.IsType(GinkgoT(), vcsEntryDownloader{}, downloader)
		})
	})
	Describe("Downloading Entry from VCS", func() {
		table.DescribeTable("DownloadRepo", func(err error) {
			remoteSource := new(commonMocks.RemoteSource)
			remoteSource.On("GetURL").Return("https://github.com/opencontrol/notarealrepo")
			remoteSource.On("GetRevision").Return("master")
			m := new(mocks.RepoManager)
			m.On("Clone", remoteSource.GetURL(), remoteSource.GetRevision(), ".").Return(err)
			v := vcsEntryDownloader{m}
			v.DownloadRepo(remoteSource, ".")
			m.AssertExpectations(GinkgoT())
		},
			table.Entry("No error returned", nil),
			table.Entry("An error returned", errors.New("an error")),
		)
	})
})
