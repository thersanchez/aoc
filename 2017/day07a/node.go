package main

// Node represents a program.
type Node struct {
	id       string
	weight   int
	parent   *Node
	children map[string]*Node
}

// NewNode return a new node with the given id and weight.
func NewNode(id string, weight int) *Node {
	return &Node{
		id:     id,
		weight: weight,
	}
}

// AddChildren adds or update a child of n.
func (n *Node) AddChildren(child *Node) {
	n.children[child.id] = child
	child.parent = n
}
