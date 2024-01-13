package ds

import (
	"reflect"
	"testing"
)

func TestOrderedSet(t *testing.T) {
	orderedSet := NewOrderedSet[int]()

	// Test Insert and Contains
	orderedSet.Insert(1)
	if !orderedSet.Contains(1) {
		t.Error("OrderedSet should contain the item that was added")
	}

	// Test Remove
	orderedSet.Remove(1)
	if orderedSet.Contains(1) {
		t.Error("OrderedSet should no longer contain an item that was removed")
	}

	// Test Items and Order
	orderedSet.Insert(2)
	orderedSet.Insert(3)
	orderedSet.Insert(4)
	if !reflect.DeepEqual(orderedSet.Items(), []int{2, 3, 4}) {
		t.Error("Items should return the items in the order they were added")
	}

	// Test Size
	if orderedSet.Size() != 3 {
		t.Errorf("Size should return the number of unique items, got %d", orderedSet.Size())
	}

	// Test Clear
	orderedSet.Clear()
	if orderedSet.Size() != 0 {
		t.Error("Clear should remove all items from the OrderedSet")
	}
}
