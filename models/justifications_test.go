package models

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

type verificationsLenTest struct {
	verifications  Verifications
	expectedLength int
}

type verificationsLessTest struct {
	verifications Verifications
	leftIsLess    bool
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

func TestJustificationGetAndApply(t *testing.T) {
	for _, example := range justificationsGetTests {
		just := NewJustifications()
		for _, mapping := range example.mappings {
			just.Add(mapping.standardKey, mapping.controlKey, mapping.componentKey, nil)
		}
		just.GetAndApply("a", "b", func(actualVerificaitons Verifications) {
			numberofABs := actualVerificaitons.Len()
			// Check that the number of controls stored is the expected number
			if example.expectedCount != numberofABs {
				t.Errorf("Expected %d, Actual: %d", example.expectedCount, numberofABs)
			}
		})
	}
}

var verificationsLenTests = []verificationsLenTest{
	// Check that the number of verifications stored is 0
	{Verifications{}, 0},
	// Check that the number of verifications stored is 1
	{Verifications{Verification{}}, 1},
	// Check that the number of verifications stored is 2
	{Verifications{Verification{}, Verification{}}, 2},
}

func TestVerificationsLen(t *testing.T) {
	for _, example := range verificationsLenTests {
		actualLength := example.verifications.Len()
		// Check that the number of verifications is the expected number
		if example.expectedLength != actualLength {
			t.Errorf("Expected %d, Actual: %d", example.expectedLength, actualLength)
		}
	}
}

var verificationsLessTests = []verificationsLessTest{
	// Check that the left verification is less by comparing a number and letter
	{Verifications{Verification{ComponentKey: "1", SatisfiesData: nil}, Verification{ComponentKey: "2", SatisfiesData: nil}}, true},
	// Check that the left verification is not less by comparing two letters
	{Verifications{Verification{ComponentKey: "a", SatisfiesData: nil}, Verification{ComponentKey: "a", SatisfiesData: nil}}, false},
	// Check that the left verification is not less by comparing the same letter
	{Verifications{Verification{ComponentKey: "a", SatisfiesData: nil}, Verification{ComponentKey: "2", SatisfiesData: nil}}, false},
	// Check that the left verification is not less by comparing two numbers
	{Verifications{Verification{ComponentKey: "2", SatisfiesData: nil}, Verification{ComponentKey: "1", SatisfiesData: nil}}, false},
	// Check that the left verification is not less by comparing two numbers
}

func TestVerificationsLess(t *testing.T) {
	for _, example := range verificationsLessTests {
		actualLeftIsLess := example.verifications.Less(0, 1)
		// Check that the verification on the left is less as expected
		if example.leftIsLess != actualLeftIsLess {
			t.Errorf("Expected %t, Actual: %t", actualLeftIsLess, actualLeftIsLess)
		}
	}
}
