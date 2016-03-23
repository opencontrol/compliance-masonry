package vcs

import (
	"fmt"
	"log"

	"github.com/Masterminds/vcs"
)

// RepoManager is the interface for how to do jobs with VCS
type RepoManager interface {
	Clone(url string, revision string, dir string) error
}

const (
	repoInitFailed     = "Repo initialization failed"
	repoCloneFailed    = "Cloning repo failed"
	repoCheckoutFailed = "Revision Checkout failed"
	errorContainer     = "[Error: %s Repo: %s Revision: %s Dir: %s Error Details: %s]\n"
)

// Manager is the concrete implementation of RepoManager
type Manager struct{}

// Clone will clone the repo to a specified location and then checkout the repo at the particular revision.
func (m Manager) Clone(url string, revision string, dir string) error {
	log.Printf("Initializing repo %s into %s\n", url, dir)
	repo, err := vcs.NewRepo(url, dir)
	if err != nil {
		return fmt.Errorf(errorContainer, repoInitFailed, url, revision, dir, err.Error())
	}
	log.Printf("Cloning %s into %s\n", url, dir)
	err = repo.Get()
	if err != nil {
		return fmt.Errorf(errorContainer, repoCloneFailed, url, revision, dir, err.Error())
	}
	if revision != "" {
		log.Printf("Checking out revision %s for repo %s\n", revision, url)
		err = repo.UpdateVersion(revision)
		if err != nil {
			return fmt.Errorf(errorContainer, repoCheckoutFailed, url, revision, dir, err.Error())
		}
	} else {
		log.Printf("Assuming default revision for repo %s\n", url)
	}
	return nil
}
