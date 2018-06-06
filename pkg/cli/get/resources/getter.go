package resources

import (
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"github.com/opencontrol/compliance-masonry/pkg/lib/opencontrol"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/tools/fs"
	"github.com/opencontrol/compliance-masonry/tools/mapset"
	"log"
	"os"
	"path/filepath"
)

// getAllLocalResources will get try to get the resources that are in the current "source" directory and place them
// in the final "destination" workspace directory.
func getAllLocalResources(source string, destination string, opencontrol common.OpenControl, getter Getter) error {
	// Get Certifications
	log.Println("Retrieving certifications")
	err := getter.GetLocalResources(source, opencontrol.GetCertifications(), destination,
		constants.DefaultCertificationsFolder, false, constants.Certifications)
	if err != nil {
		return err
	}
	// Get Standards
	log.Println("Retrieving standards")
	err = getter.GetLocalResources(source, opencontrol.GetStandards(), destination,
		constants.DefaultStandardsFolder, false, constants.Standards)
	if err != nil {
		return err
	}
	// Get Components
	log.Println("Retrieving components")
	err = getter.GetLocalResources(source, opencontrol.GetComponents(), destination,
		constants.DefaultComponentsFolder, true, constants.Components)
	if err != nil {
		return err
	}
	return nil
}

// getAllRemoteResources will get try to get the dependencies from their respective repositories and put them
// in the final "destination" workspace directory.
func getAllRemoteResources(destination string, opencontrol common.OpenControl, getter Getter) error {
	// Get Certifications
	log.Println("Retrieving dependent certifications")
	err := getter.GetRemoteResources(destination, constants.DefaultCertificationsFolder,
		opencontrol.GetCertificationsDependencies())
	if err != nil {
		return err
	}
	// Get Standards
	log.Println("Retrieving dependent standards")
	err = getter.GetRemoteResources(destination, constants.DefaultStandardsFolder,
		opencontrol.GetStandardsDependencies())
	if err != nil {
		return err
	}
	// Get Components
	log.Println("Retrieving dependent components")
	err = getter.GetRemoteResources(destination, constants.DefaultComponentsFolder,
		opencontrol.GetComponentsDependencies())
	if err != nil {
		return err
	}
	return nil
}

// GetResources will download all the resources that are specified by the schema first by copying the
// local resources then downloading the remote ones and letting their respective schema version handle
// how to get their resources.
func GetResources(source string, destination string, opencontrol common.OpenControl, getter Getter) error {
	// Local
	err := getAllLocalResources(source, destination, opencontrol, getter)
	if err != nil {
		return err
	}

	// Remote
	err = getAllRemoteResources(destination, opencontrol, getter)
	if err != nil {
		return err
	}
	return nil
}

// Getter is an interface for how to get and place local and remote resources.
type Getter interface {
	GetLocalResources(source string, resources []string, destination string, subfolder string, recursively bool,
		resourceType constants.ResourceType) error
	GetRemoteResources(destination string, subfolder string, entries []common.RemoteSource) error
}

// NewVCSAndLocalGetter constructs a new resource getter with the type of parser to use for the files.
func NewVCSAndLocalGetter(parser opencontrol.SchemaParser) Getter {
	return &vcsAndLocalFSGetter{Downloader: NewVCSDownloader(), FSUtil: fs.OSUtil{}, Parser: parser,
		ResourceMap: mapset.Init()}
}

// vcsAndLocalFSGetter is the resource getter that uses VCS for remote resource getting and local file system
// for local resources.
type vcsAndLocalFSGetter struct {
	Downloader  Downloader
	FSUtil      fs.Util
	ResourceMap mapset.MapSet
	Parser      opencontrol.SchemaParser
}

// reserveLocalResourceDestination will attempt to make a unique reservation for a particular type of resource and make
// any necessary filesystem changes so that the resource can moved there.
func (g *vcsAndLocalFSGetter) reserveLocalResourceDestination(resourceType constants.ResourceType, subfolder string,
	destination string, resource string) (string, error) {
	// Attempt to make a unique reservation for the resource.
	if result := g.ResourceMap.Reserve(string(resourceType), resource); !result.Success {
		return "", result.Error
	}
	// Construct the folder of where the resource should be placed.
	resourceDestinationFolder := filepath.Join(destination, subfolder)
	// Construct the final path for the resource itself once placed in the destination path.
	resourceDestination := filepath.Join(resourceDestinationFolder, filepath.Base(resource))

	log.Printf("Ensuring directory %s exists\n", resourceDestinationFolder)
	err := g.FSUtil.Mkdirs(resourceDestinationFolder)
	if err != nil {
		return "", err
	}
	return resourceDestination, nil
}

func (g *vcsAndLocalFSGetter) copyLocalResource(resourceSource string,
	resourceDestination string, recursively bool) error {
	var err error
	log.Printf("Attempting to copy local resource %s into %s\n", resourceSource, resourceDestination)
	if recursively {
		log.Printf("Copying local resource %s recursively into %s\n",
			resourceSource, resourceDestination)
		err = g.FSUtil.CopyAll(resourceSource, resourceDestination)
	} else {
		log.Printf("Copying local resource %s into %s\n", resourceSource, resourceDestination)
		err = g.FSUtil.Copy(resourceSource, resourceDestination)
	}
	if err != nil {
		log.Printf("Copying local resources %s failed\n", resourceSource)
		return err
	}
	return nil
}

// GetLocalResources is the implementation that uses the local file system to get local resources.
func (g *vcsAndLocalFSGetter) GetLocalResources(source string, resources []string, destination string,
	subfolder string, recursively bool, resourceType constants.ResourceType) error {
	for _, resource := range resources {
		// Attempt to reserve a space for the local resource.
		resourceDestination, err := g.reserveLocalResourceDestination(resourceType,
			subfolder, destination, resource)
		if err != nil {
			return err
		}

		// Find the final path of where the resource is originally located.
		resourceSource := filepath.Join(source, resource)

		// Attempt to copy the resource.
		err = g.copyLocalResource(resourceSource, resourceDestination, recursively)
		if err != nil {
			return err
		}

	}
	return nil
}

// GetRemoteResources is the implementation that uses VCS to get remote resources.
func (g *vcsAndLocalFSGetter) GetRemoteResources(destination string, subfolder string,
	entries []common.RemoteSource) error {
	// Create the temporary directory for where to clone all the remote resources.
	tempResourcesDir, err := g.FSUtil.TempDir("", "opencontrol-resources")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempResourcesDir)
	for _, entry := range entries {
		// Create the final path for where to clone.
		tempPath := filepath.Join(tempResourcesDir, subfolder, filepath.Base(entry.GetURL()))

		// Clone repo
		log.Printf("Attempting to clone %v into %s\n", entry, tempPath)

		err := g.Downloader.DownloadRepo(entry, tempPath)
		if err != nil {
			return err
		}

		// If contextdir is defined, switch to that dir for content
		if entry.GetContextDir() != "" {
			tempPath = filepath.Join(tempPath, entry.GetContextDir())
		}

		// Parse the opencontrol.yaml.
		configBytes, err := g.FSUtil.OpenAndReadFile(filepath.Join(tempPath, entry.GetConfigFile()))
		if err != nil {
			return err
		}
		opencontrol, err := g.Parser.Parse(configBytes)
		if err != nil {
			return err
		}

		// Get the resources specified in the OpenControl YAML
		err = GetResources(tempPath, destination, opencontrol, g)
		if err != nil {
			return err
		}
	}
	return nil
}
