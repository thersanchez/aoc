package main

import (
	"fmt"
	"math"
)

// Do solves day03b.
func Do(int) int {
	return 0
}

// Pos is a memory position.
type Pos uint

// String pretty prints a pos.
func (p Pos) String() string {
	return fmt.Sprintf("%d", uint(p))
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
func PosToCoord(p Pos) Coord {
	return Coord{0, 0}
}

// RingSide returns the side of the ring containing the given position.
func RingSide(p Pos) int {
	if p == 0 {
		return 1
	}
	sqrt := int(math.Sqrt(float64(p)))
	if sqrt%2 == 0 {
		sqrt--
	}
	return sqrt + 2
}

// CoordToPos returns the memory position of an spiral coordinate.
func CoordToPos(c Coord) Pos {
	return 42
}
