package ds

// AugmentedMap is a non thread-safe map that keeps a sum of the values it contains.
type AugmentedMap[K comparable, V any] struct {
	inner    map[K]V
	add      func(V, V) V
	subtract func(V, V) V
	sum      V
}

// NewAugmentedMap creates a new augmented map from the given arithmetic operations.
func NewAugmentedMap[K comparable, V any](add func(V, V) V, subtract func(V, V) V) *AugmentedMap[K, V] {
	var zero V
	return &AugmentedMap[K, V]{
		inner:    make(map[K]V),
		add:      add,
		subtract: subtract,
		sum:      zero,
	}
}

// Insert inserts a new element into the map.
func (m *AugmentedMap[K, V]) Insert(k K, v V) {
	oldValue, exists := m.inner[k]
	if exists {
		m.sum = m.subtract(m.sum, oldValue)
	}
	m.inner[k] = v
	m.sum = m.add(m.sum, v)
}

// Delete deletes an element from the map.
func (m *AugmentedMap[K, V]) Delete(k K) {
	if oldValue, exists := m.inner[k]; exists {
		delete(m.inner, k)
		m.sum = m.subtract(m.sum, oldValue)
	}
}

// Get retrieves an element from the map without removing it.
func (m *AugmentedMap[K, V]) Get(k K) (V, bool) {
	value, exists := m.inner[k]
	return value, exists
}

// Sum returns the currently collected sum.
func (m *AugmentedMap[K, V]) Sum() V {
	return m.sum
}
