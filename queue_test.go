package ds

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	if !q.IsEmpty() {
		t.Errorf("NewQueue() should be empty")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := NewQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2, 3)

	if size := q.Size(); size != 3 {
		t.Errorf("Enqueue() = %d items, want %d", size, 3)
	}

	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("Dequeue() = %d, %t, want %d, %t", val, ok, 1, true)
	}

	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Errorf("Dequeue() = %d, %t, want %d, %t", val, ok, 2, true)
	}

	if size := q.Size(); size != 1 {
		t.Errorf("Size() = %d, want %d", size, 1)
	}

	// Empty the queue completely
	q.Dequeue()

	if !q.IsEmpty() {
		t.Errorf("Queue should be empty after all items are dequeued")
	}

	_, ok = q.Dequeue()
	if ok {
		t.Errorf("Dequeue() on empty queue = %t, want %t", ok, false)
	}
}

func TestPeek(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)

	val, ok := q.Peek()
	if !ok || val != 1 {
		t.Errorf("Peek() = %d, %t, want %d, %t", val, ok, 1, true)
	}

	// Ensure that the size of queue remains the same after Peek
	if size := q.Size(); size != 2 {
		t.Errorf("Size() after Peek() = %d, want %d", size, 2)
	}
}

func TestReset(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)

	q.Reset()

	if !q.IsEmpty() {
		t.Errorf("Reset() should clear all elements in the queue")
	}
}
