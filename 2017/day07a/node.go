package main

// Node represents a program.
type Node struct {
	ID       string
	Weight   int
	Children NodeSet
}

// NodeSet is a collection of unique Node IDs.
type NodeSet struct {
	m map[string]struct{}
}

// Add adds a new unique ID to the set. Returns an error if the ID already exists in the set.
func (*NodeSet) Add(id string) error {
	return nil
}

// Has returns true if the ID is in the set, or false if it is not.
func (*NodeSet) Has(id string) bool {
	return false
}
