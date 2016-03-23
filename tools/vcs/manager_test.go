package vcs

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestClone(t *testing.T) {
	var CloneTests = []struct {
		url         string
		rev         string
		errorIsNull bool
		errorString string
	}{
		{
			// Sane check.
			url:         "https://github.com/opencontrol/compliance-masonry-go",
			rev:         "master",
			errorIsNull: true,
		},
		{
			// Sane check. No revision
			url:         "https://github.com/opencontrol/compliance-masonry-go",
			errorIsNull: true,
		},
		{
			// Can't init / detect the repo.
			url:         "https://myrepo/opencontrol/compliance-masonry-go",
			rev:         "master",
			errorIsNull: false,
			errorString: repoInitFailed,
		},
		{
			// Can't init / detect the repo.
			url:         "http://user:name@github.com/opencontrol/compliance-masonry-go-blah",
			rev:         "master",
			errorIsNull: false,
			errorString: repoCloneFailed,
		},
		{
			// Get a revision that doesn't exist.
			url:         "https://github.com/opencontrol/compliance-masonry-go",
			rev:         "master-ultimate-branch-that-never-exists",
			errorIsNull: false,
			errorString: repoCheckoutFailed,
		},
	}
	for _, test := range CloneTests {
		local, err := ioutil.TempDir("", "go-vcs")
		if err != nil {
			assert.Fail(t, err.Error())
		}
		m := Manager{}
		err = m.Clone(test.url, test.rev, local)
		if test.errorIsNull {
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), test.errorString)
		}
		os.RemoveAll(local)
	}
}
