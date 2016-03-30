package common

import (
	"errors"
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
	"github.com/opencontrol/compliance-masonry-go/tools/vcs/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfigFile(t *testing.T) {
	var ConfigFileTests = []struct {
		entry        Entry
		expectedFile string
	}{
		{
			// Return the default
			entry:        Entry{},
			expectedFile: constants.DefaultConfigYaml,
		},
		{
			// Return a custom path.
			entry:        Entry{Path: "samplepath"},
			expectedFile: "samplepath",
		},
	}
	for _, test := range ConfigFileTests {
		assert.Equal(t, test.expectedFile, test.entry.GetConfigFile())
	}
}

func TestNewVCSDownloader(t *testing.T) {
	downloader := NewVCSDownloader()
	assert.IsType(t, vcsEntryDownloader{}, downloader)

}

func TestVCSDownloadEntry(t *testing.T) {
	var DownloadEntryTests = []struct {
		entry Entry
		err   error
	}{
		{
			// No error returned
			entry: Entry{URL: "link", Revision: "master"},
			err:   nil,
		},
		{
			// Error returned.
			entry: Entry{URL: "link", Revision: "master"},
			err:   errors.New("an error"),
		},
	}
	for _, test := range DownloadEntryTests {
		m := new(mocks.RepoManager)
		m.On("Clone", test.entry.URL, test.entry.Revision, ".").Return(test.err)
		v := vcsEntryDownloader{m}
		v.DownloadEntry(test.entry, ".")
		m.AssertExpectations(t)
	}
}
