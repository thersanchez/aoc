package mem_test

import (
	"testing"

	"github.com/thersanchez/aoc/2017/day06a/mem"
)

func TestStates_HasWhenEmpty(t *testing.T) {
	states := mem.NewStates()

	var m mem.Mem
	{
		banks := []int{42} // irrelevant
		var err error

		m, err = mem.NewMem(banks)
		if err != nil {
			t.Fatal(err)
		}
	}

	if states.Has(m) {
		t.Error("unexpected mem found in states")
	}
}

func TestStates_HasWhenMissing(t *testing.T) {
	states := mem.NewStates()

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

	states.Add(m2)

	if states.Has(m1) {
		t.Error("unexpected m1 found in states")
	}
}

func TestStates_HasOK(t *testing.T) {
	states := mem.NewStates()

	var m mem.Mem
	{
		banks := []int{42} // irrelevant
		var err error

		m, err = mem.NewMem(banks)
		if err != nil {
			t.Fatal(err)
		}
	}

	states.Add(m)

	if !states.Has(m) {
		t.Error("missing m")
	}
}
