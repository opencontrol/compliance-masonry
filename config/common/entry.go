package common

import (
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
	"github.com/opencontrol/compliance-masonry-go/tools/vcs"
)

// Entry is a generic holder for handling the specific location and revision of a resource.
type Entry struct {
	URL      string `yaml:"url"`
	Revision string `yaml:"revision"`
	Path     string `yaml:path"`
}

func (e Entry) GetConfigFile() string {
	if e.Path == "" {
		return constants.DefaultConfigYaml
	}
	return e.Path
}

type EntryDownloader interface {
	DownloadEntry(Entry, string) error
}

func NewVCSDownloader() EntryDownloader {
	return vcsEntryDownloader{vcs.Manager{}}
}
type vcsEntryDownloader struct {
	manager vcs.VCSManager
}

func (v vcsEntryDownloader) DownloadEntry(entry Entry, destination string) error {
	err := v.manager.Clone(entry.URL, entry.Revision, destination)
	if err != nil {
		return err
	}
	return nil
}
