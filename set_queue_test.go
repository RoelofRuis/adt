package ds

import (
	"reflect"
	"testing"
)

func TestSetQueue_NewSetQueue(t *testing.T) {
	queue := NewSetQueue[int]()
	if !queue.IsEmpty() {
		t.Error("Newly created SetQueue should be empty.")
	}
}

func TestSetQueue_EnqueueAndLength(t *testing.T) {
	queue := NewSetQueue[int]()
	queue.Enqueue(1, 2, 3)
	if length := queue.Length(); length != 3 {
		t.Errorf("Length is incorrect after Enqueue. Expected: 3, Got: %d", length)
	}
}

func TestSetQueue_Peek(t *testing.T) {
	queue := NewSetQueue[int]()
	queue.Enqueue(1, 2, 3)
	elem := queue.Peek()
	if elem != 1 {
		t.Errorf("Peek is incorrect. Expected: 1, Got: %v", elem)
	}
}

func TestSetQueue_Dequeue(t *testing.T) {
	queue := NewSetQueue[int]()
	queue.Enqueue(1, 2, 3)
	elem := queue.Dequeue()
	if elem != 1 {
		t.Errorf("Dequeue is incorrect. Expected: 1, Got: %v", elem)
	}
	if length := queue.Length(); length != 2 {
		t.Errorf("Length is incorrect after Dequeue. Expected: 2, Got: %d", length)
	}
}

func TestSetQueue_IsEmptyAfterDequeue(t *testing.T) {
	queue := NewSetQueue[int]()
	queue.Enqueue(1, 2, 3)
	queue.Dequeue()
	if queue.IsEmpty() {
		t.Error("SetQueue should not be empty after Dequeue.")
	}
}

func TestSetQueue_EnqueueWithDuplicates(t *testing.T) {
	queue := NewSetQueue[int]()
	queue.Enqueue(1, 2, 3)
	queue.Enqueue(2, 3, 4)
	if length := queue.Length(); length != 4 {
		t.Errorf("Length is incorrect after Enqueue with duplicates. Expected: 4, Got: %d", length)
	}
}

func TestSetQueue_Clone(t *testing.T) {
	queue := NewSetQueue[int]()
	queue.Enqueue(1, 2, 3)
	clone := queue.Clone()
	if !reflect.DeepEqual(queue.Elements, clone.Elements) {
		t.Error("Clone did not produce an equal SetQueue.")
	}
	if &queue.Elements == &clone.Elements {
		t.Error("Clone should create a new instance of the Elements slice.")
	}
}

func TestSetQueue_Reset(t *testing.T) {
	queue := NewSetQueue[int]()
	queue.Enqueue(1, 2, 3)
	queue.Reset()
	if !queue.IsEmpty() {
		t.Error("SetQueue should be empty after Reset.")
	}
}
