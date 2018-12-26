package mem

import "fmt"

// Mem represents the computer memory of fixed size.
type Mem struct {
	size int
	data []int
}

// New returns a new memory with the given size and all values set to zero.
// Returns an error if the size is <0.
func New(size int) (Mem, error) {
	if size < 0 {
		return Mem{}, fmt.Errorf("negative size")
	}
	return Mem{
		size: size,
		data: make([]int, size),
	}, nil
}

// Read returns the value stored in the address addr.
// Returns an error if the address is <0 or >= than the memory size.
func (m Mem) Read(addr int) (int, error) {
	return 0, fmt.Errorf("TODO test me")
}

// Write stores the value at the address addr.
// Returns an error if the address is <0 or >= than the memory size.
func (m Mem) Write(value, addr int) error {
	return fmt.Errorf("TODO test me")
}
