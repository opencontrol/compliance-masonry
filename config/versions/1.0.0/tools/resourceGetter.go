package tools

import (
	"github.com/go-utils/ufs"
	"github.com/opencontrol/compliance-masonry-go/config"
	"github.com/opencontrol/compliance-masonry-go/config/common"
	"github.com/opencontrol/compliance-masonry-go/tools/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ResourceGetter is an interface for how to get and place local and remote resources.
type ResourceGetter interface {
	GetLocalResources(resources []string, destination string, subfolder string, recursively bool) error
	GetRemoteResources(destination string, subfolder string, worker *common.ConfigWorker, entries []common.Entry) error
}

// VCSAndLocalFSGetter is the resource getter that uses VCS for remote resource getting and local file system for local resources.
type VCSAndLocalFSGetter struct{}

// GetLocalResources is the implementation that uses the local file system to get local resources.
func (g VCSAndLocalFSGetter) GetLocalResources(resources []string, destination string, subfolder string, recursively bool) error {
	for _, resource := range resources {
		var err error
		if recursively {
			err = ufs.CopyAll(resource,
				filepath.Join(destination, subfolder, filepath.Base(resource)), nil)
		} else {
			err = ufs.CopyFile(resource,
				filepath.Join(destination, subfolder, filepath.Base(resource)))
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// GetRemoteResources is the implementation that uses VCS to get remote resources.
func (g VCSAndLocalFSGetter) GetRemoteResources(destination string, subfolder string, worker *common.ConfigWorker, entries []common.Entry) error {
	tempResourcesDir, err := ioutil.TempDir("", "opencontrol-resources")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempResourcesDir)
	for _, entry := range entries {
		tempPath := filepath.Join(tempResourcesDir, subfolder, filepath.Base(entry.URL))
		// Clone repo
		log.Printf("Attempting to clone %v into %s\n", entry, tempPath)
		err := worker.Downloader.DownloadEntry(entry, tempPath)
		if err != nil {
			return err
		}
		// Parse
		configBytes, err := fs.OpenAndReadFile(filepath.Join(tempPath, entry.GetConfigFile()))
		if err != nil {
			return err
		}
		schema, err := config.Parse(worker.Parser, configBytes)
		if err != nil {
			return err
		}
		err = schema.GetResources(destination, worker)
		if err != nil {
			return err
		}
	}
	return nil
}
