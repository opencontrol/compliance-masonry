package docx

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

// readText reads text from a word document
func readText(files []*zip.File) (text string, err error) {
	var documentFile *zip.File
	documentFile, err = retrieveWordDoc(files)
	if err != nil {
		return text, err
	}
	var documentReader io.ReadCloser
	documentReader, err = documentFile.Open()
	if err != nil {
		return text, err
	}

	text, err = wordDocToString(documentReader)
	return
}

// wordDocToString converts a word document to string
func wordDocToString(reader io.Reader) (string, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// retrieveWordDoc fetches a word document.
func retrieveWordDoc(files []*zip.File) (file *zip.File, err error) {
	for _, f := range files {
		if f.Name == "word/document.xml" {
			file = f
		}
	}
	if file == nil {
		err = errors.New("document.xml file not found")
	}
	return
}

func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// normalize fixes quotation marks in documnet
func normalizeQuotes(in rune) rune {
	switch in {
	case '“', '”':
		return '"'
	case '‘', '’':
		return '\''
	}
	return in
}

// cleans template tagged text of all brakets
func normalizeAll(text string) string {
	brakets := regexp.MustCompile("<.*?>")
	quotes := regexp.MustCompile("&quot;")
	text = brakets.ReplaceAllString(text, "")
	text = quotes.ReplaceAllString(text, "\"")
	return strings.Map(normalizeQuotes, text)
}

func cleanText(text string) string {
	braketFinder := regexp.MustCompile("{{.*?}}")
	return braketFinder.ReplaceAllStringFunc(text, normalizeAll)
}
