package common

// Entry is a generic holder for handling the specific location and revision of a resource.
type Entry struct {
	URL      string `yaml:"url"`
	Revision string `yaml:"revision"`
}

