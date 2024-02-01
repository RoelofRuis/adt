package ds

import (
	"fmt"
	"strings"
	"testing"
)

func TestPartitioningTree(t *testing.T) {
	pt := NewPartitioningTree()

	pt.Insert(Interval{0, 1})

	var sb strings.Builder
	pt.TraverseInOrder(func(i Interval) {
		fmt.Fprint(&sb, i, " ")
	})
	got := strings.TrimSpace(sb.String())
	fmt.Printf("%v\n", got)

}
