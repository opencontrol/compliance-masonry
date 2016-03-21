package schema

import (
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	var parseTests = []struct {
		data                []byte
		expectedSchema      Schema
		expectedErrorExists bool
		expectedErrorText   string
	}{
		{
			data: []byte(`
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
`),
			expectedSchema: Schema{
				Base:       common.Base{SchemaVersion: "1.0.0"},
				SystemName: "test-system",
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
				Dependencies: Dependencies{
					Certification: common.Entry{
						URL:      "github.com/18F/LATO",
						Revision: "master",
					},
					Systems: []common.Entry{
						common.Entry{
							URL:      "github.com/18F/cg-complinace",
							Revision: "master",
						},
					},
					Standards: []common.Entry{
						common.Entry{
							URL:      "github.com/18F/NIST-800-53",
							Revision: "master",
						},
					},
				},
			},
			expectedErrorExists: false,
		},
		{
			// Malformed yaml. Tabbed over.
			data: []byte(`
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
			`),
			expectedSchema:      Schema{},
			expectedErrorExists: true,
			expectedErrorText:   ErrMalformedV1_0_0YamlPrefix,
		},
	}
	for _, test := range parseTests {
		s := Schema{}
		err := s.Parse(test.data)
		assert.Equal(t, test.expectedSchema, s)
		if test.expectedErrorExists {
			assert.Contains(t, err.Error(), ErrMalformedV1_0_0YamlPrefix)
		}
		assert.Equal(t, test.expectedErrorExists, err != nil)
	}
}
