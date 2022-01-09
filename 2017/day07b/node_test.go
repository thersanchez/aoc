package day07b_test

import (
	"strconv"
	"testing"

	"github.com/thersanchez/aoc/2017/day07b"
)

func TestNodeId(t *testing.T) {
	t.Parallel()

	want := "some_id"
	n, err := day07b.NewNode(want, 42)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	got := n.Id()
	if want != got {
		t.Errorf("wrong id, want %q, got %q", want, got)
	}
}

func TestNonPositiveWeightAreRejected(t *testing.T) {
	t.Parallel()

	weights := []int{0, -1, -42}

	for _, w := range weights {
		w := w
		t.Run(strconv.Itoa(w), func(t *testing.T) {
			t.Parallel()

			_, err := day07b.NewNode("id", w)
			if err == nil {
				t.Errorf("unexpected success")
			}
		})
	}
}

func TestEmptyId(t *testing.T) {
	t.Parallel()

	_, err := day07b.NewNode("", 42)
	if err == nil {
		t.Errorf("unexpected success")
	}
}

func TestTotalWeight_OneNode(t *testing.T) {
	want := 42
	n, err := day07b.NewNode("some_id", want)
	if err != nil {
		t.Fatalf("failed to create node: %v", err)
	}

	got := n.TotalWeight()

	if got != want {
		t.Errorf("wrong weight, want %d, got %d", want, got)
	}
}

func TestTotalWeight_OneParentOneChildren(t *testing.T) {
	pWeight := 42
	cWeight := 13
	want := pWeight + cWeight

	p, err := day07b.NewNode("parent", pWeight)
	if err != nil {
		t.Fatalf("failed to create parent node: %v", err)
	}

	c, err := day07b.NewNode("child", cWeight)
	if err != nil {
		t.Fatalf("failed to create child node: %v", err)
	}

	p.AddChildren(c)

	got := p.TotalWeight()

	if got != want {
		t.Errorf("wrong weight, want %d, got %d", want, got)
	}
}

func TestTotalWeight_BigTree(t *testing.T) {
	createNode := func(w int) *day07b.Node {
		id := strconv.Itoa(w)
		node, err := day07b.NewNode(id, w)
		if err != nil {
			t.Fatalf("failed to create node %s: %v", id, err)
		}
		return node
	}

	n1 := createNode(1)
	n2 := createNode(2)
	n3 := createNode(3)
	n4 := createNode(4)
	n5 := createNode(5)
	n6 := createNode(6)
	n7 := createNode(7)

	n3.AddChildren(n7)
	n2.AddChildren(n4)
	n2.AddChildren(n5)
	n2.AddChildren(n6)
	n1.AddChildren(n2)
	n1.AddChildren(n3)

	got := n1.TotalWeight()
	want := 28
	if got != want {
		t.Errorf("wrong weight, want %d, got %d", want, got)
	}
}
