package mem

import "fmt"

// Mem represents the computer memory of fixed size.
type Mem struct {
	data []int
}

// New returns a new memory with the given size and all values set to zero.
// Returns an error if the size is <0.
func New(size int) (Mem, error) {
	if size < 0 {
		return Mem{}, fmt.Errorf("negative size")
	}
	return Mem{
		data: make([]int, size),
	}, nil
}

// Read returns the value stored in the address addr.
// Returns an error if the address is <0 or >= than the memory size.
func (m Mem) Read(addr int) (int, error) {
	if addr < 0 {
		return 0, fmt.Errorf("negative addr: %d", addr)
	}
	if addr >= len(m.data) {
		return 0, fmt.Errorf("addr (%d) out of bounds(%d)", addr, len(m.data))
	}
	return m.data[addr], nil
}

// Write stores the value at the address addr.
// Returns an error if the address is <0 or >= than the memory size.
func (m Mem) Write(addr, value int) error {
	if addr < 0 {
		return fmt.Errorf("negative addr: %d", addr)
	}
	if addr >= len(m.data) {
		return fmt.Errorf("addr (%d) out of bounds(%d)", addr, len(m.data))
	}
	m.data[addr] = value
	return nil
}
