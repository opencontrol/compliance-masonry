package common

type Standard interface {
	GetName() string
	GetControls() map[string]Control
	GetControl(string) Control
	GetSortedControls() []string
}
