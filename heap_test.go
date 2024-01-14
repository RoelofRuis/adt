package ds

import (
	"fmt"
	"testing"
)

// intComparator defines a comparator for integers that creates a min-heap.
func intComparator(a, b int) bool {
	return a < b
}

func TestHeap(t *testing.T) {
	h := NewHeap(intComparator)

	// Test Push
	h.Push(4)
	h.Push(1)
	h.Push(3)

	if h.Size() != 3 {
		t.Errorf("Expected heap size 3, got %d", h.Size())
	}

	fmt.Printf("%+v\n", h)

	// Test Pop
	val, ok := h.Pop()
	if !ok || val != 1 { // Min-heap should pop smallest element first
		t.Errorf("Expected 1 got %v, ok: %v", val, ok)
	}

	fmt.Printf("%+v\n", h)

	val, ok = h.Pop()
	if !ok || val != 3 {
		t.Errorf("Expected 3 got %v, ok: %v", val, ok)
	}

	// Test Peek
	h.Push(2)
	topVal, ok := h.Peek()
	if !ok || topVal != 2 {
		t.Errorf("Expected 2 got %v, ok: %v", topVal, ok)
	}

	// Test Pop on last element
	h.Pop() // Now only one element is left

	val, ok = h.Pop()
	if !ok || val != 4 {
		t.Errorf("Expected 4 got %v, ok: %v", val, ok)
	}

	// Test Pop on empty heap
	val, ok = h.Pop()
	if ok {
		t.Errorf("Expected ok to be false, but got %v with value %v", ok, val)
	}

	// Test Peek on empty heap
	topVal, ok = h.Peek()
	if ok {
		t.Errorf("Expected ok to be false, but got %v with value %v", ok, topVal)
	}
}
