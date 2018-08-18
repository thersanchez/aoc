package main

import (
	"fmt"
	"math"
)

// Do solves day03b.
func Do(int) (int, error) {
	return 0, nil
}

// Coord represents a coordinate position in an spiral memory.
type Coord struct {
	X, Y int
}

// String pretty prints a coord.
func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

// PosToCoord returns the spiral coordinate of a memory position.
// Returns an error with negative positions.
func PosToCoord(p int) (Coord, error) {
	if p == 0 {
		return Coord{X: 0, Y: 0}, nil
	}
	side, err := RingSide(p)
	if err != nil {
		return Coord{}, fmt.Errorf(
			"calculating ring side: negative position (%d)", p)
	}
	bottomRight := (side * side) - 1
	if bottomRight == p {
		return Coord{
			X: (side - 1) / 2,
			Y: (side - 1) / -2,
		}, nil
	}
	bottomLeft := bottomRight - (side - 1)
	if bottomLeft == p {
		return Coord{
			X: (side - 1) / -2,
			Y: (side - 1) / -2,
		}, nil
	}
	topLeft := bottomLeft - (side - 1)
	if topLeft == p {
		return Coord{
			X: (side - 1) / -2,
			Y: (side - 1) / 2,
		}, nil
	}
	topRight := topLeft - (side - 1)
	if topRight == p {
		return Coord{
			X: (side - 1) / 2,
			Y: (side - 1) / 2,
		}, nil
	}
	return Coord{}, fmt.Errorf("TODO: not implemented yet")
}

// RingSide returns the side of the ring containing the given position.
// Returns an error with negative positions.
func RingSide(p int) (int, error) {
	if p < 0 {
		return 0, fmt.Errorf("negative position (%d)", p)
	}
	if p == 0 {
		return 0, nil
	}
	sqrt := int(math.Sqrt(float64(p)))
	if sqrt%2 == 0 {
		sqrt--
	}
	return sqrt + 2, nil
}

// CoordToPos returns the memory position of an spiral coordinate.
func CoordToPos(c Coord) int {
	return 42
}
