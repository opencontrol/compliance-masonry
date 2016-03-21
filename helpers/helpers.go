// Package helpers contains shared functions that assist other packages
package helpers

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// CreateDirectory creates a directory
func CreateDirectory(directory string) string {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.MkdirAll(directory, 0700)
	}
	return directory
}

// AppendToFile adds text to a file
func AppendToFile(filepath string, text string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, 0700)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err = file.WriteString(text); err != nil {
		panic(err)
	}
}

// AppendOrCreate adds text to file if it exists otherwise it creates a new
// file with the given text
func AppendOrCreate(filepath string, text string) {
	if _, err := os.Stat(filepath); err == nil {
		AppendToFile(filepath, text)
	} else {
		ioutil.WriteFile(filepath, []byte(text), 0700)
	}
}

// CopyFile function from https://www.socketloop.com/tutorials/golang-copy-directory-including-sub-directories-files
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourcefile.Close()
	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destfile.Close()
	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}
	}
	return
}

// CopyDir function from https://www.socketloop.com/tutorials/golang-copy-directory-including-sub-directories-files
func CopyDir(source string, dest string) error {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}
	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)
	for _, obj := range objects {
		sourcefilepointer := source + "/" + obj.Name()
		destinationfilepointer := dest + "/" + obj.Name()
		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
