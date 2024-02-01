package ds

import "fmt"

type Interval struct {
	Start, End int
}

func (i Interval) String() string {
	return fmt.Sprintf("[%d %d]", i.Start, i.End)
}

type PartitioningNode struct {
	Interval Interval
}

type PartitioningTree struct {
	Root *PartitioningNode
}

func NewPartitioningTree() *PartitioningTree {
	return &PartitioningTree{}
}

func (t *PartitioningTree) Insert(i Interval) {
	newNode := &PartitioningNode{Interval: i}

	if t.Root == nil {
		t.Root = newNode
		return
	}

	// test if fits
}

func (t *PartitioningTree) TraverseInOrder(f func(i Interval)) {
	var inOrder func(node *PartitioningNode)
	inOrder = func(node *PartitioningNode) {
		if node == nil {
			return
		}
		f(node.Interval)
	}

	inOrder(t.Root)
}
