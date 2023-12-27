package adt

import (
	"fmt"
	"strings"
)

// Set is a non thread-safe set implementation.
type Set[A comparable] map[A]struct{}

// emptyValue is a placeholder for the map values in the set.
var emptyValue struct{}

// NewSet creates a new set from the specified values.
func NewSet[A comparable](values ...A) *Set[A] {
	set := make(Set[A])
	for _, v := range values {
		set.Insert(v)
	}
	return &set
}

// Clone creates a deep copy of the set.
func (s *Set[A]) Clone() *Set[A] {
	clone := NewSet[A]()
	for elem := range *s {
		clone.Insert(elem)
	}
	return clone
}

// Intersect modifies the set by keeping only the elements that exist in both this and 'that' set.
func (s *Set[A]) Intersect(that Set[A]) {
	for elem := range *s {
		if !that.Contains(elem) {
			delete(*s, elem)
		}
	}
}

// Union modifies the set by adding all elements from 'that' into this set.
func (s *Set[A]) Union(that Set[A]) {
	for elem := range that {
		s.Insert(elem)
	}
}

// Difference modifies the set by removing all elements that are present in 'that'.
func (s *Set[A]) Difference(that Set[A]) {
	for elem := range that {
		delete(*s, elem)
	}
}

// Insert adds the specified value to the set.
func (s *Set[A]) Insert(value A) {
	(*s)[value] = emptyValue
}

// Contains checks if the set contains the specified value.
func (s *Set[A]) Contains(value A) bool {
	_, has := (*s)[value]
	return has
}

// ContainsOneOf checks if the set contains any of the specified values.
func (s *Set[A]) ContainsOneOf(values []A) bool {
	for _, candidate := range values {
		if s.Contains(candidate) {
			return true
		}
	}
	return false
}

// IsSubset checks if the set is a subset of 'that'.
func (s *Set[A]) IsSubset(that Set[A]) bool {
	for elem := range *s {
		if !that.Contains(elem) {
			return false
		}
	}
	return true
}

// Size returns the number of elements in the set.
func (s *Set[A]) Size() int {
	return len(*s)
}

// IsEmpty checks if the set is empty.
func (s *Set[A]) IsEmpty() bool {
	return s.Size() == 0
}

// String returns a string representation of the set.
func (s *Set[A]) String() string {
	var res []string
	for e := range *s {
		res = append(res, fmt.Sprintf("%v", e))
	}
	return fmt.Sprintf("Set[%s]", strings.Join(res, " "))
}
