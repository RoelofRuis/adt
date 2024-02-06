package ds

import (
	"fmt"
	"strings"
	"testing"
)

func TestConterminousIntervalTree_Insert(t *testing.T) {
	tests := []struct {
		name     string
		values   []Interval
		expected string
	}{
		{
			"empty",
			[]Interval{},
			"",
		},
		{
			"single node",
			[]Interval{{1, 100}},
			"{1 100}",
		},
		{
			"zero width intervals stack",
			[]Interval{{1, 1}, {1, 1}},
			"{1 1} {1 1}",
		},
		{
			"zero width interval cannot intersect interval",
			[]Interval{{1, 3}, {2, 2}, {3, 3}, {1, 1}},
			"{1 1} {1 3} {3 3}",
		},
		{
			"interval cannot overlap zero width interval",
			[]Interval{{2, 2}, {3, 3}, {1, 1}, {1, 3}},
			"{1 1} {1 2} {2 2} {2 3} {3 3}",
		},
		{
			"conterminous traversal: intervals are added between",
			[]Interval{{9, 10}, {1, 4}},
			"{1 4} {4 9} {9 10}",
		},
		{
			"intervals cannot overlap",
			[]Interval{
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
			[]Interval{
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
			it := NewConterminousIntervalTree()

			for _, v := range tt.values {
				it.Insert(v)
			}

			var sb strings.Builder
			it.TraverseInOrder(func(value Interval) {
				fmt.Fprint(&sb, value, " ")
			})
			got := strings.TrimSpace(sb.String())

			if got != tt.expected {
				t.Errorf("InOrderTraversal after Insert got = %v, expected %v", got, tt.expected)
			}
		})
	}
}
