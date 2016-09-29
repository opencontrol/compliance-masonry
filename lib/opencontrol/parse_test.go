package opencontrol_test

import (
	. "github.com/opencontrol/compliance-masonry/lib/opencontrol"

	. "github.com/onsi/ginkgo"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol/mocks"
	"github.com/opencontrol/compliance-masonry/lib/opencontrol/versions/1.0.0"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Parse", func() {
	var (
		parser      SchemaParser
		err         error
		openControl common.OpenControl
	)

	BeforeEach(func() {
		parser = new(mocks.SchemaParser)
	})

	Describe("bad input scenarios", func() {
		BeforeEach(func() {
			parser = YAMLParser{}
		})
		It("should detect there's no data to parse when given nil data", func() {
			openControl, err = parser.Parse(nil)
			assert.Equal(GinkgoT(), common.ErrNoDataToParse, err)
		})
		It("should detect there's no data to parse when given empty data", func() {
			openControl, err = parser.Parse([]byte(""))
			assert.Equal(GinkgoT(), common.ErrNoDataToParse, err)
		})
		It("should detect when it's unable to unmarshal into the base type", func() {
			openControl, err = parser.Parse([]byte("schema_version: @"))
			assert.Contains(GinkgoT(), err.Error(), ErrMalformedBaseYamlPrefix)
		})
		It("should detect when it's unable to determine the semver version because it is not in the format", func() {
			openControl, err = parser.Parse([]byte("schema_version: versionone"))
			assert.Equal(GinkgoT(), err, common.ErrCantParseSemver)
		})
		It("should detect when it's unable to determine the semver version because the version is not in string quotes", func() {
			openControl, err = parser.Parse([]byte(`schema_version: 1.0`))
			assert.Equal(GinkgoT(), err, common.ErrCantParseSemver)
		})
		It("should detect when the version is unknown", func() {
			openControl, err = parser.Parse([]byte(`schema_version: "0.0.0"`))
			assert.Equal(GinkgoT(), err, common.ErrUnknownSchemaVersion)
		})
	})
})

var _ = Describe("Parsing the scchema", func() {
	Describe("Parsing v1.0.0", func() {
		data := []byte(`
schema_version: "1.0.0"
name: test
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
components:
  - ./component-1
  - ./component-2
  - ./component-3
certifications:
  - ./cert-1.yaml
standards:
  - ./standard-1.yaml
dependencies:
  certifications:
    - url: github.com/18F/LATO
      revision: master
  systems:
    - url: github.com/18F/cg-complinace
      revision: master
  standards:
    - url: github.com/18F/NIST-800-53
      revision: master
`)
		It("should successfully parse", func() {
			parser := YAMLParser{}
			opencontrol, err := parser.Parse(data)
			assert.Nil(GinkgoT(), err)
			assert.Equal(GinkgoT(), []string{"./cert-1.yaml"}, opencontrol.GetCertifications())
			assert.Equal(GinkgoT(), []string{"./standard-1.yaml"}, opencontrol.GetStandards())
			assert.Equal(GinkgoT(), []string{"./component-1", "./component-2", "./component-3"}, opencontrol.GetComponents())
			assert.Equal(GinkgoT(), []common.RemoteSource{schema.VCSEntry{URL: "github.com/18F/NIST-800-53", Revision: "master", Path: ""}}, opencontrol.GetStandardsDependencies())
			assert.Equal(GinkgoT(), []common.RemoteSource{schema.VCSEntry{URL: "github.com/18F/cg-complinace", Revision: "master", Path: ""}}, opencontrol.GetComponentsDependencies())
			assert.Equal(GinkgoT(), []common.RemoteSource{schema.VCSEntry{URL: "github.com/18F/LATO", Revision: "master", Path: ""}}, opencontrol.GetCertificationsDependencies())

		})
	})
	Describe("Parsing a bad aligned yaml", func() {
		data := []byte(`
			schema_version: "1.0.0"
			system_name: test-system
			metadata:
			  description: "A system to test parsing"
			  maintainers:
			    - test@test.com
			components:
			  - ./component-1
			  - ./component-2
			  - ./component-3
			dependencies:
			  certification:
			    url: github.com/18F/LATO
			    revision: master
			  systems:
			    - url: github.com/18F/cg-complinace
			      revision: master
			  standards:
			    - url: github.com/18F/NIST-800-53
			      revision: master
			`)
		It("should unsuccessfully parse", func() {
			parser := YAMLParser{}
			opencontrol, err := parser.Parse(data)
			assert.Equal(GinkgoT(), "Unable to parse yaml data - yaml: line 1: found character that cannot start any token", err.Error())
			assert.Nil(GinkgoT(), opencontrol)
		})
	})
})
