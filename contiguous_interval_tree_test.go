package ds

import (
	"fmt"
	"strings"
	"testing"
)

func TestContiguousIntervalTree_Insert(t *testing.T) {
	tests := []struct {
		name     string
		values   []Interval[int]
		expected string
	}{
		{
			"empty",
			[]Interval[int]{},
			"",
		},
		{
			"single node",
			[]Interval[int]{{1, 100}},
			"{1 100}:data",
		},
		{
			"zero width intervals stack",
			[]Interval[int]{{1, 1}, {1, 1}},
			"{1 1}:data {1 1}:data",
		},
		{
			"zero width interval cannot intersect interval",
			[]Interval[int]{{1, 3}, {2, 2}, {3, 3}, {1, 1}},
			"{1 1}:data {1 3}:data {3 3}:data",
		},
		{
			"interval cannot overlap zero width interval",
			[]Interval[int]{{2, 2}, {3, 3}, {1, 1}, {1, 3}},
			"{1 1}:data {1 2}: {2 2}:data {2 3}: {3 3}:data",
		},
		{
			"contiguous traversal: intervals are added between",
			[]Interval[int]{{9, 10}, {1, 4}},
			"{1 4}:data {4 9}: {9 10}:data",
		},
		{
			"intervals cannot overlap",
			[]Interval[int]{
				{5, 10},
				{2, 6},
				{8, 14},
				{6, 9},
				{5, 10},
			},
			"{5 10}:data",
		},
		{
			"intervals are traversed in sorted order",
			[]Interval[int]{
				{6, 9},
				{1, 2},
				{3, 4},
				{2, 3},
				{12, 14},
				{10, 12},
			},
			"{1 2}:data {2 3}:data {3 4}:data {4 6}: {6 9}:data {9 10}: {10 12}:data {12 14}:data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := NewContiguousIntervalTree[int, string](CompareInt)

			for _, v := range tt.values {
				it.Insert(v, "data")
			}

			var sb strings.Builder
			it.TraverseInOrder(func(value Interval[int], s string) {
				fmt.Fprintf(&sb, "%v:%s ", value, s)
			})
			got := strings.TrimSpace(sb.String())

			if got != tt.expected {
				t.Errorf("InOrderTraversal after Insert got = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestContiguousIntervalTree_SizeAndNumIntervals(t *testing.T) {
	it := NewContiguousIntervalTree[int, string](CompareInt)

	if it.Size() != 0 {
		t.Errorf("expected tree of size 0, got %d", it.Size())
	}
	if it.NumIntervals() != 0 {
		t.Errorf("expected 0 intervals, got %d", it.NumIntervals())
	}

	// insert single interval
	it.Insert(Interval[int]{1, 4}, "data")

	if it.Size() != 1 {
		t.Errorf("expected tree of size 1, got %d", it.Size())
	}
	if it.NumIntervals() != 1 {
		t.Errorf("expected 1 interval, got %d", it.NumIntervals())
	}

	// insert non-touching interval
	it.Insert(Interval[int]{8, 12}, "data")

	if it.Size() != 2 {
		t.Errorf("expected tree of size 2, got %d", it.Size())
	}
	if it.NumIntervals() != 3 {
		t.Errorf("expected 3 intervals, got %d", it.NumIntervals())
	}

	// insert touching interval
	it.Insert(Interval[int]{12, 14}, "data")

	if it.Size() != 3 {
		t.Errorf("expected tree of size 3, got %d", it.Size())
	}
	if it.NumIntervals() != 4 {
		t.Errorf("expected 4 intervals, got %d", it.NumIntervals())
	}

	// insert filling interval
	it.Insert(Interval[int]{4, 8}, "data")

	if it.Size() != 4 {
		t.Errorf("expected tree of size 4, got %d", it.Size())
	}
	if it.NumIntervals() != 4 {
		t.Errorf("expected 4 intervals, got %d", it.NumIntervals())
	}

	// insert overlapping interval
	it.Insert(Interval[int]{6, 10}, "data")

	if it.Size() != 4 {
		t.Errorf("expected tree of size 4, got %d", it.Size())
	}
	if it.NumIntervals() != 4 {
		t.Errorf("expected 4 intervals, got %d", it.NumIntervals())
	}
}

func TestContiguousIntervalTree_Rebalance(t *testing.T) {
	it := NewContiguousIntervalTree[int, string](CompareInt)

	it.Insert(Interval[int]{1, 2}, "data")
	it.Insert(Interval[int]{2, 3}, "data")
	it.Insert(Interval[int]{3, 4}, "data")
	it.Insert(Interval[int]{4, 5}, "data")
	it.Insert(Interval[int]{5, 6}, "data")
	it.Insert(Interval[int]{6, 7}, "data")

	showTree(it.Root, 0)
}

func showTree(n *ContiguousIntervalNode[int, string], depth int) {
	if n == nil {
		return
	}
	showTree(n.Left, depth+1)
	fmt.Printf("%s%v (%v)\n", strings.Repeat("  ", depth), n.Interval, n.alphaBalanced(0.5))
	showTree(n.Right, depth+1)
}
