/*
Package collections provides Go implementations of basic collections and other data structures.

Instantiate a collection X using the collections.NewX functions, which take an initial capacity as argument:

	l := collections.NewList[string](10)
	m := collections.NewMap[string, int](10)
	s := collections.NewSet[string](10)
	q := collections.NewQueue[int](10)
	k := collections.NewStack[byte](5)

Or convert your existing slices/maps to collections using the collections.AsX functions:

	l := collections.AsList([]int{0, 1})
	m := collections.AsMap(map[int]int{0: 0, 1: 1})
	q := collections.AsQueue([]int{0, 1})
	k := collections.AsStack([]int{0, 1})
*/
package collections
