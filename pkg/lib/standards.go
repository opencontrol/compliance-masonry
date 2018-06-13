/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package lib

import (
	"sync"

	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"github.com/opencontrol/compliance-masonry/pkg/lib/standards"
)

// standardsMap struct is a thread save mapping of Standards
type standardsMap struct {
	mapping map[string]common.Standard
	sync.RWMutex
}

// newStandards creates an instance of standardsMap struct
func newStandards() *standardsMap {
	return &standardsMap{mapping: make(map[string]common.Standard)}
}

// add adds a standard to the standards mapping
func (s *standardsMap) add(standard common.Standard) {
	s.Lock()
	s.mapping[standard.GetName()] = standard
	s.Unlock()
}

// Get retrieves a standard
func (s *standardsMap) get(standardName string) (standard common.Standard, found bool) {
	s.Lock()
	defer s.Unlock()
	standard, found = s.mapping[standardName]
	return
}

// GetAll retrieves all the standards
func (s *standardsMap) getAll() []common.Standard {
	s.RLock()
	defer s.RUnlock()
	standardSlice := make([]common.Standard, len(s.mapping))
	idx := 0
	for _, value := range s.mapping {
		standardSlice[idx] = value
		idx++
	}
	return standardSlice
}

// LoadStandard imports a standard into the Standard struct and adds it to the
// main object.
func (ws *localWorkspace) LoadStandard(standardFile string) error {
	standard, err := standards.Load(standardFile)
	if err != nil {
		return common.ErrStandardSchema
	}
	ws.standards.add(standard)
	return nil
}

// GetAllStandards retrieves all standards
func (ws *localWorkspace) GetAllStandards() []common.Standard {
	return ws.standards.getAll()
}

// GetStandard retrieves a specific standard
func (ws *localWorkspace) GetStandard(standardKey string) (common.Standard, bool) {
	return ws.standards.get(standardKey)
}
