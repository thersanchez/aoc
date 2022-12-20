package main

import (
	"testing"

	"github.com/thersanchez/aoc/2017/day07b"
)

func TestCorrectedWeight(t *testing.T) {
	t.Parallel()

	want := 42

	correctNode, err := day07b.NewNode("correctId", 42)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	incorrectNode, err := day07b.NewNode("incorrectId", 3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	got := correctedWeight(correctNode, incorrectNode)

	if want != got {
		t.Errorf("wrong weight, want %d, got %d", want, got)
	}

}
