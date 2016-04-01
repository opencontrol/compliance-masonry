package docx

// Config contains data for docx template export configurations
type Config struct {
	OpencontrolDir string
	Certification  string
	TemplatePath   string
	ExportPath     string
}

//BuildDocx exports a Doxc ssp based on a template
func (config *Config) BuildDocx() {
}
