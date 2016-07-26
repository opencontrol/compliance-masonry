package resources

import (

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
	"github.com/opencontrol/compliance-masonry/tools/vcs/mocks"
	"errors"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

var _ = Describe("Downloader", func() {

	Describe("Constructing a new VCEntrySDownloader", func() {
		It("should return a downloader of type VCSEntryDownloader", func() {
			downloader := NewVCSDownloader()
			assert.IsType(GinkgoT(), vcsEntryDownloader{}, downloader)
		})
	})
	Describe("Downloading Entry from VCS", func(){
		table.DescribeTable("DownloadEntry", func(e common.Entry, err error) {
			m := new(mocks.RepoManager)
			m.On("Clone", e.URL, e.Revision, ".").Return(err)
			v := vcsEntryDownloader{m}
			v.DownloadEntry(e, ".")
			m.AssertExpectations(GinkgoT())
		},
			table.Entry("No error returned", common.Entry{}, nil),
			table.Entry("An error returned", common.Entry{}, errors.New("an error")),
		)
	})
})
