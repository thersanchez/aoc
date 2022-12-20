package day07b_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestNodeWeight(t *testing.T) {
	t.Parallel()

	want := 42
	n, err := day07b.NewNode("test_id", want)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	got := n.Weight()
	if want != got {
		t.Errorf("wrong weight, want %d, got %d", want, got)
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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

	createNode := func(w int) *day07b.Node {
		id := strconv.Itoa(w)
		node, err := day07b.NewNode(id, w)
		if err != nil {
			t.Fatalf("failed to create node %s: %v", id, err)
		}
		return node
	}

	root := createNode(1)
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
	root.AddChildren(n2)
	root.AddChildren(n3)

	got := root.TotalWeight()
	want := 28
	if got != want {
		t.Errorf("wrong weight, want %d, got %d", want, got)
	}
}

func TestChildren_NoChildren(t *testing.T) {
	t.Parallel()

	p, err := day07b.NewNode("parent", 42)
	if err != nil {
		t.Fatalf("failed to create parent node: %v", err)
	}

	children := p.Children()
	if len(children) != 0 {
		t.Errorf("got %d children, want 0", len(children))
	}
}

func TestChildren_OneChildren(t *testing.T) {
	t.Parallel()

	p, err := day07b.NewNode("parent", 42)
	if err != nil {
		t.Fatalf("failed to create parent node: %v", err)
	}

	c, err := day07b.NewNode("child", 24)
	if err != nil {
		t.Fatalf("failed to create child node: %v", err)
	}

	p.AddChildren(c)

	children := p.Children()
	if len(children) != 1 {
		t.Errorf("got %d children, want 1", len(children))
	}

	if children[0] != c {
		t.Errorf("bad child, got %v, want %v", *children[0], *c)
	}
}

func TestChildren_TwoChildren(t *testing.T) {
	t.Parallel()

	p, err := day07b.NewNode("parent", 42)
	if err != nil {
		t.Fatalf("failed to create parent node: %v", err)
	}

	c1, err := day07b.NewNode("child1", 24)
	if err != nil {
		t.Fatalf("failed to create child node: %v", err)
	}

	c2, err := day07b.NewNode("child2", 20)
	if err != nil {
		t.Fatalf("failed to create child node: %v", err)
	}

	p.AddChildren(c1)
	p.AddChildren(c2)

	children := p.Children()
	if len(children) != 2 {
		t.Errorf("got %d children, want 1", len(children))
	}

	want := []*day07b.Node{c1, c2}
	assert.ElementsMatch(t, children, want)
}

func TestIsBalanced_Leaf(t *testing.T) {
	t.Parallel()

	leaf, err := day07b.NewNode("leaf", 42)
	if err != nil {
		t.Fatalf("creating leaf: %v", err)
	}

	if !leaf.IsBalanced() {
		t.Error("leafs should be balanced")
	}
}

func TestIsBalanced_BalancedChildren(t *testing.T) {
	t.Parallel()

	parent, err := day07b.NewNode("parent", 10)
	if err != nil {
		t.Fatalf("creating parent: %v", err)
	}

	child1, err := day07b.NewNode("child1", 20)
	if err != nil {
		t.Fatalf("creating child 1: %v", err)
	}

	child2, err := day07b.NewNode("child2", 20)
	if err != nil {
		t.Fatalf("creating child 2: %v", err)
	}

	parent.AddChildren(child1)
	parent.AddChildren(child2)
	parent.TotalWeight()

	if !parent.IsBalanced() {
		t.Error("parent should be balanced")
	}
}

func TestIsBalanced_UnbalancedChildren(t *testing.T) {
	t.Parallel()

	parent, err := day07b.NewNode("parent", 10)
	if err != nil {
		t.Fatalf("creating parent: %v", err)
	}

	child1, err := day07b.NewNode("child1", 20)
	if err != nil {
		t.Fatalf("creating child 1: %v", err)
	}

	child2, err := day07b.NewNode("child2", 21)
	if err != nil {
		t.Fatalf("creating child 2: %v", err)
	}

	parent.AddChildren(child1)
	parent.AddChildren(child2)
	parent.TotalWeight()

	if parent.IsBalanced() {
		t.Error("parent should be unbalanced")
	}
}
