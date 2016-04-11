package common

import (

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/tools/vcs/mocks"
"errors"
)

var _ = Describe("Entry", func() {
	Describe("Retrieving the config file", func(){
		table.DescribeTable("GetConfigFile", func(e Entry, expectedPath string) {
			assert.Equal(GinkgoT(), e.GetConfigFile(), expectedPath)
		},
			table.Entry("Empty / new base struct to return default", Entry{}, constants.DefaultConfigYaml),
			table.Entry("overriden config file path", Entry{Path: "samplepath"}, "samplepath"),
		)
	})
	Describe("Constructing a new VCEntrySDownloader", func() {
		It("should return a downloader of type VCSEntryDownloader", func() {
			downloader := NewVCSDownloader()
			assert.IsType(GinkgoT(), vcsEntryDownloader{}, downloader)
		})
	})
	Describe("Downloading Entry from VCS", func(){
		table.DescribeTable("DownloadEntry", func(e Entry, err error) {
			m := new(mocks.RepoManager)
			m.On("Clone", e.URL, e.Revision, ".").Return(err)
			v := vcsEntryDownloader{m}
			v.DownloadEntry(e, ".")
			m.AssertExpectations(GinkgoT())
		},
			table.Entry("No error returned", Entry{}, nil),
			table.Entry("An error returned", Entry{}, errors.New("an error")),
		)
	})
})
