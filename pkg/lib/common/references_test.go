package common

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
	// Load a GeneralReferences struct that has 0 GeneralReference(s) to verify Len method returns 0
	{GeneralReferences{}, 0},
	// Load a GeneralReferences struct that has 1 GeneralReference to verify Len method returns 1
	{GeneralReferences{GeneralReference{}}, 1},
	// Load a GeneralReferences struct that has 2 GeneralReference(s) to verify Len method returns 2
	{GeneralReferences{GeneralReference{}, GeneralReference{}}, 2},
}

func TestGeneralReferencesLen(t *testing.T) {
	for _, example := range generalReferenceLenTests {
		actualLength := example.references.Len()
		// Check that the expected length of GeneralReferences is the actual length
		if example.expectedLength != actualLength {
			t.Errorf("Expected %d, Actual: %d", example.expectedLength, actualLength)
		}
	}
}

var verificationReferencesLenTests = []verificationReferencesLenTest{
	// Load a VerificationReferences struct that has 0 VerificationReference(s) to verify Len method returns 0
	{VerificationReferences{}, 0},
	// Load a VerificationReferences struct that has 0 VerificationReference to verify Len method returns 0
	{VerificationReferences{VerificationReference{}}, 1},
	// Load a VerificationReferences struct that has 0 VerificationReference(s) to verify Len method returns 0
	{VerificationReferences{VerificationReference{}, VerificationReference{}}, 2},
}

func TestVerificationReferencesLen(t *testing.T) {
	for _, example := range verificationReferencesLenTests {
		actualLength := example.references.Len()
		// Check that the expected length of VerificationReferences is the actual length
		if example.expectedLength != actualLength {
			t.Errorf("Expected %d, Actual: %d", example.expectedLength, actualLength)
		}
	}
}

var generalReferencesLessTests = []generalReferencesLessTest{
	// Verify that the left is greater than the right when given 2 numbers
	{GeneralReferences{GeneralReference{Name: "1"}, GeneralReference{Name: "2"}}, true},
	// Verify that the left is not greater than the right when 2 of the same
	{GeneralReferences{GeneralReference{Name: "a"}, GeneralReference{Name: "a"}}, false},
	// Verify that the left is not greater than the right when given a letter and number
	{GeneralReferences{GeneralReference{Name: "a"}, GeneralReference{Name: "2"}}, false},
	// Verify that the left is not greater than the right when given two numbers
	{GeneralReferences{GeneralReference{Name: "2"}, GeneralReference{Name: "1"}}, false},
}

func TestGeneralReferencesLess(t *testing.T) {
	for _, example := range generalReferencesLessTests {
		actualLeftIsLess := example.references.Less(0, 1)
		// Verify that the left is greater than the right when expected
		if example.leftIsLess != actualLeftIsLess {
			t.Errorf("Expected %t, Actual: %t", actualLeftIsLess, actualLeftIsLess)
		}
	}
}

var verificationReferencesLessTests = []verificationReferencesLessTest{
	// Verify that the left is greater than the right when given 2 numbers
	{VerificationReferences{VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "1"}}, VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "2"}}}, true},
	// Verify that the left is not greater than the right when 2 of the same
	{VerificationReferences{VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "a"}}, VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "a"}}}, false},
	// Verify that the left is not greater than the right when given a letter and number
	{VerificationReferences{VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "a"}}, VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "2"}}}, false},
	// Verify that the left is not greater than the right when given two numbers
	{VerificationReferences{VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "2"}}, VerificationReference{Key: "", GeneralReference: GeneralReference{Name: "2"}}}, false},
}

func TestVerificationReferencesLess(t *testing.T) {
	for _, example := range verificationReferencesLessTests {
		actualLeftIsLess := example.references.Less(0, 1)
		// verify that the left is greater than the right when expected
		if example.leftIsLess != actualLeftIsLess {
			t.Errorf("Expected %t, Actual: %t", actualLeftIsLess, actualLeftIsLess)
		}
	}
}

func TestGenralReferencesSwap(t *testing.T) {
	// Test that the swap method functions correctly
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
	// Test that the swap method functions correctly
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
	// Test that a VerificationReference with the key "a" is returned
	{VerificationReferences{VerificationReference{Key: "a"}}, true},
	// Test that a VerificationReference with the key "a" is returned when there are multiple keys
	{VerificationReferences{VerificationReference{Key: "a"}, VerificationReference{Key: "c"}}, true},
	// Test that a VerificationReference with the key "a" is not returned when only other keys exist
	{VerificationReferences{VerificationReference{Key: "1"}, VerificationReference{Key: "b"}}, false},
	// Test that a VerificationReference with the key "a" is not returned when empty
	{VerificationReferences{}, false},
}

func TestVerificationReferencesGet(t *testing.T) {
	for _, example := range verificationReferencesGetTests {
		found := example.references.Get("a")
		actuallyFound := false
		// Verify key a is found
		if found.Key == "a" {
			actuallyFound = true
		}
		// Check that the correct key was returned
		if example.found != actuallyFound {
			t.Errorf("Expected %t, Actual: %t", example.found, actuallyFound)
		}

	}
}
