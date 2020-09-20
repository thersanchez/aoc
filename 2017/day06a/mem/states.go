package mem

// States is a collection of memory hashes.
type States struct {
	data map[string]struct{}
}

// NewStates returns a new empty States ready to use.
func NewStates() *States {
	return &States{
		data: map[string]struct{}{},
	}
}

// Add adds the memory hash of m to s.
func (s *States) Add(m Mem) {
	s.data[m.Hash()] = struct{}{}
}

// Has returns true if the hash of m is already stored in s.
func (s *States) Has(m Mem) bool {
	_, ok := s.data[m.Hash()]
	return ok
}
