package common

type Base struct {
	SchemaVersion float32 `yaml:"schema_version"`
}

func (b Base) GetSchemaVersion() float32 {
	return b.SchemaVersion
}

type SchemaParser interface {
	ParseV1_0(data[] byte) (BaseSchema, error)
}

type BaseSchema interface {
	Parse(data []byte) error
}
