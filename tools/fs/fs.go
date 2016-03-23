package fs

import (
	"fmt"
	"io/ioutil"
	"os"
)

// OpenAndReadFile is a util that will check if the file exists, open and then read the file.
func OpenAndReadFile(file string) ([]byte, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, fmt.Errorf("Error: %s does not exist\n", file)
	}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
