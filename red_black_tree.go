package ds

// Red and Black define the color representation for the nodes.
const (
	RED   = true
	BLACK = false
)

// RedBlackTreeNode represents a node in a red-black tree.
type RedBlackTreeNode[K any, V any] struct {
	Key    K
	Value  V
	Color  bool
	Left   *RedBlackTreeNode[K, V]
	Right  *RedBlackTreeNode[K, V]
	Parent *RedBlackTreeNode[K, V]
}

// isRed returns whether the node color is red.
func (n *RedBlackTreeNode[K, V]) isRed() bool {
	if n == nil {
		return BLACK
	}
	return n.Color == RED
}

// RedBlackTree represents a non thread-safe generic red-black tree.
type RedBlackTree[K any, V any] struct {
	Root       *RedBlackTreeNode[K, V]
	Comparator Comparator[K]
}

// NewRedBlackTree creates a new red-black tree that uses the given Comparator.
func NewRedBlackTree[K any, V any](comparator Comparator[K]) *RedBlackTree[K, V] {
	return &RedBlackTree[K, V]{
		Root:       nil,
		Comparator: comparator,
	}
}

// Insert inserts a new value into the red-black tree.
func (t *RedBlackTree[K, V]) Insert(key K, value V) {
	newNode := &RedBlackTreeNode[K, V]{Key: key, Value: value, Color: RED}

	if t.Root == nil {
		t.Root = newNode
	} else {
		current := t.Root
		var parent *RedBlackTreeNode[K, V]

		for current != nil {
			parent = current
			compResult := t.Comparator(key, current.Key)
			if compResult == 0 {
				return
			} else if compResult < 0 {
				current = current.Left
			} else {
				current = current.Right
			}
		}

		newNode.Parent = parent
		if t.Comparator(key, parent.Key) < 0 {
			parent.Left = newNode
		} else {
			parent.Right = newNode
		}
	}

	t.insertFixup(newNode)
}

// TraverseInOrder performs an in-order traversal of the tree, applying the given function to each node.
func (t *RedBlackTree[K, V]) TraverseInOrder(f func(K)) {
	var inOrder func(node *RedBlackTreeNode[K, V])
	inOrder = func(node *RedBlackTreeNode[K, V]) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		f(node.Key)
		inOrder(node.Right)
	}

	inOrder(t.Root)
}

// rotateLeft performs a left rotation around the given node.
func (t *RedBlackTree[K, V]) rotateLeft(x *RedBlackTreeNode[K, V]) {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// rotateRight performs a right rotation around the given node.
func (t *RedBlackTree[K, V]) rotateRight(x *RedBlackTreeNode[K, V]) {
	y := x.Left
	x.Left = y.Right
	if y.Right != nil {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}
	y.Right = x
	x.Parent = y
}

// insertFixup fixes up the tree after an insertion to maintain red-black properties.
func (t *RedBlackTree[K, V]) insertFixup(z *RedBlackTreeNode[K, V]) {
	for z.Parent.isRed() {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y.isRed() {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					t.rotateLeft(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.rotateRight(z.Parent.Parent)
			}
		} else {
			y := z.Parent.Parent.Left
			if y.isRed() {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					t.rotateRight(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.rotateLeft(z.Parent.Parent)
			}
		}
	}
	t.Root.Color = BLACK
}
