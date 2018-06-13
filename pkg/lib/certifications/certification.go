/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package certifications

import (
	"encoding/json"
	"errors"
	v1_0_0 "github.com/opencontrol/compliance-masonry/pkg/lib/certifications/versions/1_0_0"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Load will read the file at the given path and attempt to return a Certification object.
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

// MarshalJSON accounts for different versions
func MarshalJSON(certification common.Certification) (b []byte, e error) {
	// ABr: *punt* on getting marshal to work with interface
	var (
		bytesCertification []byte
		err                error
	)
	vCertification1_0_0, ok := certification.(v1_0_0.Certification)
	if ok {
		bytesCertification, err = json.Marshal(&vCertification1_0_0)
	} else {
		return nil, errors.New("unsupported certification version")
	}
	if err != nil {
		return nil, err
	}
	return bytesCertification, nil
}
