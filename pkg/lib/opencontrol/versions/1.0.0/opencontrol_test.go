/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package schema

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"

	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

var _ = Describe("Opencontrol", func() {
	Describe("Getter functions for v1.0.0", func() {
		opencontrol := OpenControl{
			Name: "test",
			Meta: Metadata{
				Description: "A system to test parsing",
				Maintainers: []string{
					"test@test.com",
				},
			},
			Components: []string{
				"./component-1",
				"./component-2",
				"./component-3",
			},
			Certifications: []string{
				"./cert-1.yaml",
			},
			Standards: []string{
				"./standard-1.yaml",
			},
			Dependencies: Dependencies{
				Certifications: []VCSEntry{
					{
						URL:      "github.com/18F/LATO",
						Revision: "master",
					},
				},
				Systems: []VCSEntry{
					{
						URL:      "github.com/18F/cg-complinace",
						Revision: "master",
					},
				},
				Standards: []VCSEntry{
					{
						URL:      "github.com/18F/NIST-800-53",
						Revision: "master",
					},
				},
			},
		}
		assert.Equal(GinkgoT(), []string{"./cert-1.yaml"}, opencontrol.GetCertifications())
		assert.Equal(GinkgoT(), []string{"./standard-1.yaml"}, opencontrol.GetStandards())
		assert.Equal(GinkgoT(), []string{"./component-1", "./component-2", "./component-3"}, opencontrol.GetComponents())
		assert.Equal(GinkgoT(), []common.RemoteSource{VCSEntry{URL: "github.com/18F/NIST-800-53", Revision: "master", Path: ""}}, opencontrol.GetStandardsDependencies())
		assert.Equal(GinkgoT(), []common.RemoteSource{VCSEntry{URL: "github.com/18F/cg-complinace", Revision: "master", Path: ""}}, opencontrol.GetComponentsDependencies())
		assert.Equal(GinkgoT(), []common.RemoteSource{VCSEntry{URL: "github.com/18F/LATO", Revision: "master", Path: ""}}, opencontrol.GetCertificationsDependencies())
	})

})

var _ = Describe("VCSEntry", func() {
	Describe("Retrieving the config file", func() {
		DescribeTable("GetConfigFile", func(e VCSEntry, expectedPath string) {
			assert.Equal(GinkgoT(), e.GetConfigFile(), expectedPath)
		},
			Entry("Empty / new base struct to return default", VCSEntry{}, constants.DefaultConfigYaml),
			Entry("overridden config file path", VCSEntry{Path: "samplepath"}, "samplepath"),
		)
	})
	Describe("GetRevision", func() {
		e := VCSEntry{Revision: "master"}
		assert.Equal(GinkgoT(), "master", e.GetRevision())
	})
	Describe("GetURL", func() {
		e := VCSEntry{URL: "testurl"}
		assert.Equal(GinkgoT(), "testurl", e.GetURL())
	})
})
