package memory_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/thersanchez/aoc/2017/day03b/memory"
)

func TestNewCorners(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		side int
		want []memory.Corner
	}{
		{
			side: 1,
			want: []memory.Corner{
				{
					Pos:   0,
					Coord: memory.Coord{X: 0, Y: 0},
				},
			},
		}, {
			side: 3,
			want: []memory.Corner{
				{
					Pos:   2,
					Coord: memory.Coord{X: 1, Y: 1},
				}, {
					Pos:   4,
					Coord: memory.Coord{X: -1, Y: 1},
				}, {
					Pos:   6,
					Coord: memory.Coord{X: -1, Y: -1},
				}, {
					Pos:   8,
					Coord: memory.Coord{X: 1, Y: -1},
				},
			},
		}, {
			side: 5,
			want: []memory.Corner{
				{
					Pos:   12,
					Coord: memory.Coord{X: 2, Y: 2},
				}, {
					Pos:   16,
					Coord: memory.Coord{X: -2, Y: 2},
				}, {
					Pos:   20,
					Coord: memory.Coord{X: -2, Y: -2},
				}, {
					Pos:   24,
					Coord: memory.Coord{X: 2, Y: -2},
				},
			},
		}, {
			side: 7,
			want: []memory.Corner{
				{
					Pos:   30,
					Coord: memory.Coord{X: 3, Y: 3},
				}, {
					Pos:   36,
					Coord: memory.Coord{X: -3, Y: 3},
				}, {
					Pos:   42,
					Coord: memory.Coord{X: -3, Y: -3},
				}, {
					Pos:   48,
					Coord: memory.Coord{X: 3, Y: -3},
				},
			},
		},
	} {
		test := test
		desc := fmt.Sprintf("side=%d", test.side)
		t.Run(desc, func(t *testing.T) {
			t.Parallel()
			got, err := memory.NewCorners(test.side)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("\nwant=%#v\n got=%#v", test.want, got)
			}
		})
	}
}

func TestNewCornersError(t *testing.T) {
	t.Parallel()
	for _, side := range []int{
		0, 2, 4, 6, -1, -2, -3, -4,
	} {
		side := side
		desc := fmt.Sprintf("side=%d", side)
		t.Run(desc, func(t *testing.T) {
			t.Parallel()
			_, err := memory.NewCorners(side)
			if err == nil {
				t.Errorf("unexpected success")
			}
		})
	}
}
