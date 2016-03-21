package common

import "github.com/opencontrol/compliance-masonry-go/tools/vcs"

// Entry is a generic holder for handling the specific location and revision of a resource.
type Entry struct {
	URL      string `yaml:"url"`
	Revision string `yaml:"revision"`
	Path     string `yaml:path"`
}

type EntryDownloader interface {
	DownloadEntry(Entry, string) error
}

type VCSEntryDownloader struct {
}

func (v VCSEntryDownloader) DownloadEntry(entry Entry, destination string) error {
	err := vcs.Clone(entry.URL, entry.Revision, destination)
	if err != nil {
		return err
	}
	return nil
}
