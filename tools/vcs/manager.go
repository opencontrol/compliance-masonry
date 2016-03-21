package vcs

import (
	"fmt"
	"github.com/Masterminds/vcs"
	"log"
)

const (
	repoInitFailed     = "Repo initialization failed"
	repoCloneFailed    = "Cloning repo failed"
	repoCheckoutFailed = "Revision Checkout failed"
	errorContainer     = "[Error: %s Repo: %s Revision: %s Dir: %s Error Details: %s]\n"
)

func Clone(url string, revision string, dir string) error {
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
