package mem

import "fmt"

// AutoInc is a Mem that increments in 1 every value it reads.
type AutoInc struct {
	m Mem // tiene dentro una memeoria Mem
}

// NewAutoInc returns a new AutoInc using the given Mem as its
// internal memory.
func NewAutoInc(m Mem) AutoInc {
	return AutoInc{
		m: m,
	}
}

// ReadAndInc returns the memory value at addr and increments
// that value in 1 in the memory.
// Returns an error if the address is <0 or >= than the memory size.
func (ai AutoInc) ReadAndInc(addr int) (int, error) {
	v, err := ai.m.Read(addr)
	if err != nil {
		return 0, fmt.Errorf("reading: %v", err)
	}

	err = ai.m.Write(v+1, addr)
	if err != nil {
		return 0, fmt.Errorf("writing: %v", err)
	}

	return v, nil
}
