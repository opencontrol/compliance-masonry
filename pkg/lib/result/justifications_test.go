package result

import "testing"

type singleMapping struct {
	standardKey  string
	controlKey   string
	componentKey string
}

type justificationsTest struct {
	mappings      []singleMapping
	expectedCount int
}

var justificationsAddTests = []justificationsTest{
	// Check that justifications can be stored
	{[]singleMapping{{"s1", "c", "1"}, {"s2", "c", "2"}, {"s3", "c", "3"}}, 3},
	// Check that justifications with the same standard and control are put into the same slice
	{[]singleMapping{{"s", "c", "1"}, {"s", "c", "2"}, {"s", "c", "3"}}, 1},
}

func TestJustificationAdd(t *testing.T) {
	for _, example := range justificationsAddTests {
		just := NewJustifications()
		for _, mapping := range example.mappings {
			just.Add(mapping.standardKey, mapping.controlKey, mapping.componentKey, nil)
		}
		// Check that the expected stored standards are the actual standards
		if example.expectedCount != len(just.mapping) {
			t.Errorf("Expected %d, Actual: %d", example.expectedCount, len(just.mapping))
		}
	}
}

var justificationsGetTests = []justificationsTest{
	// Check that the number of controls stored is 3
	{[]singleMapping{{"a", "b", "1"}, {"a", "b", "2"}, {"a", "b", "3"}}, 3},
	// Check that the number of controls stored is 2
	{[]singleMapping{{"a", "b", "1"}, {"a", "b", "2"}, {"f", "g", "3"}}, 2},
	// Check that the number of controls stored is 1
	{[]singleMapping{{"a", "b", "1"}, {"d", "e", "2"}, {"f", "g", "3"}}, 1},
}

func TestJustificationGet(t *testing.T) {
	for _, example := range justificationsGetTests {
		just := NewJustifications()
		for _, mapping := range example.mappings {
			just.Add(mapping.standardKey, mapping.controlKey, mapping.componentKey, nil)
		}
		numberofABs := len(just.Get("a", "b"))
		// Check that the number of controls stored is the expected number
		if example.expectedCount != numberofABs {
			t.Errorf("Expected %d, Actual: %d", example.expectedCount, numberofABs)
		}
	}
}
