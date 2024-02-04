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
	Left     *PartitioningNode
	Right    *PartitioningNode
}

type PartitioningTree struct {
	Root *PartitioningNode
}

func NewPartitioningTree() *PartitioningTree {
	return &PartitioningTree{}
}

func (t *PartitioningTree) Insert(i Interval) bool {
	newNode := &PartitioningNode{Interval: i}

	if t.Root == nil {
		t.Root = newNode
		return true
	}

	currentNode := t.Root
	for {
		if i.Start >= currentNode.Interval.End {
			// It can be placed after
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return true
			}
			currentNode = currentNode.Right
			continue
		}
		if i.End <= currentNode.Interval.Start {
			// It can be placed before
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return true
			}
			currentNode = currentNode.Left
			continue
		}

		return false
	}
}

func (t *PartitioningTree) TraverseInOrder(f func(i Interval)) {
	var inOrder func(node *PartitioningNode)
	inOrder = func(node *PartitioningNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		f(node.Interval)
		inOrder(node.Right)
	}

	inOrder(t.Root)
}
