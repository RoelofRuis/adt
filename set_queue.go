package ds

// SetQueue is a non thread-safe queue implementation that only queues unique elements.
type SetQueue[T comparable] struct {
	Elements []T
	set      Set[T]
}

// NewSetQueue creates a new SetQueue.
func NewSetQueue[T comparable]() *SetQueue[T] {
	return &SetQueue[T]{set: make(Set[T])}
}

func (q *SetQueue[T]) Clone() *SetQueue[T] {
	clone := NewSetQueue[T]()
	for _, elem := range q.Elements {
		clone.Enqueue(elem)
	}
	return clone
}

// Length return the number of elements in the set queue.
func (q *SetQueue[T]) Length() int {
	return len(q.Elements)
}

// Reset resets the set queue to an empty state.
func (q *SetQueue[T]) Reset() {
	q.Elements = []T{}
	q.set = make(Set[T])
}

// Enqueue adds unique elements to the set queue.
func (q *SetQueue[T]) Enqueue(elems ...T) {
	for _, elem := range elems {
		if q.set.Contains(elem) {
			continue
		}
		q.set.Insert(elem)
		q.Elements = append(q.Elements, elem)
	}
}

// IsEmpty checks if the set queue is empty.
func (q *SetQueue[T]) IsEmpty() bool {
	return len(q.Elements) == 0
}

// Peek returns the first element in the set queue without removing it.
// It panics with an empty queue; use Length to check the length first.
func (q *SetQueue[T]) Peek() T {
	if q.IsEmpty() {
		panic("empty SetQueue")
	}
	return q.Elements[0]
}

// Dequeue removes and returns the first element in the set queue.
// It panics with an empty queue; use Length to check the length first.
func (q *SetQueue[T]) Dequeue() T {
	element := q.Peek()
	delete(q.set, element)
	if len(q.Elements) == 1 {
		q.Elements = nil
	} else {
		q.Elements = q.Elements[1:]
	}
	return element
}
