package cpu

// AutoIncMem is a memory that increments in 1 every value it reads.
type AutoIncMem interface {
	ReadAndInc(int) (int, error)
}

// Run execute the program at the address 0 of memory m.
func Run(m AutoIncMem) int {
	var pc, ir, steps int
	var err error

	for {
		// Stores the Mem Value at PC Address into IR
		ir, err = m.ReadAndInc(pc)
		if err != nil {
			return steps
		}

		// Jump
		pc = pc + ir
	}
}
