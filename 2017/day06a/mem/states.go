package mem

// States is a collection of memory hashes.
type States struct {
	data map[string]struct{}
}

// Add adds the memory hash h to s.
func (s *States) Add(h string) {

}

// Find returns if the state of m is already stored in s.
func (s *States) Has(m Mem) bool {
	return false
}
