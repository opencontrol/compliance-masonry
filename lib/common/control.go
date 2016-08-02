package common

// Control is the interface for getting all the attributes for a given control.
// Schema info: https://github.com/opencontrol/schemas#standards-documentation
//
// GetName returns the string representation of the control.
//
// GetFamily returns which family the control belongs to.
type Control interface {
	GetName() string
	GetFamily() string
}
