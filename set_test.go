package ds

import (
	"reflect"
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet(1, 2, 3)

	expected := Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("NewSet() = %v, want %v", s, expected)
	}
}

func TestSet_InsertAndContains(t *testing.T) {
	s := NewSet[int]()
	s.Insert(1)

	if !s.Contains(1) {
		t.Errorf("Set.Contains() = false, want true after Insert")
	}
}

func TestSet_Delete(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Delete(2)

	if s.Contains(2) {
		t.Errorf("Set.Contains() = true, want false after Delete")
	}
}

func TestSet_Union(t *testing.T) {
	s1 := NewSet(1, 2)
	s2 := NewSet(3, 4)

	union := s1.Union(s2)
	expected := NewSet(1, 2, 3, 4)

	if !reflect.DeepEqual(union, expected) {
		t.Errorf("Set.Union() = %v, want %v", union, expected)
	}
}

func TestSet_Intersect(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3, 4)

	intersection := s1.Intersect(s2)
	expected := NewSet(2, 3)

	if !reflect.DeepEqual(intersection, expected) {
		t.Errorf("Set.Intersect() = %v, want %v", intersection, expected)
	}
}

func TestSet_Difference(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(2, 3)

	diff := s1.Difference(s2)
	expected := NewSet(1)

	if !reflect.DeepEqual(diff, expected) {
		t.Errorf("Set.Difference() = %v, want %v", diff, expected)
	}
}

func TestSet_IsSubset(t *testing.T) {
	s1 := NewSet(1, 2)
	s2 := NewSet(1, 2, 3)

	if !s1.IsSubset(s2) {
		t.Errorf("Set.IsSubset() = false, want true")
	}

	if s2.IsSubset(s1) {
		t.Errorf("Set.IsSubset() = true, want false")
	}
}

func TestSet_SizeAndIsEmpty(t *testing.T) {
	s := NewSet[int]()

	if !s.IsEmpty() || s.Size() != 0 {
		t.Errorf("New Set should be empty and size = 0, got IsEmpty() = %v, Size() = %d", s.IsEmpty(), s.Size())
	}

	s.Insert(1)

	if s.IsEmpty() || s.Size() != 1 {
		t.Errorf("Set should have one element and size = 1, got IsEmpty() = %v, Size() = %d", s.IsEmpty(), s.Size())
	}
}
