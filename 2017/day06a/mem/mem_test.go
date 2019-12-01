package mem_test

import (
	"fmt"
	"testing"

	"github.com/thersanchez/aoc/2017/day06a/mem"
)

func TestFindMostCrowded(t *testing.T) {
	t.Parallel()

	tests := []struct {
		banks [16]int
		want  int
	}{
		{
			banks: [16]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:  0,
		},
		{
			banks: [16]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:  1,
		},
		{
			banks: [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6},
			want:  8,
		},
		{
			banks: [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			want:  15,
		},
		{
			banks: [16]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
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

}
func TestRedistributeBlocksError(t *testing.T) {
	t.Parallel()

	invalidPositions := []int{mem.NumBanks, -1, mem.NumBanks + 10, -27}

	for _, v := range invalidPositions {
		desc := fmt.Sprintf("%d", v)
		v := v
		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			banks := [16]int{}
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
