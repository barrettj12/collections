package collections

import "fmt"

// Map is an implementation of a map using Go's hashmap.
type Map[K comparable, V any] map[K]V

// Constructors

// NewMap makes a new Map with the specified initial capacity.
func NewMap[K comparable, V any](capacity int) *Map[K, V] {
	m := make(Map[K, V], capacity)
	return &m
}

// AsMap returns a Map backed by the given map.
func AsMap[K comparable, V any](m map[K]V) *Map[K, V] {
	M := Map[K, V](m)
	return &M
}

// Underlying returns the underlying map for this Map.
func (m *Map[K, V]) AsSlice() map[K]V {
	return *m
}

// Basic (non-mutating) functions

// Size returns the number of entries in this Map.
func (m *Map[K, V]) Size() int {
	return len(*m)
}

// IsEmpty returns true if this Map is empty.
func (m *Map[K, V]) IsEmpty() bool {
	return m.Size() == 0
}

// Contains returns true if the given key is in the Map.
func (m *Map[K, V]) Contains(k K) bool {
	_, ok := (*m)[k]
	return ok
}

// Get returns the value associated with k. If k is not a key in this Map,
// Get returns an error.
func (m *Map[K, V]) Get(k K) (v V, err error) {
	if !m.Contains(k) {
		err = errKeyNotFound(k)
	}
	v = (*m)[k]
	return
}

// Basic (mutating) functions

// Set adds the given key-value pair to this Map. If there is already a value
// associated with k, it will be overwritten.
func (m *Map[K, V]) Set(k K, v V) {
	(*m)[k] = v
}

// Remove removes k and its associated value from this Map. It returns false
// if k was not in the Map to begin with, and returns true if k was removed.
func (m *Map[K, V]) Remove(k K) bool {
	ret := m.Contains(k)
	delete(*m, k)
	return ret
}

// Copying functions

// Copy returns a copy of the given Map.
func (m *Map[K, V]) Copy() *Map[K, V] {
	cp := NewMap[K, V](m.Size())
	for k, v := range *m {
		cp.Set(k, v)
	}
	return cp
}

// Errors

func errKeyNotFound(k any) error {
	return fmt.Errorf("key not found in Map: %v", k)
}
