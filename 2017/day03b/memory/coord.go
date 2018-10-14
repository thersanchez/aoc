package memory

import "fmt"

// Coord represents a coordinate position in an spiral memory.
type Coord struct {
	X, Y int
}

// String pretty prints a coord.
func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

// MaxAbsComponent returns the maximun of the absolute value of
// the components of c.
func (c Coord) MaxAbsComponent() int {
	x := abs(c.X)
	y := abs(c.Y)
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
