/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package common_test

import (
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"testing"
)

type verificationsLenTest struct {
	verifications  common.Verifications
	expectedLength int
}

type verificationsLessTest struct {
	verifications common.Verifications
	leftIsLess    bool
}

var verificationsLenTests = []verificationsLenTest{
	// Check that the number of verifications stored is 0
	{common.Verifications{}, 0},
	// Check that the number of verifications stored is 1
	{common.Verifications{common.Verification{}}, 1},
	// Check that the number of verifications stored is 2
	{common.Verifications{common.Verification{}, common.Verification{}}, 2},
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
	{common.Verifications{common.Verification{ComponentKey: "1", SatisfiesData: nil}, common.Verification{ComponentKey: "2", SatisfiesData: nil}}, true},
	// Check that the left verification is not less by comparing two letters
	{common.Verifications{common.Verification{ComponentKey: "a", SatisfiesData: nil}, common.Verification{ComponentKey: "a", SatisfiesData: nil}}, false},
	// Check that the left verification is not less by comparing the same letter
	{common.Verifications{common.Verification{ComponentKey: "a", SatisfiesData: nil}, common.Verification{ComponentKey: "2", SatisfiesData: nil}}, false},
	// Check that the left verification is not less by comparing two numbers
	{common.Verifications{common.Verification{ComponentKey: "2", SatisfiesData: nil}, common.Verification{ComponentKey: "1", SatisfiesData: nil}}, false},
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
