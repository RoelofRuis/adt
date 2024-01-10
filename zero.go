package ds

// zeroValue returns the zero value for a variable type A.
func zeroValue[A any]() A {
	var zero A
	return zero
}
