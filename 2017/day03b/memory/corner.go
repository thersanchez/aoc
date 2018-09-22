package memory

import "fmt"

type Corner struct {
	Pos   int
	Coord Coord
}

func NewCorners(side int) ([]Corner, error) {
	if side <= 0 {
		return nil, fmt.Errorf("negative side (%d)", side)
	}
	if side%2 == 0 {
		return nil, fmt.Errorf("even side (%d)", side)
	}
	if side == 1 {
		return []Corner{{Pos: 0, Coord: Coord{X: 0, Y: 0}}}, nil
	}

	halfSide := (side - 1) / 2
	br := Corner{
		Pos:   (side * side) - 1,
		Coord: Coord{X: halfSide, Y: -halfSide},
	}
	bl := Corner{
		Pos:   br.Pos - (side - 1),
		Coord: Coord{X: -halfSide, Y: -halfSide},
	}
	tl := Corner{
		Pos:   bl.Pos - (side - 1),
		Coord: Coord{X: -halfSide, Y: halfSide},
	}
	tr := Corner{
		Pos:   tl.Pos - (side - 1),
		Coord: Coord{X: halfSide, Y: halfSide},
	}

	return []Corner{tr, tl, bl, br}, nil
}
