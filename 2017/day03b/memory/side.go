package memory

import (
	"fmt"
	"math"
)

// RingSideFromPos returns the side of the ring containing the
// given position. Returns an error with negative positions.
func RingSideFromPos(p int) (int, error) {
	if p < 0 {
		return 0, fmt.Errorf("negative position (%d)", p)
	}
	if p == 0 {
		return 1, nil
	}
	sqrt := int(math.Sqrt(float64(p)))
	if sqrt%2 == 0 {
		sqrt--
	}
	return sqrt + 2, nil
}

// RingSideFromCoord returns the side of the ring containing the
// given coordinate.
func RingSideFromCoord(c Coord) int {
	return 2*c.MaxAbsComponent() + 1
}
