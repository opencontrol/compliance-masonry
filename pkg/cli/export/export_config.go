package export

import (
	"bytes"
	"encoding/json"

	lib_certifications "github.com/opencontrol/compliance-masonry/pkg/lib/certifications"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
)

////////////////////////////////////////////////////////////////////////
// Package structures

// Config contains settings for this object
type Config struct {
	// remainder are configuration settings local to Export
	Certification   string
	OpencontrolDir  string
	DestinationFile string
	OutputFormat    OutputFormat
	Flatten         bool
	InferKeys       bool
	Docxtemplater   bool
	KeySeparator    string
}

// internal - structure for JSON / YAML output
type exportData struct {
	Certification common.Certification
	Components    []common.Component
	Standards     []common.Standard
}

// MarshalJSON provides JSON support
func (p *exportData) MarshalJSON() (b []byte, e error) {
	// start the output
	buffer := bytes.NewBufferString("{")

	// certification
	buffer.WriteString("\"certification\":")
	bytesJSON, err := lib_certifications.MarshalJSON(p.Certification)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(string(bytesJSON))

	// iterate over components
	if len(p.Components) > 0 {
		buffer.WriteString(",\"components\":[")
		for i, v := range p.Components {
			bytesJSON, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			if i > 0 {
				buffer.WriteString(",")
			}
			buffer.WriteString(string(bytesJSON))
		}
		buffer.WriteString("]")
	}

	// iterate over standards
	if len(p.Standards) > 0 {
		buffer.WriteString(",\"standards\":[")
		for i, v := range p.Standards {
			bytesJSON, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			if i > 0 {
				buffer.WriteString(",")
			}
			buffer.WriteString(string(bytesJSON))
		}
		buffer.WriteString("]")
	}

	// finish json
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

// internal - structure for all exported data
type exportOutput struct {
	Config *Config
	Data   exportData
}

// MarshalJSON provides JSON support
func (p *exportOutput) MarshalJSON() (b []byte, e error) {
	// start the output
	buffer := bytes.NewBufferString("{")

	// config section
	buffer.WriteString("\"config\":")
	bytesConfig, err := json.Marshal(p.Config)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(string(bytesConfig))

	// data section
	buffer.WriteString(",\"data\":")
	bytesData, err := json.Marshal(&p.Data)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(string(bytesData))

	// close output
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}
