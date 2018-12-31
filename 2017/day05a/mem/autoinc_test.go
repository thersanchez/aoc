package mem_test

import (
	"testing"

	"github.com/thersanchez/aoc/2017/day05a/mem"
)

// Checks that everytime you read from a memory address you get the
// previous value plus one.
func TestAutoInc(t *testing.T) {
	m, err := mem.New(2)
	if err != nil {
		t.Fatal(err)
	}
	sut := mem.NewAutoInc(m)

	// reading from 0 should return 0 (next read should return 1)
	check(t, sut, 0, 0)
	// reading again from 0 should return 1 (next read should return 2)
	check(t, sut, 0, 1)
	check(t, sut, 0, 2)
	check(t, sut, 0, 3)
	check(t, sut, 0, 4)

	// reading from another addr should return 0, then 1...
	check(t, sut, 1, 0)
	check(t, sut, 1, 1)
}

func check(t *testing.T, sut mem.AutoInc, addr, want int) {
	t.Helper()
	got, err := sut.ReadAndInc(addr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
