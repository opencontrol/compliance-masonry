package vcs

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
)

var _ = Describe("Manager", func() {
	DescribeTable("Clone", func(url, rev string, errorIsNull bool, errorString string) {
		local, err := ioutil.TempDir("", "go-vcs")
		if err != nil {
			assert.Fail(GinkgoT(), err.Error())
		}
		m := Manager{}
		err = m.Clone(url, rev, local)
		if errorIsNull {
			assert.Nil(GinkgoT(), err)
		} else {
			assert.NotNil(GinkgoT(), err)
			assert.Contains(GinkgoT(), err.Error(), errorString)
		}
		os.RemoveAll(local)
	},
		Entry("sane check", "https://github.com/opencontrol/compliance-masonry", "master", true, ""),
		Entry("sane check no revision", "https://github.com/opencontrol/compliance-masonry", "", true, ""),
		Entry("Can't init / detect the repo", "https://myrepo/opencontrol/compliance-masonry", "master", false, repoInitFailed),
		Entry("Can't clone repo", "http://user:name@github.com/opencontrol/compliance-masonry-blah", "master", false, repoCloneFailed),
		Entry("Get a revision that doesn't exist", "https://github.com/opencontrol/compliance-masonry", "master-ultimate-branch-that-never-exists", false, repoCheckoutFailed),
	)
})
