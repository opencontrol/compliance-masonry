package inventory

type Inventory struct {
}

type Config struct {
	Certification string
	OpencontrolDir string
}

func (i Inventory) ComputeGapAnalysis(config Config) []interface{} {
	return nil
}