package ds

// OrderedSet is a non thread-safe ordered set data structure that maintains the insertion order of items.
type OrderedSet[A comparable] struct {
	set   Set[A]
	items []A
}

// NewOrderedSet creates a new ordered set from the specified values.
func NewOrderedSet[A comparable]() *OrderedSet[A] {
	return &OrderedSet[A]{
		set:   make(Set[A]),
		items: make([]A, 0),
	}
}

// Insert adds an item to the ordered set if it is not already present.
func (os *OrderedSet[A]) Insert(item A) {
	if !os.set.Contains(item) {
		os.set.Insert(item)
		os.items = append(os.items, item)
	}
}

// Remove deletes an item from the ordered set if it exists.
func (os *OrderedSet[A]) Remove(item A) {
	if os.Contains(item) {
		os.set.Delete(item)
		for index, currentItem := range os.items {
			if currentItem == item {
				os.items = append(os.items[:index], os.items[index+1:]...)
				break
			}
		}
	}
}

// Contains returns true if the item is present in the ordered set.
func (os *OrderedSet[A]) Contains(item A) bool {
	return os.set.Contains(item)
}

// Items returns all the items in the ordered set, in the order they were added.
func (os *OrderedSet[A]) Items() []A {
	return os.items
}

// Size returns the number of items in the ordered set.
func (os *OrderedSet[A]) Size() int {
	return os.set.Size()
}

// Clear removes all the items from the ordered set.
func (os *OrderedSet[A]) Clear() {
	os.set = make(Set[A])
	os.items = []A{}
}
