package ds

// Stack is a non thread-safe LIFO structure implementation.
type Stack[A any] []A

// NewStack creates a new stack.
func NewStack[A any]() *Stack[A] {
	return &Stack[A]{}
}

// Clone creates a shallow copy of the stack.
func (s *Stack[A]) Clone() *Stack[A] {
	clone := make(Stack[A], len(*s))
	copy(clone, *s)
	return &clone
}

// Push adds an item to the top of the stack.
func (s *Stack[A]) Push(item A) {
	*s = append(*s, item)
}

// Pop removes and returns the item from the top of the stack.
// It returns false if the stack is empty.
func (s *Stack[A]) Pop() (A, bool) {
	if s.IsEmpty() {
		return zeroValue[A](), false
	}

	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item, true
}

// Peek returns the item from the top of the stack without removing it.
// It returns false if the stack is empty.
func (s *Stack[A]) Peek() (A, bool) {
	if s.IsEmpty() {
		return zeroValue[A](), false
	}

	return (*s)[len(*s)-1], true
}

// Size returns the number of elements in the stack.
func (s *Stack[A]) Size() int {
	return len(*s)
}

// IsEmpty returns whether the stack contains no elements.
func (s *Stack[A]) IsEmpty() bool {
	return s.Size() == 0
}
