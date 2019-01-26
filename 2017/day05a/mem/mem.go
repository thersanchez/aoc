package mem

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

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

// NewFromReader creates a new Mem with just enough size to store the
// data in the reader and loads those data into it.
func NewFromReader(r io.Reader) (Mem, error) {
	scanner := bufio.NewScanner(r)
	data := []int{}
	for scanner.Scan() {
		str := scanner.Text()
		i, err := strconv.Atoi(str)
		if err != nil {
			return Mem{}, fmt.Errorf("scanning line %d: %v", len(data)+1, err)
		}
		data = append(data, i)
	}
	if err := scanner.Err(); err != nil {
		return Mem{}, fmt.Errorf("scanning: %v", err)
	}

	m, err := New(len(data))
	if err != nil {
		return Mem{},
			fmt.Errorf("creating Mem of size %d: %v", len(data), err)
	}
	for addr, value := range data {
		if err := m.Write(addr, value); err != nil {
			return Mem{},
				fmt.Errorf("writing %d at address %d: %v", value, addr, err)
		}
	}
	return m, nil
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
