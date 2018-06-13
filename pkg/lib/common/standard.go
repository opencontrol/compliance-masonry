package common

//go:generate mockery -name Standard -testonly

// Standard is the container of all the information for a particular Standard.
// Schema info: https://github.com/opencontrol/schemas#standards-documentation
//
// GetName returns the name
//
// GetControls returns all controls associated with the standard
//
// GetControl returns a particular control
//
// GetSortedControls returns a list of sorted controls
type Standard interface {
	GetName() string
	GetControls() map[string]Control
	GetControl(string) Control
	GetSortedControls() []string
}
