package ds

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if size := stack.Size(); size != 3 {
		t.Errorf("stack.Size() = %d, want %d", size, 3)
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack[int]()
	values := []int{1, 2, 3}

	for _, v := range values {
		stack.Push(v)
	}

	for i := len(values) - 1; i >= 0; i-- {
		popped, ok := stack.Pop()
		if !ok {
			t.Fatal("stack.Pop() returned ok = false, want true")
		}
		if popped != values[i] {
			t.Errorf("stack.Pop() = %d, want %d", popped, values[i])
		}
	}

	if _, ok := stack.Pop(); ok {
		t.Error("stack.Pop() returned ok = true, want false on empty stack")
	}
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack[int]()
	values := []int{1, 2, 3}

	for _, v := range values {
		stack.Push(v)
	}

	peeked, ok := stack.Peek()
	if !ok {
		t.Fatal("stack.Peek() returned ok = false, want true")
	}
	if peeked != values[len(values)-1] {
		t.Errorf("stack.Peek() = %d, want %d", peeked, values[len(values)-1])
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack[int]()
	if !stack.IsEmpty() {
		t.Error("stack.IsEmpty() = false, want true for new stack")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("stack.IsEmpty() = true, want false for stack with elements")
	}
}

func TestStack_Reset(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Reset()

	if !stack.IsEmpty() {
		t.Error("stack.IsEmpty() = false, want true after Reset()")
	}
}

func TestStack_Size(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)

	if size := stack.Size(); size != 2 {
		t.Errorf("stack.Size() = %d, want %d", size, 2)
	}

	stack.Pop()
	if size := stack.Size(); size != 1 {
		t.Errorf("stack.Size() = %d, want %d after Pop()", size, 1)
	}
}
