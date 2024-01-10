package ds

// Queue is a non thread-safe queue (FIFO) implementation.
type Queue[A any] []A

// NewQueue creates a new Queue
func NewQueue[A any]() *Queue[A] {
	return &Queue[A]{}
}

func (q *Queue[A]) Clone() *Queue[A] {
	clone := make(Queue[A], len(*q))
	copy(clone, *q)
	return &clone
}

// Size returns the number of elements in the queue.
func (q *Queue[A]) Size() int {
	return len(*q)
}

// IsEmpty returns whether the queue is empty.
func (q *Queue[A]) IsEmpty() bool {
	return q.Size() == 0
}

// Reset resets the queue to an empty state.
func (q *Queue[A]) Reset() {
	*q = (*q)[:0]
}

// Enqueue adds elements to the back of the queue.
func (q *Queue[A]) Enqueue(elems ...A) {
	*q = append(*q, elems...)
}

// Dequeue removes and returns the item from the front of the queue.
// It returns false if the queue is empty.
func (q *Queue[A]) Dequeue() (A, bool) {
	if q.IsEmpty() {
		return zeroValue[A](), false
	}

	item := (*q)[0]
	(*q)[0] = zeroValue[A]()
	*q = (*q)[1:]
	return item, true
}

// Peek returns the first element in the queue without removing it.
// It returns false if the queue is empty.
func (q *Queue[A]) Peek() (A, bool) {
	if q.IsEmpty() {
		return zeroValue[A](), false
	}

	return (*q)[0], true
}
