package models

type Control struct {
	Family string `yaml:"family" json:"family"`
	Name   string `yaml:"name" json:"name"`
}

type Standard struct {
	Key      string             `yaml:"name" json:"name"`
	Controls map[string]Control `yaml:",inline"`
}
