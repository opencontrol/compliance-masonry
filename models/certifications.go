package models

type Certification struct {
	Key       string              `yaml:"name" json:"name"`
	Standards map[string]Standard `yaml:"standards" json:"standards"`
}
