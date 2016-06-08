package component

import (
	"testing"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/stretchr/testify/assert"
)

func TestComponentGetterAndSetter(t *testing.T) {
	component := Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    common.GeneralReferences{{}},
		Verifications: common.VerificationReferences{{}, {}},
		Satisfies:     []Satisfies{{}, {}, {}, {}},
		SchemaVersion: 2.0,
	}
	assert.Equal(t, "EC2", component.GetKey())
	assert.Equal(t, "Amazon Elastic Compute Cloud", component.GetName())
	component.SetKey("FooKey")
	assert.Equal(t, "FooKey", component.GetKey())
	assert.Equal(t, &common.GeneralReferences{{}}, component.GetReferences())
	assert.Equal(t, &common.VerificationReferences{{}, {}}, component.GetVerifications())
	testSatisfies := []Satisfies{{}, {}, {}, {}}
	assert.Equal(t, len(testSatisfies), len(component.GetAllSatisfies()))
	for idx, satisfies := range component.GetAllSatisfies() {
		assert.Equal(t, satisfies.GetControlKey(), testSatisfies[idx].GetControlKey())
		assert.Equal(t, satisfies.GetStandardKey(), testSatisfies[idx].GetStandardKey())
		assert.Equal(t, satisfies.GetNarrative(), testSatisfies[idx].GetNarrative())
		assert.Equal(t, satisfies.GetCoveredBy(), testSatisfies[idx].GetCoveredBy())
	}
}
