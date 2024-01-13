package ds

import (
	"fmt"
	"testing"
)

func BenchmarkSet_Insert(b *testing.B) {
	for _, size := range []int{100, 200, 400, 800, 1600, 3200} {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				set := NewSet[int]()
				for j := 0; j < size; j++ {
					set.Insert(j)
				}
			}
		})
	}
}
