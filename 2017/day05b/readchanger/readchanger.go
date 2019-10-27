package readchanger

import (
	"fmt"
)

// ReadChanger can read memories and change its contets based on the
// value being read (See ReadAndChange).
type ReadChanger struct {
	rw ReadWriter
}

// ReadWriter describes a memory from the point of view of a ReadChanger.
type ReadWriter interface {
	Read(addr int) (int, error)
	Write(addr, value int) error
}

// New returns a ReadChanger using m as its memory.
func New(m ReadWriter) *ReadChanger {
	return &ReadChanger{rw: m}
}

// ReadAndChange reads the memory address addr and returns its contents.
// Also changes the value stored there in two different ways:
// - Adds 1 if the value was less than 3.
// - Substracts 1 otherwise.
// Returns an error if addr is out memory bounds.
func (rc *ReadChanger) ReadAndChange(addr int) (int, error) {
	v, err := rc.rw.Read(addr)
	if err != nil {
		return 0, fmt.Errorf("reading: %v", err)
	}

	newValue := v - 1
	if v < 3 {
		newValue = v + 1
	}

	err = rc.rw.Write(addr, newValue)
	if err != nil {
		return 0, fmt.Errorf("writing: %v", err)
	}
	return v, nil
}
