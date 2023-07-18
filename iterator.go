package collections

// Iterator allows iteration over a collection.
type Iterator[T any] interface {
	// HasNext returns true if there are more values in the iterator.
	HasNext() bool

	// Next returns the next value in the iterator, and advances iteration.
	// Implementations of Next may panic if there are no more values to return.
	// To be safe, always call HasNext before calling Next.
	Next() T
}

// Iterator2 allows iteration over a bivalent collection (e.g. a Map).
type Iterator2[T, U any] interface {
	// HasNext returns true if there are more values in the iterator.
	HasNext() bool

	// Next returns the next value in the iterator, and advances iteration.
	// Implementations of Next may panic if there are no more values to return.
	// To be safe, always call HasNext before calling Next.
	Next() (T, U)
}
