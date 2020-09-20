package mem_test

import (
	"fmt"
	"testing"

	"github.com/thersanchez/aoc/2017/day06a/mem"
)

func TestMem_NewMemError(t *testing.T) {
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

func TestMem_FindMostCrowded(t *testing.T) {
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

func TestMem_RedistributeBlocksOK(t *testing.T) {
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

func TestMem_RedistributeBlocksError(t *testing.T) {
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

func TestMem_HashSame(t *testing.T) {
	t.Parallel()

	tests := [][]int{
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4, 5, 6},
	}

	for _, banks := range tests {
		banks := banks
		name := fmt.Sprintf("%v", banks)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			m1, err := mem.NewMem(banks)
			if err != nil {
				t.Fatal(err)
			}

			m2, err := mem.NewMem(banks)
			if err != nil {
				t.Fatal(err)
			}

			h1, h2 := m1.Hash(), m2.Hash()

			if h1 != h2 {
				t.Errorf("different hashes:\n%s\n%s", h1, h2)
			}
		})
	}
}

func TestMem_HashDifferent(t *testing.T) {
	t.Parallel()

	type pair struct {
		a, b []int
	}

	tests := []pair{
		{[]int{1}, []int{2}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2}, []int{1, 2, 3}},
		{[]int{0}, []int{0, 0}},
		{[]int{1, 2}, []int{0, 1, 0, 2, 0}},
	}

	for _, p := range tests {
		p := p
		name := fmt.Sprintf("%v, %v", p.a, p.b)

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			m1, err := mem.NewMem(p.a)
			if err != nil {
				t.Fatal(err)
			}

			m2, err := mem.NewMem(p.b)
			if err != nil {
				t.Fatal(err)
			}

			h1, h2 := m1.Hash(), m2.Hash()

			if h1 == h2 {
				t.Errorf("same hashes: %q", h1)
			}
		})
	}
}
