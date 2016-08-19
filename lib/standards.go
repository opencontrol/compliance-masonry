package lib

import (
	"sync"

	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/lib/standards"
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

// Add adds a standard to the standards mapping
func (s *standardsMap) Add(standard common.Standard) {
	s.Lock()
	s.mapping[standard.GetName()] = standard
	s.Unlock()
}

// Get retrieves a standard
func (s *standardsMap) Get(standardName string) common.Standard {
	s.Lock()
	defer s.Unlock()
	return s.mapping[standardName]
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
func (ws *LocalWorkspace) LoadStandard(standardFile string) error {
	standard, err := standards.Load(standardFile)
	if err != nil {
		return common.ErrStandardSchema
	}
	ws.Standards.Add(standard)
	return nil
}
