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

// Keys returns all keys present in this Map.
func (m *Map[K, V]) Keys() *List[K] {
	keys := NewList[K](m.Size())
	for k := range *m {
		keys.Append(k)
	}
	return keys
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

// Iteration

type mapIterator[K comparable, V any] struct {
	m     *Map[K, V]
	keys  *List[K]
	index int
}

func (i *mapIterator[K, V]) HasNext() bool {
	return i.index < i.keys.Size()
}

func (i *mapIterator[K, V]) Next() (K, V) {
	k, _ := i.keys.Get(i.index)
	i.index++
	v, _ := i.m.Get(k)
	return k, v
}

// Iterate returns an Iterator2 iterating over the given map.
// If a non-nil comparator keyOrder is provided, then the iteration will be in
// the order determined on the keys.
func (m *Map[K, V]) Iterate(keyOrder func(K, K) bool) Iterator2[K, V] {
	keys := m.Keys()
	if keyOrder != nil {
		keys.Sort(keyOrder)
	}
	return &mapIterator[K, V]{
		m:     m,
		keys:  keys,
		index: 0,
	}
}

// Errors

func errKeyNotFound(k any) error {
	return fmt.Errorf("key not found in Map: %v", k)
}
