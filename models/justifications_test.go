package models

import "testing"

type singleMapping struct {
	standardKey  string
	controlKey   string
	componentKey string
}

type justificationsTest struct {
	mappings          []singleMapping
	expectedStandards int
}

var justificationsTests = []justificationsTest{
	{[]singleMapping{{"s", "c", "1"}, {"s", "c", "2"}, {"s", "c", "3"}}, 1},
	{[]singleMapping{{"s1", "c", "1"}, {"s2", "c", "2"}, {"s3", "c", "3"}}, 3},
}

func TestJustificationAdd(t *testing.T) {
	for _, example := range justificationsTests {
		just := NewJustifications()
		for _, mapping := range example.mappings {
			just.Add(mapping.standardKey, mapping.controlKey, mapping.componentKey)
		}
		if example.expectedStandards != len(just.mapping) {
			t.Errorf("Expected %d, Actual: %d", example.expectedStandards, len(just.mapping))
		}
	}
}
