package collections

import "fmt"

// List is an implementation of a list using a slice.
type List[T comparable] []T

// Constructors

// NewList makes a new List with the specified initial capacity.
func NewList[T comparable](capacity int) *List[T] {
	l := make(List[T], 0, capacity)
	return &l
}

// AsList returns a List backed by the given slice.
func AsList[T comparable](s []T) *List[T] {
	l := List[T](s)
	return &l
}

// AsSlice returns the underlying slice for this List.
func (l *List[T]) AsSlice() []T {
	return *l
}

// Basic (non-mutating) functions

// Len returns the number of elements in this List.
func (l *List[T]) Len() int {
	return len(*l)
}

// Size returns the number of elements in this List.
func (l *List[T]) Size() int {
	return len(*l)
}

// IsEmpty returns true if this List is empty.
func (l *List[T]) IsEmpty() bool {
	return l.Size() == 0
}

// Capacity returns the current capacity of this List.
func (l *List[T]) Capacity() int {
	return cap(*l)
}

// Contains returns true if the given element is in the List.
func (l *List[T]) Contains(t T) bool {
	for _, s := range *l {
		if s == t {
			return true
		}
	}
	return false
}

// Basic (mutating) functions

// Append appends the given elements to the end of the List.
func (l *List[T]) Append(t ...T) {
	*l = append(*l, t...)
}

// Insert inserts t at position pos in this List.
// It returns an error if the given index is out of bounds.
func (l *List[T]) Insert(t T, pos int) error {
	if pos < 0 || pos >= l.Size() {
		return l.errIndexOutOfBounds(pos)
	}

	var z T
	*l = append(*l, z)
	copy((*l)[pos+1:], (*l)[pos:])
	(*l)[pos] = t
	return nil
}

// Remove removes the element at the given index in this List,
// shifting all other elements down to "fill in the gap".
// It returns an error if the given index is out of bounds.
func (l *List[T]) Remove(pos int) error {
	if pos < 0 || pos >= l.Size() {
		return l.errIndexOutOfBounds(pos)
	}

	*l = append((*l)[:pos], (*l)[pos+1:]...)
	return nil
}

// RemoveAll removes all occurrences of the given element in the List.
func (l *List[T]) RemoveAll(t T) {
	n := 0
	for _, s := range *l {
		if s == t {
			(*l)[n] = s
			n++
		}
	}
	*l = (*l)[:n]
}

// Slice slices this List at the given indices, removing all elements with
// index smaller than low or at least high.
// It returns an error if either index is out of bounds, or if low is
// greater than high.
//
// Slice accepts "Python-style" indices which can be negative:
// -1 refers to the last element of the List, -2 the second last, etc.
func (l *List[T]) Slice(low, high int) error {
	low, high, err := l.checkIndices(low, high)
	if err != nil {
		return err
	}

	*l = (*l)[low:high]
	return nil
}

// Copying functions

// Copy returns a copy of the given List.
func (l *List[T]) Copy() *List[T] {
	lcopy, _ := l.CopyPart(0, l.Size())
	return lcopy
}

// CopyPart returns a partial copy of the given List, i.e. it copies the part
// of the List starting at low and ending at high-1.
// It returns an error if either index is out of bounds, or if low is
// greater than high.
//
// CopyPart accepts "Python-style" indices which can be negative:
// -1 refers to the last element of the List, -2 the second last, etc.
func (l *List[T]) CopyPart(low, high int) (*List[T], error) {
	low, high, err := l.checkIndices(low, high)
	if err != nil {
		return nil, err
	}

	slice := make([]T, high-low)
	copy(slice, (*l)[low:high])
	lcopy := List[T](slice)
	return &lcopy, nil
}

// Internal methods

func (l *List[T]) checkIndices(low, high int) (lo int, hi int, err error) {
	lo, err = l.normaliseIndex(low)
	if err != nil {
		return
	}
	hi, err = l.normaliseIndex(low)
	if err != nil {
		return
	}
	if lo > hi {
		err = l.errLowAboveHigh(lo, hi)
	}
	return
}

func (l *List[T]) normaliseIndex(i int) (int, error) {
	if i <= -l.Size() || i > l.Size() {
		return 0, l.errIndexOutOfBounds(i)
	}
	if i < 0 {
		i += l.Size()
	}
	return i, nil
}

// Errors

func (l *List[T]) errIndexOutOfBounds(pos int) error {
	return fmt.Errorf("index %d out of bounds in List (size %d)", pos, l.Size())
}

func (l *List[T]) errLowAboveHigh(low, high int) error {
	return fmt.Errorf("low index %d greater than high index %d", low, high)
}
