package day07b

type Node struct {
	id          string
	weight      int
	parent      *Node
	children    map[string]*Node
	totalWeight int
}

// CalculateTotalWeight calculates the total weight of a node, this is, its weight plus the total weight of its children.
// This function requires that the total weight of the childen has been previously calculated.
func (n *Node) CalculateTotalWeight() {
	n.totalWeight = n.weight
	for _, c := range n.children {
		n.totalWeight += c.totalWeight
	}
}

// GetTotalWeight returns the total weight of the node. It only returns correct values if CalculateTotalWeight have been called previously.
func (n *Node) GetTotalWeight() int {
	return n.totalWeight
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
