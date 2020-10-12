package states_test

import (
	"testing"

	"github.com/thersanchez/aoc/2017/day06a/mem"
	"github.com/thersanchez/aoc/2017/day06b/states"
)

func TestStates_HasWhenEmpty(t *testing.T) {
	t.Parallel()

	states := states.NewStates()

	var m mem.Mem
	{
		banks := []int{42} // irrelevant
		var err error

		m, err = mem.NewMem(banks)
		if err != nil {
			t.Fatal(err)
		}
	}

	_, ok := states.Has(m)
	if ok {
		t.Error("unexpected mem found in states")
	}
}

func TestStates_HasWhenMissing(t *testing.T) {
	t.Parallel()

	states := states.NewStates()

	// m1 and m2 have different hashes
	var m1, m2 mem.Mem
	{
		var err error

		banks1 := []int{1}
		banks2 := []int{2}

		m1, err = mem.NewMem(banks1)
		if err != nil {
			t.Fatal(err)
		}

		m2, err = mem.NewMem(banks2)
		if err != nil {
			t.Fatal(err)
		}
	}

	var step int // irrelevant

	states.Add(m2, step)

	_, ok := states.Has(m1)
	if ok {
		t.Error("unexpected m1 found in states")
	}
}

func TestStates_HasOK(t *testing.T) {
	t.Parallel()

	states := states.NewStates()

	var m mem.Mem
	{
		banks := []int{42} // irrelevant
		var err error

		m, err = mem.NewMem(banks)
		if err != nil {
			t.Fatal(err)
		}
	}

	step := 42

	states.Add(m, step)

	got, ok := states.Has(m)
	if !ok {
		t.Error("missing m")
	}

	if step != got {
		t.Errorf("wrong step, got %d, want %d", got, step)
	}
}
