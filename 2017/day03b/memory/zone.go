package memory

import "fmt"

// Zone defines the four regions of a ring.
type Zone int

const (
	// Top is the top row.
	Top Zone = iota
	// Bottom is the bottom row.
	Bottom
	// Left is the left column.
	Left
	// Right is the right column.
	Right
)

// String pretty prints a Zone.
func (z Zone) String() string {
	switch z {
	case Top:
		return "Top"
	case Bottom:
		return "Bottom"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		panic(fmt.Sprintf("unknown zone (%d)", int(z)))
	}
}
