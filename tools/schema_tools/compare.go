package schema_tools

const (
	SchemaVersionNotNeeded float32 = -18.0
)

func VerifyVersion(file string, fileType string, version, minVersion, maxVersion float32) error {
	if minVersion == SchemaVersionNotNeeded && maxVersion == SchemaVersionNotNeeded {
		return nil
	} else if version >= minVersion && maxVersion == SchemaVersionNotNeeded {
		return nil
	} else if minVersion == SchemaVersionNotNeeded && version <= SchemaVersionNotNeeded {
		return nil
	} else if version >= minVersion && version <= maxVersion {
		return nil
	}
	return NewIncompatibleSchemaError(file, fileType, version, minVersion, maxVersion)
}
