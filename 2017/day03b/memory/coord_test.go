package memory_test

import (
	"fmt"
	"testing"

	"github.com/thersanchez/aoc/2017/day03b/memory"
)

func TestCoordString(t *testing.T) {
	t.Parallel()
	got := memory.Coord{X: 42, Y: -3}.String()
	want := "(42, -3)"
	if got != want {
		t.Errorf("want=%s, got=%s", want, got)
	}
}

func TestMaxAbsComponent(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		input memory.Coord
		want  int
	}{
		{input: memory.Coord{X: 0, Y: 0}, want: 0},
		{input: memory.Coord{X: 1, Y: 0}, want: 1},
		{input: memory.Coord{X: 1, Y: 1}, want: 1},
		{input: memory.Coord{X: 0, Y: 1}, want: 1},
		{input: memory.Coord{X: -1, Y: 1}, want: 1},
		{input: memory.Coord{X: -1, Y: 0}, want: 1},
		{input: memory.Coord{X: -1, Y: -1}, want: 1},
		{input: memory.Coord{X: 0, Y: -1}, want: 1},
		{input: memory.Coord{X: 1, Y: -1}, want: 1},
		{input: memory.Coord{X: 2, Y: -1}, want: 2},
		{input: memory.Coord{X: 2, Y: 2}, want: 2},
		{input: memory.Coord{X: -1, Y: 2}, want: 2},
		{input: memory.Coord{X: -2, Y: 0}, want: 2},
		{input: memory.Coord{X: -2, Y: -2}, want: 2},
		{input: memory.Coord{X: 1, Y: -2}, want: 2},
		{input: memory.Coord{X: 2, Y: -2}, want: 2},
		{input: memory.Coord{X: 3, Y: 0}, want: 3},
		{input: memory.Coord{X: 3, Y: 3}, want: 3},
		{input: memory.Coord{X: 1, Y: 3}, want: 3},
		{input: memory.Coord{X: -2, Y: 3}, want: 3},
		{input: memory.Coord{X: -3, Y: 3}, want: 3},
		{input: memory.Coord{X: -3, Y: 0}, want: 3},
		{input: memory.Coord{X: -3, Y: -2}, want: 3},
		{input: memory.Coord{X: -3, Y: -3}, want: 3},
		{input: memory.Coord{X: -1, Y: -3}, want: 3},
		{input: memory.Coord{X: 2, Y: -3}, want: 3},
		{input: memory.Coord{X: 3, Y: -3}, want: 3},
	} {
		tt := tt
		description := fmt.Sprint(tt.input)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got := tt.input.MaxAbsComponent()
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
