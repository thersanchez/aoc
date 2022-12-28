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

func TestChildrenRepresentatives(t *testing.T) {
	t.Parallel()

	parent, err := day07b.NewNode("parent", 41)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	child1, err := day07b.NewNode("id_child_1", 45)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	child2, err := day07b.NewNode("id_child_2", 45)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	child3, err := day07b.NewNode("id_child_3", 68)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	parent.AddChildren(child1)
	parent.AddChildren(child2)
	parent.AddChildren(child3)

	wantCorrect := child2 // or child1
	wantIncorrect := child3

	gotCorrect, gotIncorrect, err := childrenRepresentatives(parent)
	if err != nil {
		t.Fatal(err)
	}

	if wantIncorrect != gotIncorrect {
		t.Errorf("wrong incorrect node, want %s, got %s",
			wantIncorrect.Id(), gotIncorrect.Id())
	}

	if wantCorrect.TotalWeight() != gotCorrect.TotalWeight() {
		t.Errorf("wrong correct node, want %s, got %s",
			wantCorrect.Id(), gotCorrect.Id())
	}

}
