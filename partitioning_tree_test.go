package ds

import (
	"fmt"
	"strings"
	"testing"
)

func TestPartitioningTree(t *testing.T) {
	pt := NewPartitioningTree()

	pt.Insert(Interval{4, 6})
	pt.Insert(Interval{0, 4})
	pt.Insert(Interval{9, 11})
	pt.Insert(Interval{6, 7})
	pt.Insert(Interval{7, 8})
	pt.Insert(Interval{1, 2})

	debug(pt)
	var sb strings.Builder
	pt.TraverseInOrder(func(i Interval) {
		fmt.Fprint(&sb, i, " ")
	})
	got := strings.TrimSpace(sb.String())
	fmt.Printf("%v\n", got)
}

func debug(t *PartitioningTree) {
	debugTree(t.Root, 0)
}

func debugTree(n *PartitioningNode, depth int) {
	if n.Left != nil {
		debugTree(n.Left, depth+1)
	}
	fmt.Printf("%s%v\n", strings.Repeat(" ", depth), n.Interval)
	if n.Right != nil {
		debugTree(n.Right, depth+1)
	}
}
