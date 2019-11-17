package mem

// States stores copies of the states of a memory.
type States []Mem

// StoreCopy stores a copy of the memory.
func (s States) StoreCopy(m Mem) {

}

// Find returns if the state of m is already stored in s.
func (s States) Find(m Mem) bool {
	return false
}
