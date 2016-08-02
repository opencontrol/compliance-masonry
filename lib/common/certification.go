package common


type Certification interface {
	GetKey()string
	GetSortedStandards() []string
	GetStandards() map[string]Standard
}