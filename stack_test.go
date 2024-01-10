package ds

import "testing"

func TestStack(t *testing.T) {
	// Test NewStack and Size
	stack := NewStack[int]()
	if size := stack.Size(); size != 0 {
		t.Errorf("Expected size 0, got %d", size)
	}

	// Test Push and Size
	stack.Push(1)
	if size := stack.Size(); size != 1 {
		t.Errorf("Expected size 1, got %d", size)
	}

	stack.Push(2)
	stack.Push(3)
	if size := stack.Size(); size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}

	// Test Peek
	peekedItem, ok := stack.Peek()
	if !ok || peekedItem != 3 {
		t.Errorf("Expected Peeked item 3, got %v (ok: %t)", peekedItem, ok)
	}
	if size := stack.Size(); size != 3 {
		t.Errorf("Expected size 3 after Peek, got %d", size)
	}

	// Test Pop and Size
	poppedItem, ok := stack.Pop()
	if !ok || poppedItem != 3 {
		t.Errorf("Expected Popped item 3, got %v (ok: %t)", poppedItem, ok)
	}
	if size := stack.Size(); size != 2 {
		t.Errorf("Expected size 2 after Pop, got %d", size)
	}

	// Test IsEmpty
	empty := stack.IsEmpty()
	if empty {
		t.Errorf("Expected IsEmpty=false, got IsEmpty=true")
	}

	// Test Pop on an empty stack
	emptyStack := NewStack[int]()
	_, ok = emptyStack.Pop()
	if ok {
		t.Errorf("Expected ok=false on Pop from an empty stack, got ok=true")
	}

	// Test Clone
	original := NewStack[int]()
	original.Push(1)
	original.Push(2)
	original.Push(3)
	cloned := original.Clone()
	if size := cloned.Size(); size != 3 {
		t.Errorf("Expected size 3 for cloned stack, got %d", size)
	}
}
