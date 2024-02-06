package ds

type Interval struct {
	Start int
	End   int
}

type ConterminousIntervalNode struct {
	Interval Interval
	Left     *ConterminousIntervalNode
	Right    *ConterminousIntervalNode
}

type ConterminousIntervalTree struct {
	Root *ConterminousIntervalNode
}

func NewConterminousIntervalTree() *ConterminousIntervalTree {
	return &ConterminousIntervalTree{}
}

func (t *ConterminousIntervalTree) Insert(i Interval) bool {
	newNode := &ConterminousIntervalNode{Interval: i}

	if t.Root == nil {
		t.Root = newNode
		return true
	}

	currentNode := t.Root
	for {
		if i.Start <= currentNode.Interval.Start {
			if i.End > currentNode.Interval.Start {
				return false // overlap
			}
			if currentNode.Left == nil {
				currentNode.Left = newNode
				return true
			}
			currentNode = currentNode.Left
		} else if i.Start >= currentNode.Interval.End {
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

func (t *ConterminousIntervalTree) TraverseInOrder(f func(interval Interval)) {
	var interval *Interval
	var inOrder func(node *ConterminousIntervalNode)
	inOrder = func(node *ConterminousIntervalNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		if interval != nil && interval.End < node.Interval.Start {
			f(Interval{interval.End, node.Interval.Start})
		}
		f(node.Interval)
		interval = &node.Interval
		inOrder(node.Right)
	}

	inOrder(t.Root)
}
