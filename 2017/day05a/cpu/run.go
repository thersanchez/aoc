package cpu

// AutoIncMem is a memory that increments in 1 every value it reads.
// Returns an error if the address is <0 or >= than the memory size.
type AutoIncMem interface {
	ReadAndInc(int) (int, error)
}

// Run executes the program at the address 0 of AutoIncMem m.
func Run(m AutoIncMem) int {
	var pc, ir, jumps int
	var err error

	for {
		// Loads next instruction.
		ir, err = m.ReadAndInc(pc)
		if err != nil {
			return jumps
		}

		// Jump
		pc = pc + ir
		jumps++
	}
}
