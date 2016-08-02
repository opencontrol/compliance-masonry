package certifications

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/opencontrol/compliance-masonry/lib/common"
	v1_0_0 "github.com/opencontrol/compliance-masonry/lib/certifications/versions/1_0_0"
)


func Load(certificationFile string) (common.Certification, error) {
	// Only one version right now and there's no schema version right now to indicate which version.
	var certification v1_0_0.Certification
	certificationData, err := ioutil.ReadFile(certificationFile)
	if err != nil {
		return nil, common.ErrReadFile
	}
	err = yaml.Unmarshal(certificationData, &certification)
	if err != nil {
		return nil, common.ErrCertificationSchema
	}
	return certification, nil
}