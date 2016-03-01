package models

import "testing"

func TestLoadComponent(t *testing.T) {
	system := &System{Components: make(map[string]*Component)}
	componentDir := "./opencontrol_fixtures/components/AWS/EC2"
	system.LoadComponent(componentDir)
	componentKey := system.Components["EC2"].Key
	if "EC2" != componentKey {
		t.Errorf("Expected %s, Actual: %s", "EC2", componentKey)
	}

}
