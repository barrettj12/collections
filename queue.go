package collections

import "fmt"

// Queue is an implementation of a queue using a slice.
type Queue[T any] struct {
	elems []T
}

// Constructors

// NewQueue makes a new Queue with the specified initial capacity.
func NewQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{make([]T, 0, capacity)}
}

// AsQueue returns a Queue backed by the given slice.
func AsQueue[T any](elems []T) *Queue[T] {
	return &Queue[T]{elems}
}

// AsSlice returns the underlying slice for this Queue.
func (q *Queue[T]) AsSlice() []T {
	return q.elems
}

// Basic (non-mutating) functions

// Size returns the number of elements in this Queue.
func (q *Queue[T]) Size() int {
	return len(q.elems)
}

// IsEmpty returns true if this Queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// Capacity returns the current capacity of this Queue.
func (q *Queue[T]) Capacity() int {
	return cap(q.elems)
}

// Peek returns the front element of this Queue, without removing it from
// the Queue. It returns an error if the Queue is empty.
func (q *Queue[T]) Peek() (t T, err error) {
	if q.IsEmpty() {
		err = errQueueEmpty
	} else {
		t = q.elems[0]
	}
	return
}

// Basic (mutating) functions

// Enqueue adds the given element to the back of this Queue.
func (q *Queue[T]) Enqueue(t T) {
	q.elems = append(q.elems, t)
}

// Dequeue removes the front element of the Queue and returns it.
// It returns an error if the Queue is empty.
func (q *Queue[T]) Dequeue() (t T, err error) {
	if q.IsEmpty() {
		err = errQueueEmpty
	} else {
		t = q.elems[0]
		q.elems = q.elems[1:]
	}
	return
}

// Copying functions

// Copy returns a copy of the given Queue.
func (q *Queue[T]) Copy() *Queue[T] {
	elemsCp := make([]T, 0, q.Size())
	copy(elemsCp, q.elems)
	return &Queue[T]{elemsCp}
}

// Errors
var errQueueEmpty = fmt.Errorf("queue is empty")
