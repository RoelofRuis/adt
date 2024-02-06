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
			"{1 100}",
		},
		{
			"zero width intervals stack",
			[]Interval[int]{{1, 1}, {1, 1}},
			"{1 1} {1 1}",
		},
		{
			"zero width interval cannot intersect interval",
			[]Interval[int]{{1, 3}, {2, 2}, {3, 3}, {1, 1}},
			"{1 1} {1 3} {3 3}",
		},
		{
			"interval cannot overlap zero width interval",
			[]Interval[int]{{2, 2}, {3, 3}, {1, 1}, {1, 3}},
			"{1 1} {1 2} {2 2} {2 3} {3 3}",
		},
		{
			"contiguous traversal: intervals are added between",
			[]Interval[int]{{9, 10}, {1, 4}},
			"{1 4} {4 9} {9 10}",
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
			"{5 10}",
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
			"{1 2} {2 3} {3 4} {4 6} {6 9} {9 10} {10 12} {12 14}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := NewContiguousIntervalTree[int](CompareInt)

			for _, v := range tt.values {
				it.Insert(v)
			}

			var sb strings.Builder
			it.TraverseInOrder(func(value Interval[int]) {
				fmt.Fprint(&sb, value, " ")
			})
			got := strings.TrimSpace(sb.String())

			if got != tt.expected {
				t.Errorf("InOrderTraversal after Insert got = %v, expected %v", got, tt.expected)
			}
		})
	}
}
