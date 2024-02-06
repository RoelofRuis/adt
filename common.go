package ds

// Comparator specifies a comparison function that determines the ordering of values.
// It returns an int indicating the ordering: negative if a < b, zero if a == b, and positive if a > b.
type Comparator[A any] func(a, b A) int

// CompareInt implements a Comparator for int.
func CompareInt(a, b int) int {
	return a - b
}

// zeroValue returns the zero value for a variable type A.
func zeroValue[A any]() A {
	var zero A
	return zero
}
