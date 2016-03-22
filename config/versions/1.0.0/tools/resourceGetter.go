package tools

import (
	"github.com/opencontrol/compliance-masonry-go/config/common"
	"github.com/go-utils/ufs"
	"path/filepath"
	"io/ioutil"
	"os"
	"log"
	"github.com/opencontrol/compliance-masonry-go/tools/fs"
	"github.com/opencontrol/compliance-masonry-go/config"
)

type ResourceGetter interface {
	GetLocalResources(resources []string, destination string, subfolder string, recursively bool) error
	GetRemoteResources(destination string, subfolder string, worker *common.ConfigWorker, entries []common.Entry) error
}


type VCSAndLocalFSGetter struct {}

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

func (g VCSAndLocalFSGetter) GetRemoteResources(destination string, subfolder string, worker *common.ConfigWorker, entries []common.Entry) error {
	tempResourcesDir, err := ioutil.TempDir("", "opencontrol-resources")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempResourcesDir)
	for _, entry := range entries{
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
