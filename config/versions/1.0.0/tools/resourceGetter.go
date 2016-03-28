package tools

import (
	"github.com/go-utils/ufs"
	"github.com/opencontrol/compliance-masonry-go/config"
	"github.com/opencontrol/compliance-masonry-go/config/common"
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
	"github.com/opencontrol/compliance-masonry-go/tools/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ResourceGetter is an interface for how to get and place local and remote resources.
type ResourceGetter interface {
	GetLocalResources(source string, resources []string, destination string, subfolder string, recursively bool, worker *common.ConfigWorker, resourceType constants.ResourceType) error
	GetRemoteResources(destination string, subfolder string, worker *common.ConfigWorker, entries []common.Entry) error
}

// VCSAndLocalFSGetter is the resource getter that uses VCS for remote resource getting and local file system for local resources.
type VCSAndLocalFSGetter struct{}

// GetLocalResources is the implementation that uses the local file system to get local resources.
func (g VCSAndLocalFSGetter) GetLocalResources(source string, resources []string, destination string, subfolder string, recursively bool, worker *common.ConfigWorker, resourceType constants.ResourceType) error {
	for _, resource := range resources {
		if result := worker.ResourceMap.Reserve(string(resourceType), resource); !result.Success {
			return result.Error
		}
		resourceSource := filepath.Join(source, resource)
		resourceDestination := filepath.Join(destination, subfolder, filepath.Base(resource))
		log.Printf("Attempting to copy local resource %s into %s\n", resourceSource, resourceDestination)
		var err error
		if recursively {
			log.Printf("Copying local resource %s reursively into %s\n", resourceSource, resourceDestination)
			err = ufs.CopyAll(resourceSource, resourceDestination, nil)
		} else {
			log.Printf("Copying local resource %s reursively into %s\n", resourceSource, resourceDestination)
			err = ufs.CopyFile(resourceSource, resourceDestination)
		}
		if err != nil {
			log.Printf("Copying local resources %s failed\n", resourceSource)
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
		err = schema.GetResources(tempPath, destination, worker)
		if err != nil {
			return err
		}
	}
	return nil
}
