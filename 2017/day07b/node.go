package day07b

import (
	"fmt"
)

type Node struct {
	id     string
	weight int

	// as part of a tree, the node also has these other fields
	parent   *Node
	children map[string]*Node
	// totalWeight is the weight of the node plus the total weight of its children.
	// The zero value means it has not yet been calculated. See TotalWeight.
	totalWeight int
}

// NewNode return a new node with the given id and weight.
func NewNode(id string, weight int) (*Node, error) {
	if id == "" {
		return nil, fmt.Errorf("invalid id (%q), empty string", id)
	}

	if weight < 1 {
		return nil, fmt.Errorf("invalid weight (%d), should be > 0", weight)
	}

	return &Node{
		id:       id,
		weight:   weight,
		children: make(map[string]*Node),
	}, nil
}

// Id returns the id of the node.
func (n *Node) Id() string {
	return n.id
}

// Weight returns the weight of the node.
func (n *Node) Weight() int {
	return n.weight
}

// AddChildren adds or update a child of n.
// When you add a child, you invalidate the total weight of its ancestors.
func (n *Node) AddChildren(child *Node) {
	n.children[child.id] = child
	child.parent = n
	n.invalidateTotalWeight()
}

func (n *Node) invalidateTotalWeight() {
	n.totalWeight = 0
	if n.parent != nil {
		n.parent.invalidateTotalWeight()
	}
}

// TotalWeight returns the total weight of a node, this is, its weight plus the total weight of its children.
func (n *Node) TotalWeight() int {
	if n.totalWeight != 0 {
		return n.totalWeight
	}

	n.totalWeight = n.weight
	for _, c := range n.children {
		n.totalWeight += c.TotalWeight()
	}
	return n.totalWeight
}

// Children returns the children of n.
func (n *Node) Children() []*Node {
	c := []*Node{}

	for _, v := range n.children {
		c = append(c, v)
	}

	return c
}

// IsBalanced returns if the children have the same weight.
func (n *Node) IsBalanced() bool {
	children := n.Children()
	if len(children) == 0 {
		return true
	}

	childrenWeights := []int{}
	for _, c := range children {
		childrenWeights = append(childrenWeights, c.TotalWeight())
	}

	for _, n := range childrenWeights {
		if n != childrenWeights[0] {
			return false
		}
	}

	return true
}
