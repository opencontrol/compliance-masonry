package lib

import (
	"github.com/opencontrol/compliance-masonry/lib/certifications"
)

// LoadCertification struct loads certifications into a Certification struct
// and add it to the main object.
func (ws *LocalWorkspace) LoadCertification(certificationFile string) error {
	cert, err := certifications.Load(certificationFile)
	if err != nil {
		return err
	}
	ws.Certification = cert
	return nil
}