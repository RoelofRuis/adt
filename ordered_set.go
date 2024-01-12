package ds

// OrderedSet is a non thread-safe ordered set implementation.
type OrderedSet[A comparable] struct {
	set   Set[A]
	items []A
}

// NewOrderedSet creates a new ordered set from the specified values.
func NewOrderedSet[A comparable](values ...A) OrderedSet[A] {
	set := OrderedSet[A]{
		set:   make(Set[A]),
		items: make([]A, 0, len(values)),
	}
	for _, v := range values {
		set.Insert(v)
	}
	return set
}

func (s OrderedSet[A]) Insert(value A) {
	// if it can be inserted into the set, insert it into the list as well
}
