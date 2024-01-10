package ds

import (
	"testing"
)

func TestSetQueue_Enqueue(t *testing.T) {
	sq := NewSetQueue[int]()
	values := []int{1, 2, 3, 2, 1}

	for _, v := range values {
		sq.Enqueue(v)
	}

	expectedSize := 3 // Only 1, 2, and 3 should be in the queue.
	if sq.Size() != expectedSize {
		t.Errorf("expected queue size %d, got %d", expectedSize, sq.Size())
	}
}

func TestSetQueue_Dequeue(t *testing.T) {
	sq := NewSetQueue[int]()
	values := []int{1, 2, 3}

	for _, v := range values {
		sq.Enqueue(v)
	}

	for i := 0; i < len(values); i++ {
		val, ok := sq.Dequeue()
		if !ok || val != values[i] {
			t.Errorf("expected %d, got %d, successful: %v", values[i], val, ok)
		}
	}

	if !sq.IsEmpty() {
		t.Errorf("expected queue to be empty, got size: %d", sq.Size())
	}
}

func TestSetQueue_Peek(t *testing.T) {
	sq := NewSetQueue[int]()
	values := []int{1, 2, 3}

	for _, v := range values {
		sq.Enqueue(v)
	}

	peekedValue, ok := sq.Peek()
	if !ok || peekedValue != values[0] {
		t.Errorf("expected peek to return %d, got %d, successful: %v", values[0], peekedValue, ok)
	}
}

func TestSetQueue_IsEmpty(t *testing.T) {
	sq := NewSetQueue[int]()
	if !sq.IsEmpty() {
		t.Error("expected new queue to be empty")
	}

	sq.Enqueue(1)
	if sq.IsEmpty() {
		t.Error("expected non-empty queue")
	}
}

func TestSetQueue_Size(t *testing.T) {
	sq := NewSetQueue[int]()
	values := []int{1, 2, 3, 4, 5}

	for _, v := range values {
		sq.Enqueue(v)
	}

	if sq.Size() != len(values) {
		t.Errorf("expected size %d, got %d", len(values), sq.Size())
	}
}

func TestSetQueue_Reset(t *testing.T) {
	sq := NewSetQueue[int]()
	values := []int{1, 2, 3}

	for _, v := range values {
		sq.Enqueue(v)
	}

	sq.Reset()
	if !sq.IsEmpty() {
		t.Error("expected empty queue after reset")
	}
}

func TestSetQueue_Clone(t *testing.T) {
	sq := NewSetQueue[int]()
	values := []int{1, 2, 3}

	for _, v := range values {
		sq.Enqueue(v)
	}

	clone := sq.Clone()

	if clone.Size() != sq.Size() {
		t.Errorf("cloned queue size %d, expected %d", clone.Size(), sq.Size())
	}

	for i, _ := range values {
		if origVal, _ := sq.Dequeue(); origVal != i+1 {
			t.Errorf("expected %d, got %d", i+1, origVal)
		}
		if cloneVal, _ := clone.Dequeue(); cloneVal != i+1 {
			t.Errorf("expected %d, got %d in clone", i+1, cloneVal)
		}
	}
}
