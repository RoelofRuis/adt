package ds

// SetQueue is a non thread-safe queue implementation that only queues unique elements.
type SetQueue[T comparable] struct {
	queue *Queue[T]
	set   Set[T]
}

// NewSetQueue creates a new SetQueue.
func NewSetQueue[T comparable]() *SetQueue[T] {
	return &SetQueue[T]{
		queue: NewQueue[T](),
		set:   make(Set[T]),
	}
}

// Clone creates a shallow copy of the SetQueue.
func (q *SetQueue[T]) Clone() *SetQueue[T] {
	clone := NewSetQueue[T]()
	for _, elem := range *q.queue {
		clone.Enqueue(elem)
	}
	return clone
}

// Size return the number of elements in the set queue.
func (q *SetQueue[T]) Size() int {
	return q.queue.Size()
}

// Reset resets the set queue to an empty state.
func (q *SetQueue[T]) Reset() {
	q.queue.Reset()
	q.set = make(Set[T])
}

// Enqueue adds unique elements to the set queue.
func (q *SetQueue[T]) Enqueue(elems ...T) {
	for _, elem := range elems {
		if q.set.Contains(elem) {
			continue
		}
		q.set.Insert(elem)
		q.queue.Enqueue(elem)
	}
}

// IsEmpty checks if the set queue is empty.
func (q *SetQueue[T]) IsEmpty() bool {
	return q.queue.IsEmpty()
}

// Peek returns the first element in the set queue without removing it.
func (q *SetQueue[T]) Peek() (T, bool) {
	return q.queue.Peek()
}

// Dequeue removes and returns the first element in the set queue.
func (q *SetQueue[T]) Dequeue() T {
	elem, has := q.queue.Dequeue()
	if has {
		delete(q.set, elem)
	}
	return elem
}
