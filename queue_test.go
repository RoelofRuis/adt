package ds

import "testing"

func TestQueue(t *testing.T) {
	// Test NewQueue and Size
	queue := NewQueue[int]()
	if size := queue.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test Enqueue and Size
	queue.Enqueue(1)
	if size := queue.Size(); size != 1 {
		t.Errorf("Expected size 1, got %d", size)
	}

	queue.Enqueue(2)
	queue.Enqueue(3)
	if size := queue.Size(); size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}

	// Test Peek
	frontItem, ok := queue.Peek()
	if !ok || frontItem != 1 {
		t.Errorf("Expected Front item 1, got %v (ok: %t)", frontItem, ok)
	}
	if size := queue.Size(); size != 3 {
		t.Errorf("Expected size 3 after Front, got %d", size)
	}

	// Test Dequeue and Size
	dequeuedItem, ok := queue.Dequeue()
	if !ok || dequeuedItem != 1 {
		t.Errorf("Expected Dequeued item 1, got %v (ok: %t)", dequeuedItem, ok)
	}
	if size := queue.Size(); size != 2 {
		t.Errorf("Expected size 2 after Dequeue, got %d", size)
	}

	// Test IsEmpty
	empty := queue.IsEmpty()
	if empty {
		t.Errorf("Expected IsEmpty=false, got IsEmpty=true")
	}

	// Test Dequeue on an empty queue
	emptyQueue := NewQueue[int]()
	_, ok = emptyQueue.Dequeue()
	if ok {
		t.Errorf("Expected ok=false on Dequeue from an empty queue, got ok=true")
	}

	// Test Clone
	original := NewQueue[int]()
	original.Enqueue(1)
	original.Enqueue(2)
	original.Enqueue(3)
	cloned := original.Clone()
	if size := cloned.Size(); size != 3 {
		t.Errorf("Expected size 3 for cloned queue, got %d", size)
	}
}
