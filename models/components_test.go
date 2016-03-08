package models

import "testing"

func TestLoadComponent(t *testing.T) {
	openControl := &OpenControl{Components: make(map[string]*Component)}
	componentDir := "./opencontrol_fixtures/components/EC2"
	openControl.LoadComponent(componentDir)
	componentKey := openControl.Components["EC2"].Key
	if "EC2" != componentKey {
		t.Errorf("Expected %s, Actual: %s", "EC2", componentKey)
	}
}
