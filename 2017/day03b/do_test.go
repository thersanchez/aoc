package main_test

import (
	"fmt"
	"testing"

	day03b "github.com/thersanchez/aoc/2017/day03b"
)

func TestDoOK(t *testing.T) {
	for _, tt := range []struct {
		input int
		want  int
	}{
		{input: -10, want: 1},
		{input: -1, want: 1},
		{input: 0, want: 1},
		{input: 1, want: 2},
		{input: 2, want: 4},
		{input: 3, want: 4},
		{input: 4, want: 5},
		{input: 5, want: 10},
		{input: 6, want: 10},
		{input: 7, want: 10},
		{input: 8, want: 10},
		{input: 9, want: 10},
		{input: 10, want: 11},
		{input: 11, want: 23},
		{input: 12, want: 23},
		{input: 13, want: 23},
		{input: 23, want: 25},
		{input: 24, want: 25},
		{input: 25, want: 26},
		{input: 26, want: 54},
		{input: 53, want: 54},
		{input: 54, want: 57},
		{input: 56, want: 57},
		{input: 57, want: 59},
		{input: 58, want: 59},
		{input: 59, want: 122},
		{input: 121, want: 122},
		{input: 122, want: 133},
		{input: 132, want: 133},
		{input: 133, want: 142},
		{input: 141, want: 142},
		{input: 142, want: 147},
		{input: 146, want: 147},
		{input: 147, want: 304},
		{input: 303, want: 304},
		{input: 304, want: 330},
		{input: 330, want: 351},
		{input: 351, want: 362},
		{input: 362, want: 747},
		{input: 747, want: 806},
		{input: 806, want: 880},
		{input: 880, want: 931},
		{input: 931, want: 957},
		{input: 37408, want: 39835},
		{input: 368078, want: 369601},
	} {
		description := fmt.Sprintf("%d", tt.input)
		t.Run(description, func(t *testing.T) {
			got := day03b.Do(tt.input)
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}

var posCoordTable = []struct {
	pos   day03b.Pos
	coord day03b.Coord
}{
	{pos: 0, coord: day03b.Coord{X: 0, Y: 0}},
	{pos: 1, coord: day03b.Coord{X: 1, Y: 0}},
	{pos: 2, coord: day03b.Coord{X: 1, Y: 1}},
	{pos: 3, coord: day03b.Coord{X: 0, Y: 1}},
	{pos: 4, coord: day03b.Coord{X: -1, Y: 1}},
	{pos: 5, coord: day03b.Coord{X: -1, Y: 0}},
	{pos: 6, coord: day03b.Coord{X: -1, Y: -1}},
	{pos: 7, coord: day03b.Coord{X: 0, Y: -1}},
	{pos: 8, coord: day03b.Coord{X: 1, Y: -1}},
	{pos: 9, coord: day03b.Coord{X: 2, Y: -1}},
	{pos: 10, coord: day03b.Coord{X: 2, Y: 0}},
	{pos: 11, coord: day03b.Coord{X: 2, Y: 1}},
	{pos: 12, coord: day03b.Coord{X: 2, Y: 2}},
	{pos: 13, coord: day03b.Coord{X: 1, Y: 2}},
	{pos: 14, coord: day03b.Coord{X: 0, Y: 2}},
	{pos: 15, coord: day03b.Coord{X: -1, Y: 2}},
	{pos: 16, coord: day03b.Coord{X: -2, Y: 2}},
	{pos: 17, coord: day03b.Coord{X: -2, Y: 1}},
	{pos: 18, coord: day03b.Coord{X: -2, Y: 0}},
	{pos: 19, coord: day03b.Coord{X: -2, Y: -1}},
	{pos: 20, coord: day03b.Coord{X: -2, Y: -2}},
	{pos: 21, coord: day03b.Coord{X: -1, Y: -2}},
	{pos: 22, coord: day03b.Coord{X: 0, Y: -2}},
	{pos: 23, coord: day03b.Coord{X: 1, Y: -2}},
	{pos: 24, coord: day03b.Coord{X: 2, Y: -2}},
	{pos: 25, coord: day03b.Coord{X: 3, Y: -2}},
	{pos: 26, coord: day03b.Coord{X: 3, Y: -1}},
	{pos: 27, coord: day03b.Coord{X: 3, Y: 0}},
	{pos: 28, coord: day03b.Coord{X: 3, Y: 1}},
	{pos: 29, coord: day03b.Coord{X: 3, Y: 2}},
	{pos: 30, coord: day03b.Coord{X: 3, Y: 3}},
	{pos: 31, coord: day03b.Coord{X: 2, Y: 3}},
	{pos: 32, coord: day03b.Coord{X: 1, Y: 3}},
	{pos: 33, coord: day03b.Coord{X: 0, Y: 3}},
	{pos: 34, coord: day03b.Coord{X: -1, Y: 3}},
	{pos: 35, coord: day03b.Coord{X: -2, Y: 3}},
	{pos: 36, coord: day03b.Coord{X: -3, Y: 3}},
	{pos: 37, coord: day03b.Coord{X: -3, Y: 2}},
	{pos: 38, coord: day03b.Coord{X: -3, Y: 1}},
	{pos: 39, coord: day03b.Coord{X: -3, Y: 0}},
	{pos: 40, coord: day03b.Coord{X: -3, Y: -1}},
	{pos: 41, coord: day03b.Coord{X: -3, Y: -2}},
	{pos: 42, coord: day03b.Coord{X: -3, Y: -3}},
	{pos: 43, coord: day03b.Coord{X: -2, Y: -3}},
	{pos: 44, coord: day03b.Coord{X: -1, Y: -3}},
	{pos: 45, coord: day03b.Coord{X: 0, Y: -3}},
	{pos: 46, coord: day03b.Coord{X: 1, Y: -3}},
	{pos: 47, coord: day03b.Coord{X: 2, Y: -3}},
	{pos: 48, coord: day03b.Coord{X: 3, Y: -3}},
	{pos: 49, coord: day03b.Coord{X: 4, Y: -3}},
}

func TestPosString(t *testing.T) {
	got := day03b.Pos(42).String()
	want := "42"
	if got != want {
		t.Errorf("want=%s, got=%s", want, got)
	}
}
func TestCoordString(t *testing.T) {
	got := day03b.Coord{X: 42, Y: -3}.String()
	want := "(42, -3)"
	if got != want {
		t.Errorf("want=%s, got=%s", want, got)
	}
}
func TestPosToCoord(t *testing.T) {
	for _, tt := range posCoordTable {
		description := fmt.Sprintf("%v", tt.pos)
		t.Run(description, func(t *testing.T) {
			got := day03b.PosToCoord(tt.pos)
			if tt.coord != got {
				t.Errorf("want=%d, got=%d", tt.coord, got)
			}
		})
	}
}

func TestCoordToPos(t *testing.T) {
	for _, tt := range posCoordTable {
		description := fmt.Sprintf("%v", tt.coord)
		t.Run(description, func(t *testing.T) {
			got := day03b.CoordToPos(tt.coord)
			if tt.pos != got {
				t.Errorf("want=%d, got=%d", tt.pos, got)
			}
		})
	}
}

func TestRingSide(t *testing.T) {
	for _, tt := range []struct {
		input day03b.Pos
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
		description := fmt.Sprintf("%v", tt.input)
		t.Run(description, func(t *testing.T) {
			got := day03b.RingSide(tt.input)
			if tt.want != got {
				t.Errorf("want=%d, got=%d", tt.want, got)
			}
		})
	}
}
