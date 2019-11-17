package mem

// Mem is a memory that contains 16 banks, each bank contains any number of blocks.
type Mem struct {
	banks [16]int
}

// NewMem returns a new Mem with the given banks.
func NewMem(banks [16]int) (Mem, error) {
	return Mem{banks: banks}, nil
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
func (Mem) RedistributeBlocks(pos int) {

}
