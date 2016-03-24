package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

type buildGitbookTest struct {
	config           gitbookConfig
	expectedMessages []string
}

var buildGitbookTests = []buildGitbookTest{
	//Check that the gitbook is correctly exported given the fixtures
	{
		gitbookConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures/",
			certification:  "LATO",
			markdownPath:   "",
		},
		[]string{"Warning: markdown directory does not exist", "New Gitbook Documentation Created"},
	},
	{
		gitbookConfig{
			opencontrolDir: "",
			certification:  "LATO",
			markdownPath:   "",
		},
		[]string{"Error: `opencontrols/certifications` directory does exist"},
	},
	{
		gitbookConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
			certification:  "LATO",
			markdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
		},
		[]string{"New Gitbook Documentation Created"},
	},
	{
		gitbookConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
			certification:  "LAT",
			markdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
		},
		[]string{
			"`compliance-masonry-go docs gitbook LATO`",
			"Error: `fixtures/opencontrol_fixtures_with_markdown/certifications/LAT.yaml` does not exist\nUse one of the following:",
		},
	},
	{
		gitbookConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
			certification:  "",
			markdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
		},
		[]string{"Error: New Missing Certification Argument", "Usage: masonry-go docs gitbook FedRAMP-low"},
	},
}

func TestMakeGitbook(t *testing.T) {
	for _, example := range buildGitbookTests {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		example.config.exportPath = tempDir
		actualMessages := example.config.makeGitbook()
		for _, actualMessage := range actualMessages {
			if !LookForString(actualMessage, example.expectedMessages) {
				t.Errorf("Could not find `%s` in the expected messages", actualMessage)
			}
		}
		if len(actualMessages) != len(example.expectedMessages) {
			t.Errorf("The expected number of messages is %d, but %d messages were returned",
				len(example.expectedMessages), len(actualMessages),
			)

		}
	}
}

func LookForString(searchString string, stringSlice []string) bool {
	for _, singeString := range stringSlice {
		if strings.Compare(singeString, searchString) == 0 {
			return true
		}
	}
	return false
}
