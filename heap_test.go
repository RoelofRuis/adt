package ds

import (
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap(CompareInt) // creates a min-heap.

	// Test Push
	h.Push(4)
	h.Push(1)
	h.Push(3)
	h.Push(1)

	if h.Size() != 4 {
		t.Errorf("Expected heap size 4, got %d", h.Size())
	}

	// Test Pop
	val, ok := h.Pop()
	if !ok || val != 1 { // Min-heap should pop smallest element first
		t.Errorf("Expected 1 got %v, ok: %v", val, ok)
	}

	val, ok = h.Pop()
	if !ok || val != 1 { // Min-heap should pop smallest element first
		t.Errorf("Expected 1 got %v, ok: %v", val, ok)
	}

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
