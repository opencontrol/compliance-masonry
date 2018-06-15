/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package vcs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Masterminds/vcs"
)

//go:generate mockery -name RepoManager

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

	files := GetVCSFolderContents(dir)
	if len(files) == 1 {
		log.Printf("Cloning %s into %s\n", url, dir)
		err = repo.Get()
		if err != nil {
			return fmt.Errorf(errorContainer, repoCloneFailed, url, revision, dir, err.Error())
		}
	} else {
		log.Printf("Repository already exists. Skipping....")
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

// GetVCSFolderContents determines if there are actual files in the VCS dir
func GetVCSFolderContents(directory string) []string {
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}
