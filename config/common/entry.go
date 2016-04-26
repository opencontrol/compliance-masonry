package common

import (
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/tools/vcs"
)

// Entry is a generic holder for handling the specific location and revision of a resource.
type Entry struct {
	URL      string `yaml:"url"`
	Revision string `yaml:"revision"`
	Path     string `yaml:"path"`
}

// GetConfigFile is a getter for the config file name. Will return DefaultConfigYaml value if none has been set.
func (e Entry) GetConfigFile() string {
	if e.Path == "" {
		return constants.DefaultConfigYaml
	}
	return e.Path
}

// EntryDownloader is a generic interface for how to download entries.
type EntryDownloader interface {
	DownloadEntry(Entry, string) error
}

// NewVCSDownloader is a constructor for downloading entries using VCS methods.
func NewVCSDownloader() EntryDownloader {
	return vcsEntryDownloader{vcs.Manager{}}
}

type vcsEntryDownloader struct {
	manager vcs.RepoManager
}

// DownloadEntry is a implementation for downloading entries using VCS methods.
func (v vcsEntryDownloader) DownloadEntry(entry Entry, destination string) error {
	err := v.manager.Clone(entry.URL, entry.Revision, destination)
	if err != nil {
		return err
	}
	return nil
}
