package ds

// TreeNode represents a node in the Scapegoat Tree.
type ScapegoatTreeNode struct {
	Value int
	Left  *ScapegoatTreeNode
	Right *ScapegoatTreeNode
}

// ScapegoatTree represents a Scapegoat Tree data structure.
type ScapegoatTree struct {
	Root    *ScapegoatTreeNode
	Size    int
	MaxSize int
	Alpha   float64
}

// NewScapegoatTree creates a new ScapegoatTree with the given alpha.
func NewScapegoatTree(alpha float64) *ScapegoatTree {
	return &ScapegoatTree{
		Alpha: alpha,
	}
}

// insertHelper is the helper for the Insert operation that also identifies
// if the tree is unbalanced.
func (st *ScapegoatTree) insertHelper(value int, node *ScapegoatTreeNode) (*ScapegoatTreeNode, bool) {
	// Base case: Found insertion point
	if node == nil {
		st.Size++
		return &ScapegoatTreeNode{Value: value}, false // return unbalanced as false because a new node is always balanced
	}

	var unbalanced bool
	if value < node.Value {
		node.Left, unbalanced = st.insertHelper(value, node.Left)
	} else if value > node.Value {
		node.Right, unbalanced = st.insertHelper(value, node.Right)
	}

	// Check balance of the current node
	if !unbalanced {
		leftSize := st.nodeSize(node.Left)
		rightSize := st.nodeSize(node.Right)

		// Check for scapegoat condition
		if float64(leftSize+rightSize+1) > 1/st.Alpha {
			if float64(leftSize) > st.Alpha*float64(leftSize+rightSize+1) || float64(rightSize) > st.Alpha*float64(leftSize+rightSize+1) {
				unbalanced = true
			}
		}
	}
	return node, unbalanced
}

// Insert adds a value to the Scapegoat Tree.
func (st *ScapegoatTree) Insert(value int) {
	var unbalanced bool
	st.Root, unbalanced = st.insertHelper(value, st.Root)
	if unbalanced {
		// Update maximum tree size if needed
		if st.Size > st.MaxSize {
			st.MaxSize = st.Size
		}
		// Rebuild the tree starting from the root, since unbalance may have occurred
		st.Root = st.rebuildTree(st.Root)
	}
}

// rebuildTree rebuilds the given subtree into a balanced tree.
func (st *ScapegoatTree) rebuildTree(node *ScapegoatTreeNode) *ScapegoatTreeNode {
	nodes := st.flatten(node)
	return st.buildBalancedTree(nodes)
}

// flatten performs an in-order traversal of the tree rooted at node,
// flattening the tree into a slice.
func (st *ScapegoatTree) flatten(node *ScapegoatTreeNode) []*ScapegoatTreeNode {
	if node == nil {
		return nil
	}
	nodes := st.flatten(node.Left)
	nodes = append(nodes, node)
	nodes = append(nodes, st.flatten(node.Right)...)
	node.Left = nil
	node.Right = nil
	return nodes
}

// buildBalancedTree builds a balanced tree from a sorted slice of nodes.
func (st *ScapegoatTree) buildBalancedTree(nodes []*ScapegoatTreeNode) *ScapegoatTreeNode {
	if len(nodes) == 0 {
		return nil
	}
	mid := len(nodes) / 2
	node := nodes[mid]
	node.Left = st.buildBalancedTree(nodes[:mid])
	node.Right = st.buildBalancedTree(nodes[mid+1:])
	return node
}

// nodeSize computes the size of the subtree rooted at node.
func (st *ScapegoatTree) nodeSize(node *ScapegoatTreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + st.nodeSize(node.Left) + st.nodeSize(node.Right)
}

// InOrderTraversal recursively traverses the tree in-order and applies
// the function f to each node's value.
func (st *ScapegoatTree) InOrderTraversal(node *ScapegoatTreeNode, f func(int)) {
	if node == nil {
		return
	}
	st.InOrderTraversal(node.Left, f)
	f(node.Value)
	st.InOrderTraversal(node.Right, f)
}
