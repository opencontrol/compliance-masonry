package common

import (
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfigFile(t *testing.T) {
	var ConfigFileTests = []struct {
		entry        Entry
		expectedFile string
	}{
		{
			entry:        Entry{},
			expectedFile: constants.DefaultConfigYaml,
		},
		{
			entry:        Entry{Path: "samplepath"},
			expectedFile: "samplepath",
		},
	}
	for _, test := range ConfigFileTests {
		assert.Equal(t, test.expectedFile, test.entry.GetConfigFile())
	}
}
