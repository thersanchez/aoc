package memory_test

import (
	"fmt"
	"testing"

	"github.com/thersanchez/aoc/2017/day03b/memory"
)

func TestRingSideFromPosOK(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		input int
		want  int
	}{
		{input: 0, want: 1},
		{input: 1, want: 3},
		{input: 2, want: 3},
		{input: 3, want: 3},
		{input: 4, want: 3},
		{input: 5, want: 3},
		{input: 6, want: 3},
		{input: 7, want: 3},
		{input: 8, want: 3},
		{input: 9, want: 5},
		{input: 10, want: 5},
		{input: 11, want: 5},
		{input: 12, want: 5},
		{input: 13, want: 5},
		{input: 14, want: 5},
		{input: 15, want: 5},
		{input: 16, want: 5},
		{input: 17, want: 5},
		{input: 18, want: 5},
		{input: 19, want: 5},
		{input: 20, want: 5},
		{input: 21, want: 5},
		{input: 22, want: 5},
		{input: 23, want: 5},
		{input: 24, want: 5},
		{input: 25, want: 7},
		{input: 48, want: 7},
		{input: 49, want: 9},
		{input: 50, want: 9},
	} {
		tt := tt
		description := fmt.Sprint(tt.input)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := memory.RingSideFromPos(tt.input)
			if err != nil {
				t.Errorf("unexpected error (%v)", err)
			}
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}

func TestRingSideFromPosError(t *testing.T) {
	t.Parallel()
	for _, pos := range []int{
		-1, -2, -1000,
	} {
		pos := pos
		t.Run(fmt.Sprint(pos), func(t *testing.T) {
			t.Parallel()
			if _, err := memory.RingSideFromPos(pos); err == nil {
				t.Errorf("unexpected success (pos=%d)", pos)
			}
		})
	}
}

func TestRingSideFromCoord(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		input memory.Coord
		want  int
	}{
		{input: memory.Coord{X: 0, Y: 0}, want: 1},
		{input: memory.Coord{X: 1, Y: 0}, want: 3},
		{input: memory.Coord{X: 1, Y: 1}, want: 3},
		{input: memory.Coord{X: 0, Y: 1}, want: 3},
		{input: memory.Coord{X: -1, Y: 1}, want: 3},
		{input: memory.Coord{X: -1, Y: 0}, want: 3},
		{input: memory.Coord{X: -1, Y: -1}, want: 3},
		{input: memory.Coord{X: 0, Y: -1}, want: 3},
		{input: memory.Coord{X: 1, Y: -1}, want: 3},
		{input: memory.Coord{X: 2, Y: -1}, want: 5},
		{input: memory.Coord{X: 2, Y: 2}, want: 5},
		{input: memory.Coord{X: -1, Y: 2}, want: 5},
		{input: memory.Coord{X: -2, Y: 0}, want: 5},
		{input: memory.Coord{X: -2, Y: -2}, want: 5},
		{input: memory.Coord{X: 1, Y: -2}, want: 5},
		{input: memory.Coord{X: 2, Y: -2}, want: 5},
		{input: memory.Coord{X: 3, Y: 0}, want: 7},
		{input: memory.Coord{X: 3, Y: 3}, want: 7},
		{input: memory.Coord{X: 1, Y: 3}, want: 7},
		{input: memory.Coord{X: -2, Y: 3}, want: 7},
		{input: memory.Coord{X: -3, Y: 3}, want: 7},
		{input: memory.Coord{X: -3, Y: 0}, want: 7},
		{input: memory.Coord{X: -3, Y: -2}, want: 7},
		{input: memory.Coord{X: -3, Y: -3}, want: 7},
		{input: memory.Coord{X: -1, Y: -3}, want: 7},
		{input: memory.Coord{X: 2, Y: -3}, want: 7},
		{input: memory.Coord{X: 3, Y: -3}, want: 7},
	} {
		tt := tt
		description := fmt.Sprint(tt.input)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got := memory.RingSideFromCoord(tt.input)
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
