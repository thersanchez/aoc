package day07a

// Node represents a program.
type Node struct {
	ID       string
	Weight   int
	Parent   *Node
	Children map[string]*Node
}

// NewNode return a new node with the given id and weight.
func NewNode(id string, weight int) *Node {
	return &Node{
		ID:     id,
		Weight: weight,
	}
}

// AddChildren adds or update a child of n.
func (n *Node) AddChildren(child *Node) {
	n.Children[child.ID] = child
	child.Parent = n
}
