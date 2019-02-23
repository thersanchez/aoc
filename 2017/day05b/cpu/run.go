package cpu

// ReadChanger is a memory that changes every value it reads.
// Returns an error if the address is <0 or >= than the memory size.
type ReadChanger interface {
	ReadAndChange(int) (int, error)
}

// Run executes the program at the address 0 of ReadChanger rc.
func Run(rc ReadChanger) int {
	var pc, ir, jumps int
	var err error

	for {
		// Loads next instruction.
		ir, err = rc.ReadAndChange(pc)
		if err != nil {
			return jumps
		}

		// Jump
		pc = pc + ir
		jumps++
	}
}
