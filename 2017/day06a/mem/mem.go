package mem

// Mem is a memory that contains 16 banks, each bank contains any number of blocks.
type Mem [16]int

// FindMostCrowded retuns the index of the bank with most blocks.
// Ties won by the lowest-numbered memory bank.
func (Mem) FindMostCrowded() int {
	return 0
}

// RedistributeBlocks removes all of the blocks from the given bank, then moves to the next (by index)
// memory bank and inserts one of the blocks. It continues doing this until it runs out of blocks;
// if it reaches the last memory bank, it wraps around to the first one.
func (Mem) RedistributeBlocks(pos int) {

}
