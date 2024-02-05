package ds

import (
	"fmt"
	"strings"
	"testing"
)

// TestBinaryTree_Insert tests the Insert method to ensure it properly adds elements.
func TestBinaryTree_Insert(t *testing.T) {
	bt := NewBinarySearchTree(CompareInt)

	values := []int{10, 5, 15, 10, 5, 15}
	for _, v := range values {
		bt.Insert(v)
	}

	// Convert the tree to a string to compare.
	var sb strings.Builder
	bt.TraverseInOrder(func(value int) {
		fmt.Fprint(&sb, value, " ")
	})
	got := strings.TrimSpace(sb.String())
	want := "5 5 10 10 15 15"

	if got != want {
		t.Errorf("InOrderTraversal after Insert got = %v, want %v", got, want)
	}
}

// TestBinaryTree_InOrderTraversal tests the InOrderTraversal method by comparing the output to a known sorted order.
func TestBinaryTree_InOrderTraversal(t *testing.T) {
	bt := NewBinarySearchTree(CompareInt)

	values := []int{5, 3, 8, 1, 4, 7, 9}
	for _, v := range values {
		bt.Insert(v)
	}

	got := make([]int, 0, len(values))
	want := []int{1, 3, 4, 5, 7, 8, 9}

	bt.TraverseInOrder(func(value int) {
		got = append(got, value)
	})

	if len(got) != len(want) {
		t.Fatalf("InOrderTraversal got list of size %d, want size %d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("InOrderTraversal at index %d got = %v, want %v", i, got[i], want[i])
		}
	}
}
