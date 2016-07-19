package component

import (
	"testing"

	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/stretchr/testify/assert"
)

func TestComponentGetters(t *testing.T) {
	testSatisfies := []Satisfies{
		{
			ControlOrigin:          "inherited",
			ControlOrigins:         []string{"inherited"},
			ImplementationStatus:   "partial",
			ImplementationStatuses: []string{"partial"},
			Parameters:             []Section{Section{Key: "key", Text: "text"}},
			Narrative: []NarrativeSection{
				NarrativeSection{Key: "key", Text: "text"},
				NarrativeSection{Text: "text"},
			},
		}, {}}
	component := Component{
		Name:            "Amazon Elastic Compute Cloud",
		Key:             "EC2",
		ResponsibleRole: "AWS Staff",
		References:      common.GeneralReferences{{}},
		Verifications:   common.VerificationReferences{{}, {}},
		Satisfies:       testSatisfies,
		SchemaVersion:   semver.MustParse("3.1.0"),
	}
	// Test the getters
	assert.Equal(t, "EC2", component.GetKey())
	assert.Equal(t, "Amazon Elastic Compute Cloud", component.GetName())
	assert.Equal(t, &common.GeneralReferences{{}}, component.GetReferences())
	assert.Equal(t, &common.VerificationReferences{{}, {}}, component.GetVerifications())
	assert.Equal(t, semver.MustParse("3.1.0"), component.GetVersion())
	assert.Equal(t, "AWS Staff", component.GetResponsibleRole())
	assert.Equal(t, len(testSatisfies), len(component.GetAllSatisfies()))
	for idx, satisfies := range component.GetAllSatisfies() {
		assert.Equal(t, satisfies.GetControlKey(), testSatisfies[idx].GetControlKey())
		assert.Equal(t, satisfies.GetStandardKey(), testSatisfies[idx].GetStandardKey())
		assert.Equal(t, satisfies.GetNarratives(), testSatisfies[idx].GetNarratives())
		for i, narrative := range satisfies.GetNarratives() {
			assert.Equal(t, satisfies.GetNarratives()[i].GetKey(), narrative.GetKey())
			assert.Equal(t, satisfies.GetNarratives()[i].GetText(), narrative.GetText())
		}
		assert.Equal(t, satisfies.GetParameters(), testSatisfies[idx].GetParameters())
		for i, parameter := range satisfies.GetParameters() {
			assert.Equal(t, satisfies.GetParameters()[i].GetKey(), parameter.GetKey())
			assert.Equal(t, satisfies.GetParameters()[i].GetText(), parameter.GetText())
		}
		assert.Equal(t, satisfies.GetCoveredBy(), testSatisfies[idx].GetCoveredBy())
		assert.Equal(t, satisfies.GetControlOrigin(), testSatisfies[idx].GetControlOrigin())
		assert.Equal(t, satisfies.GetControlOrigins(), testSatisfies[idx].GetControlOrigins())
		assert.Equal(t, satisfies.GetImplementationStatus(), testSatisfies[idx].GetImplementationStatus())
		assert.Equal(t, satisfies.GetImplementationStatuses(), testSatisfies[idx].GetImplementationStatuses())
	}
}

func TestComponentSetters(t *testing.T) {
	component := Component{}
	// Test the setters.
	// Change the version.
	component.SetVersion(semver.MustParse("3.1.0"))
	assert.Equal(t, semver.MustParse("3.1.0"), component.GetVersion())
	// Change the key.
	component.SetKey("FooKey")
	assert.Equal(t, "FooKey", component.GetKey())
}
