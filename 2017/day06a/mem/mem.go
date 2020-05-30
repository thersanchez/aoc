package mem

import "fmt"

// Mem is a memory that contains a number of banks,
// each bank contains any number of blocks.
type Mem struct {
	banks []int
}

// NewMem returns a new Mem with a copy of the given banks.
//
// Returns an error if there are no banks or if any of the banks has a
// negative number of blocks.
func NewMem(banks []int) (Mem, error) {
	if len(banks) == 0 {
		return Mem{}, fmt.Errorf("empty banks")
	}

	for i, b := range banks {
		if b < 0 {
			return Mem{}, fmt.Errorf("bank #%d has negative number of blocks", i+1)
		}
	}

	dup := make([]int, len(banks))
	copy(dup, banks)

	return Mem{banks: dup}, nil
}

// FindMostCrowded retuns the index of the bank with most blocks.
// Ties won by the lowest-numbered memory bank.
func (m Mem) FindMostCrowded() int {
	var iMax, vMax int
	for i, v := range m.banks {
		if i == 0 {
			iMax = i
			vMax = v
			continue
		}
		if v > vMax {
			iMax = i
			vMax = v
		}
	}
	return iMax
}

// RedistributeBlocks removes all of the blocks from the given bank,
// and inserts them one by one in the consecutive banks.
// If it reaches the last memory bank, it wraps around to the first one.
//
// Returns an error if the given bank position is out of memory bounds.
func (m Mem) RedistributeBlocks(pos int) error {
	if pos < 0 || pos >= len(m.banks) {
		return fmt.Errorf("invalid pos (%d)", pos)
	}

	blocks := m.banks[pos]
	m.banks[pos] = 0

	next := func(i int) int {
		if i == len(m.banks)-1 {
			return 0
		}

		return i + 1
	}

	current := next(pos)
	for ; blocks > 0; blocks-- {
		m.banks[current]++
		current = next(current)
	}

	return nil
}

// Hash returns a hash of the Mem.
func (m Mem) Hash() string {
	return fmt.Sprintf("%#v", m)
}
