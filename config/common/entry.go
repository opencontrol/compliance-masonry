package common

import (
	"github.com/opencontrol/compliance-masonry-go/tools/vcs"
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
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

type VCSEntryDownloader struct {
}

func (v VCSEntryDownloader) DownloadEntry(entry Entry, destination string) error {
	err := vcs.Clone(entry.URL, entry.Revision, destination)
	if err != nil {
		return err
	}
	return nil
}
