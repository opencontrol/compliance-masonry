package docx

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type docxTest struct {
	fixture string
	content string
	err     error
}

var readDocTests = []docxTest{
	//  Check that reading a document works
	{fixture: "fixtures/test.docx", content: "This is a test document", err: nil},
}

func TestReadFile(t *testing.T) {
	for _, example := range readDocTests {
		actualDoc := new(Docx)
		actualErr := actualDoc.ReadFile(example.fixture)
		assert.Equal(t, example.err, actualErr)
		if actualErr == nil {
			assert.Contains(t, actualDoc.content, example.content)
		}
	}
}

var writeDocTests = []docxTest{
	//  Check that writing a document works
	{fixture: "fixtures/test.docx", content: "This is an addition", err: nil},
}

func TestWriteToFile(t *testing.T) {
	for _, example := range writeDocTests {
		exportTempDir, _ := ioutil.TempDir("", "exports")
		// Overwrite content
		actualDoc := new(Docx)
		actualDoc.ReadFile(example.fixture)
		currentContent := actualDoc.GetContent()
		actualDoc.UpdateConent(strings.Replace(currentContent, "This is a test document", example.content, -1))
		newFilePath := filepath.Join(exportTempDir, "test.docx")
		actualDoc.WriteToFile(newFilePath, actualDoc.GetContent())
		// Check content
		newActualDoc := new(Docx)
		newActualDoc.ReadFile(newFilePath)
		assert.Contains(t, newActualDoc.GetContent(), example.content)
		os.RemoveAll(exportTempDir)
	}

}
