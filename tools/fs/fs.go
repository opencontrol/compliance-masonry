package fs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-utils/ufs"
)

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
		return nil, fmt.Errorf("Error: %s does not exist\n", file)
	}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// CopyAll copies recursively from source to destination
func (fs OSUtil) CopyAll(source string, destination string) error {
	return ufs.CopyAll(source, destination, nil)
}

// Copy copies one file from source to destination
func (fs OSUtil) Copy(source string, destination string) error {
	log.Printf("source %s dest %s\n", source, destination)
	return ufs.CopyFile(source, destination)
}

// TempDir creates a temp directory that the user is responsible for cleaning up
func (fs OSUtil) TempDir(dir string, prefix string) (string, error) {
	return ioutil.TempDir(dir, prefix)
}

// Mkdirs ensures that the directory is created.
func (fs OSUtil) Mkdirs(dir string) error {
	return ufs.EnsureDirExists(dir)
}

// AppendOrCreate adds text to file if it exists otherwise it creates a new
// file with the given text
func (fs OSUtil) AppendOrCreate(filePath string, text string) error {
	var err error
	if _, err := os.Stat(filePath); err == nil {
		err = AppendToFile(filePath, text)
	} else {
		err = ioutil.WriteFile(filePath, []byte(text), 0700)
	}
	return err
}

// AppendToFile adds text to a file
func AppendToFile(filePath string, text string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0700)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err = file.WriteString(text); err != nil {
		return err
	}
	return nil
}
