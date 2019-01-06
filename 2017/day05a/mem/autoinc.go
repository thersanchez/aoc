package mem

import "fmt"

// AutoInc is a Mem that increments in 1 every value it reads.

// AutoInc wraps a ReadWritter so that each time it reads,
// the value at the address being read is incremented in one.
type AutoInc struct {
	rw ReadWriter
}

// ReadWriter knows how to read and write integers from addresses.
type ReadWriter interface {
	Read(addr int) (int, error)
	Write(addr, value int) error
}

// NewAutoInc returns a new AutoInc using the given Mem as its
// internal memory.
func NewAutoInc(rw ReadWriter) AutoInc {
	return AutoInc{
		rw: rw,
	}
}

// ReadAndInc returns the memory value at addr and increments
// that value in 1 in the memory.
// Returns an error if the address is <0 or >= than the memory size.
func (ai AutoInc) ReadAndInc(addr int) (int, error) {
	v, err := ai.rw.Read(addr)
	if err != nil {
		return 0, fmt.Errorf("reading: %v", err)
	}

	err = ai.rw.Write(addr, v+1)
	if err != nil {
		return 0, fmt.Errorf("writing: %v", err)
	}

	return v, nil
}
