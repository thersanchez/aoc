package memory_test

import (
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
