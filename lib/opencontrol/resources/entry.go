package resources

import (
	"github.com/opencontrol/compliance-masonry/tools/vcs"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

// EntryDownloader is a generic interface for how to download entries.
type EntryDownloader interface {
	DownloadEntry(common.Entry, string) error
}

// NewVCSDownloader is a constructor for downloading entries using VCS methods.
func NewVCSDownloader() EntryDownloader {
	return vcsEntryDownloader{vcs.Manager{}}
}

type vcsEntryDownloader struct {
	manager vcs.RepoManager
}

// DownloadEntry is a implementation for downloading entries using VCS methods.
func (v vcsEntryDownloader) DownloadEntry(entry common.Entry, destination string) error {
	err := v.manager.Clone(entry.URL, entry.Revision, destination)
	if err != nil {
		return err
	}
	return nil
}
