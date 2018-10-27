package memory

import (
	"fmt"
)

// CalculateValue calculates the value at the memory position p.
func CalculateValue(p int) int {
	vv := map[int]int{0: 1}
	return calculateValue(p, vv)
}

func calculateValue(p int, vv map[int]int) int {
	if p == 0 {
		return vv[0]
	}

	sum := 0
	for _, n := range neighbours(p) {
		if n < p {
			if _, ok := vv[n]; !ok {
				vv[n] = calculateValue(n, vv)
			}
			sum += vv[n]
		}
	}

	return sum
}

// PosToCoord returns the spiral coordinate of a memory position.
// Returns an error with negative positions.
func PosToCoord(p int) (Coord, error) {
	if p == 0 {
		return Coord{X: 0, Y: 0}, nil
	}
	side, err := RingSideFromPos(p)
	if err != nil {
		return Coord{}, fmt.Errorf(
			"calculating ring side: negative position (%d)", p)
	}

	corners, err := NewCorners(side)
	if err != nil {
		return Coord{}, fmt.Errorf("calculating corners: %v", err)
	}

	for _, c := range corners {
		if p == c.Pos {
			return c.Coord, nil
		}
	}

	zone, err := PosToZone(p)
	if err != nil {
		return Coord{}, fmt.Errorf("calculating zone: %v", err)
	}

	halfSide := (side - 1) / 2
	switch zone {
	case Top:
		dist := (corners[1].Pos - p)
		return Coord{X: -halfSide + dist, Y: halfSide}, nil
	case Left:
		dist := (corners[2].Pos - p)
		return Coord{X: -halfSide, Y: -halfSide + dist}, nil
	case Bottom:
		dist := (corners[3].Pos - p)
		return Coord{X: halfSide - dist, Y: -halfSide}, nil
	case Right:
		dist := (corners[0].Pos - p)
		return Coord{X: halfSide, Y: halfSide - dist}, nil
	default:
		panic(fmt.Sprintf("unknown zone: %d", zone))
	}
}

func neighbours(p int) [8]int {
	c, err := PosToCoord(p)
	if err != nil {
		panic(fmt.Sprintf("calculating coordinate for pos %d: %v", p, err))
	}
	neighbours := [8]Coord{
		{X: c.X + 1, Y: c.Y},
		{X: c.X + 1, Y: c.Y + 1},
		{X: c.X, Y: c.Y + 1},
		{X: c.X - 1, Y: c.Y + 1},
		{X: c.X - 1, Y: c.Y},
		{X: c.X - 1, Y: c.Y - 1},
		{X: c.X, Y: c.Y - 1},
		{X: c.X + 1, Y: c.Y - 1},
	}
	var pp [8]int
	for i, n := range neighbours {
		pp[i] = CoordToPos(n)
	}
	return pp
}

// CoordToPos returns the memory position of an spiral coordinate.
func CoordToPos(c Coord) int {
	zero := Coord{X: 0, Y: 0}
	if c == zero {
		return 0
	}

	side := RingSideFromCoord(c)

	corners, err := NewCorners(side)
	if err != nil {
		panic(fmt.Sprintf("calculating corners: %v", err))
	}

	// Solve if c is in a corner
	for _, e := range corners {
		if c == e.Coord {
			return e.Pos
		}
	}

	// Solve if c is in a side
	zone := CoordToZone(c)
	switch zone {
	case Bottom:
		d := corners[3].Coord.X - c.X
		return corners[3].Pos - d
	case Left:
		d := c.Y - corners[2].Coord.Y
		return corners[2].Pos - d
	case Top:
		d := c.X - corners[1].Coord.X
		return corners[1].Pos - d
	case Right:
		d := corners[0].Coord.Y - c.Y
		return corners[0].Pos - d
	default:
		panic(fmt.Sprintf("unknown zone: %d", zone))
	}
}

// CoordToZone returns the ring's zone for a given coordinate (c).
// Returns an indetermined value if c is located in more than one zone,
// for instance when c is a corner of the ring.
func CoordToZone(c Coord) Zone {
	absX := abs(c.X)
	absY := abs(c.Y)
	switch {
	case absY > absX && c.Y > 0:
		return Top
	case absX > absY && c.X < 0:
		return Left
	case absY > absX && c.Y < 0:
		return Bottom
	default:
		return Right
	}
}

// PosToZone returns the ring's zone for a given position (p).
// Returns an error if p is negative.
// Returns an indetermined value if p is located in more than one zone,
// for instance when p=0 or when p is located in a corner of the ring.
func PosToZone(p int) (Zone, error) {
	side, err := RingSideFromPos(p)
	if err != nil {
		return Top, fmt.Errorf(
			"calculating ring side: negative position (%d)", p)
	}

	corners, err := NewCorners(side)
	if err != nil {
		return Top, fmt.Errorf("calculating corners: %v", err)
	}

	switch {
	case p <= corners[0].Pos:
		return Right, nil
	case p <= corners[1].Pos:
		return Top, nil
	case p <= corners[2].Pos:
		return Left, nil
	case p <= corners[3].Pos:
		return Bottom, nil
	default:
		panic("internal error")
	}
}
