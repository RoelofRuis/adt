package adt

import (
	"reflect"
	"testing"
)

func TestNewSetAndString(t *testing.T) {
	set := NewSet(1, 2, 3)
	expectedString := "Set[1 2 3]"
	if result := set.String(); result != expectedString {
		t.Errorf("String representation is incorrect. Expected: %s, Got: %s", expectedString, result)
	}
}

func TestInsertAndSize(t *testing.T) {
	set := NewSet(1, 2, 3)
	set.Insert(4)
	if size := set.Size(); size != 4 {
		t.Errorf("Size is incorrect after insertion. Expected: 4, Got: %d", size)
	}
}

func TestSet_Contains(t *testing.T) {
	set := NewSet(1, 2, 3)
	if !set.Contains(3) {
		t.Error("Set should contain the element 3.")
	}
	if set.Contains(5) {
		t.Error("Set should not contain the element 5.")
	}
}

func TestSet_ContainsOneOf(t *testing.T) {
	set := NewSet(1, 2, 3)
	if !set.ContainsOneOf([]int{2, 5, 6}) {
		t.Error("Set should contain at least one of the elements [2, 5, 6].")
	}
	if set.ContainsOneOf([]int{5, 6, 7}) {
		t.Error("Set should not contain any of the elements [5, 6, 7].")
	}
}

func TestSet_Intersect(t *testing.T) {
	set := NewSet(1, 2, 3)
	otherSet := NewSet(2, 3, 4)
	set.Intersect(*otherSet)
	expectedIntersect := NewSet(2, 3)
	if !reflect.DeepEqual(set, expectedIntersect) {
		t.Errorf("Intersect operation is incorrect. Expected: %v, Got: %v", expectedIntersect, set)
	}
}

func TestSet_Union(t *testing.T) {
	set := NewSet(1, 2, 3)
	otherSet := NewSet(3, 4, 5)
	set.Union(*otherSet)
	expectedUnion := NewSet(1, 2, 3, 4, 5)
	if !reflect.DeepEqual(set, expectedUnion) {
		t.Errorf("Union operation is incorrect. Expected: %v, Got: %v", expectedUnion, set)
	}
}

func TestSet_Difference(t *testing.T) {
	set := NewSet(1, 2, 3)
	otherSet := NewSet(3, 5)
	set.Difference(*otherSet)
	expectedDifference := NewSet(1, 2)
	if !reflect.DeepEqual(set, expectedDifference) {
		t.Errorf("Difference operation is incorrect. Expected: %v, Got: %v", expectedDifference, set)
	}
}

func TestSet_IsSubset(t *testing.T) {
	set := NewSet(1, 2, 3)
	otherSet := NewSet(2, 3)
	if !otherSet.IsSubset(*set) {
		t.Error("Set should be a subset of the other set.")
	}
	otherSet.Insert(6)
	if set.IsSubset(*otherSet) {
		t.Error("Set should not be a subset after insertion.")
	}
}

func TestSet_Clone(t *testing.T) {
	set := NewSet(1, 2, 3)
	clone := set.Clone()
	if !reflect.DeepEqual(*set, *clone) {
		t.Error("Clone operation did not produce an equal set.")
	}
}
