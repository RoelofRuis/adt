package ds

import (
	"fmt"
	"strings"
)

// Set is a non thread-safe set implementation.
type Set[A comparable] map[A]struct{}

// NewSet creates a new set from the specified values.
func NewSet[A comparable](values ...A) Set[A] {
	set := make(Set[A])
	for _, v := range values {
		set.Insert(v)
	}
	return set
}

// Clone creates a deep copy of the set.
func (s Set[A]) Clone() Set[A] {
	clone := make(Set[A], len(s))
	for elem := range s {
		clone.Insert(elem)
	}
	return clone
}

// Values returns the values in the set as a slice.
func (s Set[A]) Values() []A {
	values := make([]A, 0, s.Size())
	for elem := range s {
		values = append(values, elem)
	}
	return values
}

// Intersect creates the set that contains only the elements that exist in both this and 'that' set.
func (s Set[A]) Intersect(that Set[A]) Set[A] {
	res := NewSet[A]()

	var smaller, larger Set[A]
	if s.Size() < that.Size() {
		smaller, larger = s, that
	} else {
		smaller, larger = that, s
	}

	for elem := range smaller {
		if larger.Contains(elem) {
			res.Insert(elem)
		}
	}

	return res
}

// Union creates the set that contains all elements from this and 'that' set.
func (s Set[A]) Union(that Set[A]) Set[A] {
	res := make(Set[A], len(s)+len(that))
	for elem := range s {
		res.Insert(elem)
	}
	for elem := range that {
		res.Insert(elem)
	}
	return res
}

// Difference creates the set that contains all elements in this but not in 'that'.
func (s Set[A]) Difference(that Set[A]) Set[A] {
	res := s.Clone()
	for elem := range that {
		res.Delete(elem)
	}
	return res
}

// Insert adds the specified value to the set.
func (s Set[A]) Insert(value A) {
	(s)[value] = struct{}{}
}

// Delete deletes the specified value from the set.
func (s Set[A]) Delete(value A) {
	delete(s, value)
}

// Contains checks if the set contains the specified value.
func (s Set[A]) Contains(value A) bool {
	_, has := (s)[value]
	return has
}

// ContainsOneOf checks if the set contains any of the specified values.
func (s Set[A]) ContainsOneOf(values []A) bool {
	for _, candidate := range values {
		if s.Contains(candidate) {
			return true
		}
	}
	return false
}

// IsSubset checks if the set is a subset of 'that'.
func (s Set[A]) IsSubset(that Set[A]) bool {
	for elem := range s {
		if !that.Contains(elem) {
			return false
		}
	}
	return true
}

// Size returns the number of elements in the set.
func (s Set[A]) Size() int {
	return len(s)
}

// IsEmpty checks if the set is empty.
func (s Set[A]) IsEmpty() bool {
	return s.Size() == 0
}

// String returns a string representation of the set.
func (s Set[A]) String() string {
	var b strings.Builder
	b.WriteString("Set[")
	first := true
	for e := range s {
		if !first {
			b.WriteString(" ")
		}
		b.WriteString(fmt.Sprintf("%v", e))
		first = false
	}
	b.WriteString("]")
	return b.String()
}
