package ds

// Interval represents a range over the type A.
type Interval[A any] struct {
	Start A
	End   A
}

// NewIntervalBetween creates a new interval between the given values.
func NewIntervalBetween[A any](start A, end A) Interval[A] {
	return Interval[A]{Start: start, End: end}
}

// NewInstantInterval creates a new interval with zero length at the given position.
func NewInstantInterval[A any](at A) Interval[A] {
	return Interval[A]{Start: at, End: at}
}
