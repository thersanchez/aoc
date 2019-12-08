package mem

import "fmt"

// Mem is a memory that contains a number of banks, each bank contains any number of blocks.
type Mem struct {
	banks []int
}

// NewMem returns a new Mem with a copy of the given banks.
func NewMem(banks []int) (Mem, error) {
	dup := make([]int, len(banks))
	for i, v := range banks {
		dup[i] = v
	}

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

// RedistributeBlocks removes all of the blocks from the given bank, then moves to the next (by index)
// memory bank and inserts one of the blocks. It continues doing this until it runs out of blocks;
// if it reaches the last memory bank, it wraps around to the first one.
func (m Mem) RedistributeBlocks(pos int) error {
	if pos < 0 || pos >= len(m.banks) {
		return fmt.Errorf("invalid pos (%d)", pos)
	}
	return nil
}
