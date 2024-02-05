package ds

import (
	"testing"
)

// TestInsertionAndProperties will test the insertion of items and confirm the tree properties remain intact
func TestInsertionAndProperties(t *testing.T) {
	tree := NewRedBlackTree(CompareInt)
	valuesToInsert := []int{10, 18, 7, 15, 16, 30, 25, 40, 60, 2, 1, 70}

	for _, val := range valuesToInsert {
		tree.Insert(val)

		// After each insertion, we verify the properties of the red-black tree
		if _, ok := verifyRBTProperties(t, tree.Root, true); !ok {
			t.Fatalf("Red-black tree properties violated after inserting %v", val)
		}
	}
}

// TestInOrderTraversal asserts that the in-order traversal retrieves elements in sorted order
func TestInOrderTraversal(t *testing.T) {
	tree := NewRedBlackTree(CompareInt)
	valuesToInsert := []int{10, 18, 7, 15, 16, 30, 25, 40, 60, 2, 1, 70}

	for _, val := range valuesToInsert {
		tree.Insert(val)
	}

	expectedOrder := []int{1, 2, 7, 10, 15, 16, 18, 25, 30, 40, 60, 70}
	index := 0

	tree.TraverseInOrder(func(val int) {
		if val != expectedOrder[index] {
			t.Fatalf("InOrder traversal failed. Expected %v got %v", expectedOrder[index], val)
		}
		index++
	})

	if index != len(valuesToInsert) {
		t.Fatalf("InOrder traversal failed. Expected %v values, got %v", len(valuesToInsert), index)
	}
}

// verifyRBTProperties recursively verifies the red-black tree properties
func verifyRBTProperties[A any](t *testing.T, node *RedBlackTreeNode[A], isRoot bool) (int, bool) {
	if node == nil {
		return 1, true // Leafs are considered black, and tree height here considered as 1
	}

	// Check the red node having black children property
	if node.isRed() {
		if (node.Left != nil && node.Left.isRed()) || (node.Right != nil && node.Right.isRed()) {
			t.Errorf("Red node %v has red child", node.Value)
			return 0, false
		}
	}

	leftBlackHeight, leftOk := verifyRBTProperties(t, node.Left, false)
	rightBlackHeight, rightOk := verifyRBTProperties(t, node.Right, false)

	// Check black height property
	if leftBlackHeight != rightBlackHeight || !leftOk || !rightOk {
		t.Errorf("Black height invariant violated at node %v", node.Value)
		return 0, false
	}

	// If the node is black, increment black count for the path
	if !node.isRed() {
		leftBlackHeight++
	}

	// Check the root node color property
	if isRoot && node.isRed() {
		t.Errorf("Root is not black")
		return 0, false
	}

	return leftBlackHeight, true
}
