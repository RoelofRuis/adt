package ds

type Interval[A any] struct {
	Start A
	End   A
}

type ContiguousIntervalNode[A any] struct {
	Interval Interval[A]
	Left     *ContiguousIntervalNode[A]
	Right    *ContiguousIntervalNode[A]
}

type ContiguousIntervalTree[A any] struct {
	Comparator Comparator[A]
	Root       *ContiguousIntervalNode[A]
}

func NewContiguousIntervalTree[A any](comparator Comparator[A]) *ContiguousIntervalTree[A] {
	return &ContiguousIntervalTree[A]{
		Root:       nil,
		Comparator: comparator,
	}
}

func (t *ContiguousIntervalTree[A]) Insert(i Interval[A]) bool {
	newNode := &ContiguousIntervalNode[A]{Interval: i}

	if t.Root == nil {
		t.Root = newNode
		return true
	}

	currentNode := t.Root
	for {

		if t.Comparator(i.Start, currentNode.Interval.Start) <= 0 {
			if t.Comparator(i.End, currentNode.Interval.Start) > 0 {
				return false // overlap
			}
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return true
			}
			currentNode = currentNode.Left
		} else if t.Comparator(i.Start, currentNode.Interval.End) >= 0 {
			if currentNode.Right == nil {
				currentNode.Right = newNode
				return true
			}
			currentNode = currentNode.Right
		} else {
			return false // overlap
		}
	}
}

func (t *ContiguousIntervalTree[A]) TraverseInOrder(f func(interval Interval[A])) {
	var lastInterval *Interval[A]

	var inOrder func(node *ContiguousIntervalNode[A])
	inOrder = func(node *ContiguousIntervalNode[A]) {
		if node == nil {
			return
		}
		inOrder(node.Left)

		if lastInterval != nil && t.Comparator(lastInterval.End, node.Interval.Start) < 0 {
			f(Interval[A]{lastInterval.End, node.Interval.Start})
		}

		f(node.Interval)

		lastInterval = &node.Interval
		inOrder(node.Right)
	}

	inOrder(t.Root)
}
