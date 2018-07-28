package memory

// Memory defines an infinite storage of ints addresable by (pos)ition.
type Memory interface {
	Put(Pos, int)
	// NeighboursSum
	NeighboursSum(Pos) int
}

// Pos defines a memory position.
type Pos uint
