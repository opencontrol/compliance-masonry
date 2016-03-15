package models

import "testing"

type generalReferencesLenTest struct {
	references     GeneralReferences
	expectedLength int
}

type verificationReferencesLenTest struct {
	references     VerificationReferences
	expectedLength int
}

type generalReferencesLessTest struct {
	references GeneralReferences
	leftIsLess bool
}

type verificationReferencesLessTest struct {
	references VerificationReferences
	leftIsLess bool
}

type verificationReferencesGetTest struct {
	references VerificationReferences
	found      bool
}

var generalReferenceLenTests = []generalReferencesLenTest{
	{GeneralReferences{}, 0},
	{GeneralReferences{GeneralReference{}}, 1},
	{GeneralReferences{GeneralReference{}, GeneralReference{}}, 2},
}

func TestGeneralReferencesLen(t *testing.T) {
	for _, example := range generalReferenceLenTests {
		actualLength := example.references.Len()
		if example.expectedLength != actualLength {
			t.Errorf("Expected %d, Actual: %d", example.expectedLength, actualLength)
		}
	}
}

var verificationReferencesLenTests = []verificationReferencesLenTest{
	{VerificationReferences{}, 0},
	{VerificationReferences{VerificationReference{}}, 1},
	{VerificationReferences{VerificationReference{}, VerificationReference{}}, 2},
}

func TestVerificationReferencesLen(t *testing.T) {
	for _, example := range verificationReferencesLenTests {
		actualLength := example.references.Len()
		if example.expectedLength != actualLength {
			t.Errorf("Expected %d, Actual: %d", example.expectedLength, actualLength)
		}
	}
}

var generalReferencesLessTests = []generalReferencesLessTest{
	{GeneralReferences{GeneralReference{Name: "1"}, GeneralReference{Name: "2"}}, true},
	{GeneralReferences{GeneralReference{Name: "a"}, GeneralReference{Name: "a"}}, false},
	{GeneralReferences{GeneralReference{Name: "a"}, GeneralReference{Name: "2"}}, false},
	{GeneralReferences{GeneralReference{Name: "2"}, GeneralReference{Name: "1"}}, false},
}

func TestGeneralReferencesLess(t *testing.T) {
	for _, example := range generalReferencesLessTests {
		actualLeftIsLess := example.references.Less(0, 1)
		if example.leftIsLess != actualLeftIsLess {
			t.Errorf("Expected %s, Actual: %s", actualLeftIsLess, actualLeftIsLess)
		}
	}
}

var verificationReferencesLessTests = []verificationReferencesLessTest{
	{VerificationReferences{VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "1"}}, VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "2"}}}, true},
	{VerificationReferences{VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "a"}}, VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "a"}}}, false},
	{VerificationReferences{VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "a"}}, VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "2"}}}, false},
	{VerificationReferences{VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "2"}}, VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "2"}}}, false},
}

func TestVerificationReferencesLess(t *testing.T) {
	for _, example := range verificationReferencesLessTests {
		actualLeftIsLess := example.references.Less(0, 1)
		if example.leftIsLess != actualLeftIsLess {
			t.Errorf("Expected %s, Actual: %s", actualLeftIsLess, actualLeftIsLess)
		}
	}
}

func TestGenralReferencesSwap(t *testing.T) {
	references := GeneralReferences{GeneralReference{Name: "1"}, GeneralReference{Name: "2"}}
	firstName := references[0]
	secondName := references[1]
	references.Swap(0, 1)
	firstNameSwapped := references[0]
	secondNameSwapped := references[1]
	if firstName != secondNameSwapped && secondName != firstNameSwapped {
		t.Errorf("Swap Failed")
	}
}

func TestVerificationReferencesSwap(t *testing.T) {
	references := VerificationReferences{
		VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "1"}},
		VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "2"}},
	}
	firstName := references[0]
	secondName := references[1]
	references.Swap(0, 1)
	firstNameSwapped := references[0]
	secondNameSwapped := references[1]
	if firstName != secondNameSwapped && secondName != firstNameSwapped {
		t.Errorf("Swap Failed")
	}
}

var verificationReferencesGetTests = []verificationReferencesGetTest{
	{VerificationReferences{VerificationReference{Key: "a"}}, true},
	{VerificationReferences{VerificationReference{Key: "a"}, VerificationReference{Key: "c"}}, true},
	{VerificationReferences{VerificationReference{Key: "1"}, VerificationReference{Key: "b"}}, false},
	{VerificationReferences{}, false},
}

func TestVerificationReferencesGet(t *testing.T) {
	for _, example := range verificationReferencesGetTests {
		found := example.references.Get("a")
		actuallyFound := false
		if found.Key == "a" {
			actuallyFound = true
		}
		if example.found != actuallyFound {
			t.Errorf("Expected %t, Actual: %t", example.found, actuallyFound)
		}

	}
}
