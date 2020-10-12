package mem

// States is a collection of memory hashes and steps requiered to get to that memory state.
type States struct {
	data map[string]int
}

// NewStates returns a new empty States ready to use.
func NewStates() *States {
	return &States{
		data: map[string]int{},
	}
}

// Add adds the memory hash of m to s.
func (s *States) Add(m Mem, step int) {
	s.data[m.Hash()] = step
}

// Has returns the step associated with a given memory
// or false if not found.
func (s *States) Has(m Mem) (int, bool) {
	step, ok := s.data[m.Hash()]
	return step, ok
}
