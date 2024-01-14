package ds

// Comparator is a function that defines the ordering of elements in the heap.
// It returns true if the first element should be ordered before the second.
type Comparator[A any] func(a, b A) bool

// Heap is a generic heap data structur containing elements of any type A.
// Heap ordering is maintained using the given Comparator.
type Heap[A any] struct {
	elements   []A
	comparator Comparator[A]
}

// NewHeap creates a new Heap instance with the provided Comparator.
func NewHeap[A any](comparator Comparator[A]) *Heap[A] {
	return &Heap[A]{
		elements:   make([]A, 0),
		comparator: comparator,
	}
}

// Push adds an element to the heap while maintaining the heap invariant.
func (h *Heap[A]) Push(element A) {
	h.elements = append(h.elements, element)
	h.up(len(h.elements) - 1)
}

// Pop removes and returns the highest-priority element from the heap according to the comparator.
// If the heap is empty, Pop returns the zero value of type A and false.
func (h *Heap[A]) Pop() (A, bool) {
	if len(h.elements) == 0 {
		return zeroValue[A](), false
	}

	n := len(h.elements) - 1
	h.swap(0, n)
	top := h.elements[n]
	h.elements = h.elements[:n]
	h.down(0)

	return top, true
}

// Peek returns the highest-priority element from the heap without removing it.
// If the heap is empty, Peek returns the zero value of type A and false.
func (h *Heap[A]) Peek() (A, bool) {
	if len(h.elements) == 0 {
		return zeroValue[A](), false
	}

	return h.elements[0], true
}

// Size returns the number of elements in the heap.
func (h *Heap[A]) Size() int {
	return len(h.elements)
}

// up moves the element at the given index up the heap to its correct position to maintain the heap invariant.
func (h *Heap[A]) up(index int) {
	for {
		parentIndex := (index - 1) / 2
		if index == 0 || h.comparator(h.elements[parentIndex], h.elements[index]) {
			break
		}
		h.swap(index, parentIndex)
		index = parentIndex
	}
}

// down moves the element at the given index down the heap to its correct position to maintain the heap invariant.
func (h *Heap[A]) down(index int) {
	n := len(h.elements)
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		smallestIndex := index
		if leftChildIndex < n && h.comparator(h.elements[leftChildIndex], h.elements[smallestIndex]) {
			smallestIndex = leftChildIndex
		}
		if rightChildIndex < n && h.comparator(h.elements[rightChildIndex], h.elements[smallestIndex]) {
			smallestIndex = rightChildIndex
		}

		if smallestIndex == index {
			break
		}

		h.swap(index, smallestIndex)
		index = smallestIndex
	}
}

// swap exchanges the elements at the given indices.
func (h *Heap[A]) swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}
