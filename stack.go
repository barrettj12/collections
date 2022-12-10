package collections

import "fmt"

// Stack is an implementation of a stack using a slice.
type Stack[T any] struct {
	elems []T
}

// Constructors

// NewStack makes a new Stack with the specified initial capacity.
func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{make([]T, 0, capacity)}
}

// AsStack returns a Stack backed by the given slice.
func AsStack[T any](elems []T) *Stack[T] {
	return &Stack[T]{elems}
}

// AsSlice returns the underlying slice for this Stack.
func (s *Stack[T]) AsSlice() []T {
	return s.elems
}

// Basic (non-mutating) functions

// Size returns the number of elements in this Stack.
func (s *Stack[T]) Size() int {
	return len(s.elems)
}

// IsEmpty returns true if this Stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Capacity returns the current capacity of this Stack.
func (s *Stack[T]) Capacity() int {
	return cap(s.elems)
}

// Peek returns the top element of this Stack, without removing it from
// the Stack. It returns an error if the Stack is empty.
func (s *Stack[T]) Peek() (t T, err error) {
	if s.IsEmpty() {
		err = errStackEmpty
	} else {
		t = s.elems[len(s.elems)-1]
	}
	return
}

// Basic (mutating) functions

// Push adds the given element to the top of this Stack.
func (s *Stack[T]) Push(t T) {
	s.elems = append(s.elems, t)
}

// Pop removes the top element of the Stack and returns it.
// It returns an error if the Stack is empty.
func (s *Stack[T]) Pop() (t T, err error) {
	if s.IsEmpty() {
		err = errStackEmpty
	} else {
		t = s.elems[len(s.elems)-1]
		s.elems = s.elems[:len(s.elems)-1]
	}
	return
}

// Copying functions

// Copy returns a copy of the given Stack.
func (s *Stack[T]) Copy() *Stack[T] {
	elemsCp := make([]T, 0, s.Size())
	copy(elemsCp, s.elems)
	return &Stack[T]{elemsCp}
}

// Errors
var errStackEmpty = fmt.Errorf("stack is empty")
