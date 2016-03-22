package helpers

//TestCreateDirectory

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Test that the CreateDirectory method works
func TestCreateDirectory(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(tempDir)
	newDir := filepath.Join(tempDir, "testtesttes")
	// Check that the dir doesn't exist
	if _, err := os.Stat(newDir); err == nil {
		t.Errorf("Expected a directory to not exist")
	}
	CreateDirectory(newDir)
	// Check that the dir exists
	if _, err := os.Stat(newDir); err != nil {
		if os.IsNotExist(err) {
			t.Errorf("Expected a directory to exist")
		}
	}

}

// Test that the AppendToFile method can append to a file
func TestAppendToFile(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(tempDir)
	text := "test text"
	filePath := filepath.Join(tempDir, "test.txt")
	// Write a file with some text
	ioutil.WriteFile(filePath, []byte(text), 0700)
	newText := "new text"
	AppendToFile(filePath, newText)
	fileText, _ := ioutil.ReadFile(filePath)
	// Check if text was appended
	if string(fileText) != text+newText {
		t.Errorf("Expected text to be appended")
	}
}

// Test that the AppendOrCreate method can append to a file
func TestAppendOrCreateAppend(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(tempDir)
	text := "test text"
	filePath := filepath.Join(tempDir, "test.txt")
	// Write a file with some text
	ioutil.WriteFile(filePath, []byte(text), 0700)
	newText := "new text"
	AppendOrCreate(filePath, newText)
	fileText, _ := ioutil.ReadFile(filePath)
	// Check if text was appended
	if string(fileText) != text+newText {
		t.Errorf("Expected text to be appended")
	}
}

// Test that the AppendOrCreate method can create a file
func TestAppendOrCreateCreate(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(tempDir)
	filePath := filepath.Join(tempDir, "test.txt")
	// Check that the dir doesn't exist
	if _, err := os.Stat(filePath); err == nil {
		t.Errorf("Expected a directory to not exist")
	}
	AppendOrCreate(filePath, "test")
	// Check that the dir exists
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			t.Errorf("Expected a directory to exist")
		}
	}
}

// Check that the CopyFile methods works
func TestCopyFile(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(tempDir)
	filePath := filepath.Join(tempDir, "test.txt")
	text := "test 1 2 3"
	AppendOrCreate(filePath, text)
	newFilePath := filepath.Join(tempDir, "test_copy.txt")
	CopyFile(filePath, newFilePath)
	fileText, _ := ioutil.ReadFile(filePath)
	// Check if file was copied
	if string(fileText) != text {
		t.Errorf("Expected text to be copied")
	}
}

// Validate that the copy dir function works
func TestCopyDir(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(tempDir)
	newTempDir, _ := ioutil.TempDir("", "copy")
	defer os.RemoveAll(tempDir)
	filePath := filepath.Join(tempDir, "test.txt")
	text := "test 1 2 3"
	AppendOrCreate(filePath, text)
	newFilePath := filepath.Join(newTempDir, "test.txt")
	CopyDir(tempDir, newTempDir)
	fileText, _ := ioutil.ReadFile(newFilePath)
	// Check if directory was copied
	if string(fileText) != text {
		t.Errorf("Expected text to be copied")
	}
}

// Validate the sub-directories are copied with the CopyDir method
func TestCopyDirRecursive(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "example")
	defer os.RemoveAll(tempDir)
	newTempDir, _ := ioutil.TempDir("", "copy")
	defer os.RemoveAll(tempDir)
	newDir := filepath.Join(tempDir, "testdir")
	CreateDirectory(newDir)
	filePath := filepath.Join(newDir, "test.txt")
	text := "test 1 2 3"
	AppendOrCreate(filePath, text)
	newFilePath := filepath.Join(newTempDir, "testdir", "test.txt")
	CopyDir(tempDir, newTempDir)
	fileText, _ := ioutil.ReadFile(newFilePath)
	// Check if directory was copied
	if string(fileText) != text {
		t.Errorf("Expected text to be copied")
	}
}
