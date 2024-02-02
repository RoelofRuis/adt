package ds

import "testing"

func addInts(a int, b int) int {
	return a + b
}

func subtractInts(a int, b int) int {
	return a - b
}

func TestAugmentedMap(t *testing.T) {
	m := NewAugmentedMap[string, int](addInts, subtractInts)

	// Test Insert
	m.Insert("a", 10)
	sum := m.Sum()
	if sum != 10 {
		t.Errorf("TotalSum after Inserting 'a': expected 10, got %d", sum)
	}

	// Test Update existing key
	m.Insert("a", 20)
	sum = m.Sum()
	if sum != 20 {
		t.Errorf("TotalSum after updating 'a': expected 20, got %d", sum)
	}

	// Test Insert another key
	m.Insert("b", 30)
	sum = m.Sum()
	if sum != 50 {
		t.Errorf("TotalSum after Inserting 'b': expected 50, got %d", sum)
	}

	// Test Get
	val, exists := m.Get("a")
	if !exists || val != 20 {
		t.Errorf("Get 'a': expected (20, true), got (%d, %t)", val, exists)
	}

	// Test Delete existing key
	m.Delete("a")
	sum = m.Sum()
	if sum != 30 {
		t.Errorf("TotalSum after Deleting 'a': expected 30, got %d", sum)
	}

	// Test Delete non-existing key
	m.Delete("c") // should not alter sum
	sum = m.Sum()
	if sum != 30 {
		t.Errorf("TotalSum after Deleting non-existing 'c': expected 30, got %d", sum)
	}
}
