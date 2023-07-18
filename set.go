package collections

import "fmt"

// Set is a implementation of a set using a Go hashmap.
type Set[T comparable] map[T]o

// o is a "marker type" representing containment in a Set.
type o struct{}

// String returns a string representation of this Set.
func (s *Set[T]) String() string {
	str := "{"

	for t := range *s {
		str += fmt.Sprintf("%v, ", t)
	}

	// Remove last comma
	str = str[:len(str)-2]
	str += "}"
	return str
}

// Constructors

// NewSet makes a new Set with the specified initial capacity.
func NewSet[T comparable](capacity int) *Set[T] {
	s := make(Set[T], capacity)
	return &s
}

// AsSet returns a Set containing the elements of the given slice.
func AsSet[T comparable](elems []T) *Set[T] {
	s := NewSet[T](len(elems))
	for _, t := range elems {
		s.Add(t)
	}
	return s
}

// Slice returns a slice containing the elements of this Set.
func (s *Set[T]) Slice() []T {
	slice := make([]T, 0, s.Size())
	for t := range *s {
		slice = append(slice, t)
	}
	return slice
}

// Basic (non-mutating) functions

// Size returns the number of elements in this Set.
func (s *Set[T]) Size() int {
	return len(*s)
}

// IsEmpty returns true if this Set is empty.
func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Contains returns true if the given element is in the List.
func (s *Set[T]) Contains(t T) bool {
	_, ok := (*s)[t]
	return ok
}

// Basic (mutating) functions

// Add adds t to this Set, if it is not already in the Set.
func (s *Set[T]) Add(t T) {
	(*s)[t] = o{}
}

// Remove removes t from this Set. It returns false if t was not in the Set
// to begin with, and returns true if t was removed from the Set.
func (s *Set[T]) Remove(t T) bool {
	ret := s.Contains(t)
	delete(*s, t)
	return ret
}

// Copying functions

// Copy returns a copy of the given Set.
func (s *Set[T]) Copy() *Set[T] {
	cp := NewSet[T](s.Size())
	for t := range *s {
		cp.Add(t)
	}
	return cp
}

// Union returns the Set of all elements which are in any of the given Sets.
func Union[T comparable](sets ...*Set[T]) *Set[T] {
	union := NewSet[T](0)
	for _, set := range sets {
		for t := range *set {
			union.Add(t)
		}
	}
	return union
}

// Intersection returns the Set of elements which are in all of the given sets.
func Intersection[T comparable](sets ...*Set[T]) *Set[T] {
	if len(sets) == 0 {
		return nil
	}
	intsec := NewSet[T](sets[0].Size())
elementLoop:
	for t := range *sets[0] {
		for _, set := range sets[1:] {
			if !set.Contains(t) {
				continue elementLoop
			}
		}
		intsec.Add(t)
	}
	return intsec
}

// Difference returns the set of all elements in s1 which are not in s2.
func Difference[T comparable](s1, s2 *Set[T]) *Set[T] {
	diff := NewSet[T](s1.Size())
	for t := range *s1 {
		if !s2.Contains(t) {
			diff.Add(t)
		}
	}
	return diff
}

// SymmetricDifference returns the set of all elements which are in exactly
// one of s1 and s2.
func SymmetricDifference[T comparable](s1, s2 *Set[T]) *Set[T] {
	return Union(Difference(s1, s2), Difference(s2, s1))
}

// TODO: write tests verifying that set theory laws hold
