package ds

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkHeap_Push(b *testing.B) {
	for _, size := range []int{1000, 2000, 4000, 8000, 16000, 32000} {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			heap := NewHeap[int](CompareInt)
			for j := 0; j < size; j++ {
				heap.Push(rand.Int())
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				heap.Push(i)
			}
		})
	}
}

func BenchmarkHeap_Pop(b *testing.B) {
	for _, size := range []int{1000, 2000, 4000, 8000, 16000, 32000} {
		b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
			heap := NewHeap[int](CompareInt)
			for j := 0; j < size; j++ {
				heap.Push(rand.Int())
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				heap.Pop()
			}
		})
	}
}
