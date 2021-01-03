package main

// Node represents a program.
type Node struct {
	ID       string
	Weight   int
	Parent   *Node
	children map[string]*Node
}

// AddChildren adds or update a child of n.
func (n *Node) AddChildren(child *Node) {
	n.children[child.ID] = child
	child.Parent = n
}
