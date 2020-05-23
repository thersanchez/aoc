package mem_test

import (
	"fmt"
	"testing"

	"github.com/thersanchez/aoc/2017/day06a/mem"
)

func TestNewMemError(t *testing.T) {
	t.Parallel()

	invalids := [][]int{
		[]int{0, 1, 2, 3, -3, 4, 5},
		[]int{},
		nil,
	}

	for _, banks := range invalids {
		banks := banks
		desc := fmt.Sprintf("%#v", banks)
		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			m, err := mem.NewMem(banks)
			if err == nil {
				t.Errorf("unexpected success: %#v", m)
			}
		})
	}
}

func TestFindMostCrowded(t *testing.T) {
	t.Parallel()

	tests := []struct {
		banks []int
		want  int
	}{
		{
			banks: []int{1, 0, 0},
			want:  0,
		},
		{
			banks: []int{0, 1, 0},
			want:  1,
		},
		{
			banks: []int{0, 0, 1},
			want:  2,
		},
		{
			banks: []int{1, 2, 3, 2, 1, 2},
			want:  2,
		},
		{
			banks: []int{0, 1, 0, 1},
			want:  1,
		},
	}

	for _, test := range tests {
		test := test
		desc := fmt.Sprintf("%v", test.banks)
		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			m, err := mem.NewMem(test.banks)
			if err != nil {
				t.Fatalf("cannot create Mem: %v", err)
			}

			got := m.FindMostCrowded()
			if got != test.want {
				t.Errorf("want %d, got %d", test.want, got)
			}
		})
	}
}

func TestRedistributeBlocksOK(t *testing.T) {
	t.Parallel()

	tests := []struct {
		banks []int // initial banks
		pos   int   // postion to redistribute
		want  int   // the most crowded after redistributing
	}{
		{
			banks: []int{0},
			pos:   0,
			want:  0,
		},
		{
			banks: []int{0, 1},
			pos:   0,
			want:  1,
		},
		{
			banks: []int{0, 1},
			pos:   1,
			want:  0,
		},
		{
			banks: []int{0, 4, 1},
			pos:   1,
			want:  2,
		},
		{
			banks: []int{3},
			pos:   0,
			want:  0,
		},
		{
			banks: []int{1, 2, 0, 1},
			pos:   1,
			want:  3,
		},
		{
			banks: []int{3, 2, 1, 4, 3},
			pos:   3,
			want:  0,
		},
		{
			banks: []int{3, 2, 10, 4, 3},
			pos:   2,
			want:  3,
		},
	}

	for _, test := range tests {
		test := test
		desc := fmt.Sprintf("%#v, %v", test.banks, test.pos)

		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			m, err := mem.NewMem(test.banks)
			if err != nil {
				t.Fatalf("cannot create Mem: %v", err)
			}

			if err := m.RedistributeBlocks(test.pos); err != nil {
				t.Fatalf("unexpected redistribute error: %v", err)
			}

			got := m.FindMostCrowded()
			if got != test.want {
				t.Errorf("want %d, got %d", test.want, got)
			}
		})
	}
}

func TestRedistributeBlocksError(t *testing.T) {
	t.Parallel()

	const numBanks = 3
	invalidPositions := []int{numBanks, -1, numBanks + 10, -27}

	for _, v := range invalidPositions {
		desc := fmt.Sprintf("%d", v)
		v := v
		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			banks := make([]int, numBanks)
			m, err := mem.NewMem(banks)
			if err != nil {
				t.Fatalf("cannot create Mem: %v", err)
			}

			got := m.RedistributeBlocks(v)
			if got == nil {
				t.Error("unexpected success")
			}
		})
	}
}
