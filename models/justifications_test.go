package models

import "testing"

type singleMapping struct {
	standardKey      string
	controlKey       string
	componentKey     string
	justificationKey string
}

type justificationsTest struct {
	mappings      []singleMapping
	expectedCount int
}

var justificationsAddTests = []justificationsTest{
	{[]singleMapping{{"s", "c", "1", "just1"}, {"s", "c", "2", "just2"}, {"s", "c", "3", "just3"}}, 1},
	{[]singleMapping{{"s1", "c", "1", "just4"}, {"s2", "c", "2", "just5"}, {"s3", "c", "3", "just6"}}, 3},
}

func TestJustificationAdd(t *testing.T) {
	for _, example := range justificationsAddTests {
		just := NewJustifications()
		for _, mapping := range example.mappings {
			just.Add(mapping.standardKey, mapping.controlKey, mapping.componentKey, mapping.justificationKey)
		}
		if example.expectedCount != len(just.mapping) {
			t.Errorf("Expected %d, Actual: %d", example.expectedCount, len(just.mapping))
		}
	}
}

var justificationsGetTests = []justificationsTest{
	{[]singleMapping{{"a", "b", "1", "just1"}, {"a", "b", "2", "just2"}, {"a", "b", "3", "just3"}}, 3},
	{[]singleMapping{{"a", "b", "1", "just1"}, {"a", "b", "2", "just2"}, {"f", "g", "3", "just3"}}, 2},
	{[]singleMapping{{"a", "b", "1", "just1"}, {"d", "e", "2", "just2"}, {"f", "g", "3", "just3"}}, 1},
}

func TestJustificationGet(t *testing.T) {
	for _, example := range justificationsGetTests {
		just := NewJustifications()
		for _, mapping := range example.mappings {
			just.Add(mapping.standardKey, mapping.controlKey, mapping.componentKey, mapping.justificationKey)
		}
		numberofABs := len(just.Get("a", "b"))
		if example.expectedCount != numberofABs {
			t.Errorf("Expected %d, Actual: %d", example.expectedCount, numberofABs)
		}
	}
}

type verificationsLenTest struct {
	verifications  Verifications
	expectedLength int
}

var verificationsLenTests = []verificationsLenTest{
	{Verifications{}, 0},
	{Verifications{Verification{}}, 1},
	{Verifications{Verification{}, Verification{}}, 2},
}

func TestVerificationsLen(t *testing.T) {
	for _, example := range verificationsLenTests {
		actualLength := example.verifications.Len()
		if example.expectedLength != actualLength {
			t.Errorf("Expected %d, Actual: %d", example.expectedLength, actualLength)
		}
	}
}

type verificationsLessTest struct {
	verifications Verifications
	leftIsLess    bool
}

var verificationsLessTests = []verificationsLessTest{
	{Verifications{Verification{Component: "1", Verification: "z"}, Verification{Component: "2", Verification: "a"}}, true},
	{Verifications{Verification{Component: "a", Verification: "z"}, Verification{Component: "a", Verification: "a"}}, false},
	{Verifications{Verification{Component: "a", Verification: "z"}, Verification{Component: "2", Verification: "a"}}, false},
	{Verifications{Verification{Component: "2", Verification: "z"}, Verification{Component: "1", Verification: "a"}}, false},
	// TODO: This should be false
	{Verifications{Verification{Component: "1zz", Verification: "z"}, Verification{Component: "1", Verification: "b"}}, false},
}

func TestVerificationsLess(t *testing.T) {
	for _, example := range verificationsLessTests {
		actualLeftIsLess := example.verifications.Less(0, 1)
		if example.leftIsLess != actualLeftIsLess {
			t.Errorf("Expected %s, Actual: %s", actualLeftIsLess, actualLeftIsLess)
		}
	}
}
