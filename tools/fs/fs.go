/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package fs

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/opencontrol/compliance-masonry/internal/constants"
)

//go:generate mockery -name Util

// Util is an interface for helper file system utilities.
type Util interface {
	OpenAndReadFile(file string) ([]byte, error)
	CopyAll(source string, destination string) error
	Copy(source string, destination string) error
	TempDir(dir string, prefix string) (string, error)
	Mkdirs(dir string) error
	AppendOrCreate(filePath string, text string) error
}

// OSUtil is the struct for dealing with File System Operations on the disk.
type OSUtil struct {
}

// OpenAndReadFile is a util that will check if the file exists, open and then read the file.
func (fs OSUtil) OpenAndReadFile(file string) ([]byte, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, fmt.Errorf("error: %s does not exist", file)
	}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// CopyAll copies recursively from source to destination
func (fs OSUtil) CopyAll(source string, destination string) error {
	return CopyAll(source, destination, "")
}

// Copy copies one file from source to destination
func (fs OSUtil) Copy(source string, destination string) error {
	log.Printf("source %s dest %s\n", source, destination)
	return CopyFile(source, destination)
}

// TempDir creates a temp directory that the user is responsible for cleaning up
func (fs OSUtil) TempDir(dir string, prefix string) (string, error) {
	return TempDir(dir, prefix)
}

// Mkdirs ensures that the directory is created.
func (fs OSUtil) Mkdirs(dir string) error {
	return EnsureDirExists(dir)
}

// AppendOrCreate adds text to file if it exists otherwise it creates a new
// file with the given text
func (fs OSUtil) AppendOrCreate(filePath string, text string) error {
	var err error
	if _, err = os.Stat(filePath); err == nil {
		err = AppendToFile(filePath, text)
	} else {
		err = ioutil.WriteFile(filePath, []byte(text), constants.FileReadWrite)
	}
	return err
}

// AppendToFile adds text to a file
func AppendToFile(filePath string, text string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, constants.FileReadWrite)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.WriteString(text); err != nil {
		return err
	}
	return nil
}

// CopyFile copys a file from source to destination
func CopyFile(source string, destination string) error {
	src, err := os.Open(source)
	if err != nil {
		return err
	}
	defer src.Close()

	dest, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dest.Close()
	_, err = io.Copy(dest, src)
	return err
}

// EnsureDirExists makes sure that the directory exists
func EnsureDirExists(dirpath string) (err error) {
	_, err = os.Stat(dirpath)

	if os.IsNotExist(err) && err != nil {
		if err = EnsureDirExists(filepath.Dir(dirpath)); err == nil {
			err = os.Mkdir(dirpath, constants.DirReadWriteExec)
		}
	}

	return err
}

// TempDir creates a temporary directory
func TempDir(dir string, prefix string) (string, error) {
	return ioutil.TempDir(dir, prefix)
}

// CopyAll copies all files and directories inside `srcDirPath` to `dstDirPath`.
// Based on the CopyAll function from metaleap/go-utils
func CopyAll(srcDirPath, dstDirPath string, skipFileSuffix string) (err error) {
	var (
		srcPath, destPath string
		fileInfo          []os.FileInfo
	)
	if fileInfo, err = ioutil.ReadDir(srcDirPath); err == nil {
		EnsureDirExists(dstDirPath)
		for _, fi := range fileInfo {
			if srcPath, destPath = filepath.Join(srcDirPath, fi.Name()), filepath.Join(dstDirPath, fi.Name()); fi.IsDir() {
				if skipFileSuffix == "" || !strings.HasSuffix(srcPath, skipFileSuffix) {
					CopyAll(srcPath, destPath, skipFileSuffix)
				}
			} else {
				CopyFile(srcPath, destPath)
			}
		}
	}
	return err
}
