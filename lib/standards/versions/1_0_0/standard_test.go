package standard

import "testing"

type controlOrderTest struct {
	standard      Standard
	expectedOrder string
}

var controlOrderTests = []controlOrderTest{
	// Verify that numeric controls are ordered correctly
	{
		Standard{
			Controls: map[string]Control{"3": Control{}, "2": Control{}, "1": Control{}},
		},
		"123",
	},
	// Verify that alphabetical controls are ordered correctly
	{
		Standard{
			Controls: map[string]Control{"c": Control{}, "b": Control{}, "a": Control{}},
		},
		"abc",
	},
	// Verify that alphanumeric controls are ordered correctly
	{
		Standard{
			Controls: map[string]Control{"1": Control{}, "b": Control{}, "2": Control{}},
		},
		"12b",
	},
	// Verify that complex alphanumeric controls are ordered correctly
	{
		Standard{
			Controls: map[string]Control{"AC-1": Control{}, "AB-2": Control{}, "1.1.1": Control{}, "2.1.1": Control{}},
		},
		"1.1.12.1.1AB-2AC-1",
	},
	// Verify Natural sort order
	{
		Standard{
			Controls: map[string]Control{"AC-1": Control{}, "AC-12": Control{}, "AC-2 (1)": Control{}, "AC-2 (11)": Control{}, "AC-2 (3)": Control{}, "AC-3 (1)": Control{}},
		},
		"AC-1AC-2 (1)AC-2 (3)AC-2 (11)AC-3 (1)AC-12",
	},
}

func TestControlOrder(t *testing.T) {
	for _, example := range controlOrderTests {
		actualOrder := ""
		controlKeys := example.standard.GetSortedControls()
		for _, controlKey := range controlKeys {
			actualOrder += controlKey
		}
		// Check that the expected order is the actual order
		if actualOrder != example.expectedOrder {
			t.Errorf("Expected %s, Actual: %s", example.expectedOrder, actualOrder)
		}
	}
}

func TestGetters(t *testing.T) {
	st := Standard{Name: "test", Controls: map[string]Control{"3": Control{Name: "controlName", Family: "controlFamily", Description: "controlDescription"}, "2": Control{}, "1": Control{}}}
	if st.GetName() != "test" {
		t.Errorf("Expected standard name test. Actual %s", st.GetName())
	}
	controls := st.GetControls()
	if len(controls) != 3 {
		t.Errorf("Expected 3 controls, found %d", len(controls))
	}
	control := st.GetControl("3")
	if control.GetName() != "controlName" {
		t.Errorf("Expected control name 'controlName'. Actual %s", control.GetName())
	}
	if control.GetFamily() != "controlFamily" {
		t.Errorf("Expected control family 'controlFamily'. Actual %s", control.GetFamily())
	}
	if control.GetDescription() != "controlDescription" {
		t.Errorf("Expected control family 'controlDescription'. Actual %s", control.GetDescription())
	}
}
