# collections

The `collections` package provides Go implementations of basic collections and
other data structures. This library has several motivations:

- Provide a nicer syntax for operations on Go's built-in data structures
  (slice, map, etc), by using *methods* for these operations, instead of
  *functions*. See the [comparison](#comparison) section below.
  

- Define standard implementations of data structures which are not built-in to
  Go, such as set, stack, queue.

- Provide methods which encapsulate common *compound operations* , such as
  inserting into a list, copying a map, checking containment in a list.

This library is heavily inspired by the
[Java Collections Framework](https://docs.oracle.com/en/java/javase/19/docs/api/java.base/java/util/doc-files/coll-overview.html).
It is open-source under the
[MIT license](https://github.com/barrettj12/collections/blob/main/LICENSE.md).

Read the package documentation
[here](https://pkg.go.dev/github.com/barrettj12/collections).


## Using the `collections` library

The collections defined by this library are generic, so Go 1.18+ is required to
use the library. To download the library:

```
go get github.com/barrettj12/collections
```

See the
[package documentation](https://pkg.go.dev/github.com/barrettj12/collections)
for usage and examples.

**This library is currently unstable, and the API is subject to
backwards-incompatible changes at any time. Fix a SHA version in your
`go.mod`.**


## Comparison

### List
```go
l := make([]string, 0, 10)
l = append(l, "foo", "bar")
len(l) // 2
l[0]   // foo

l2 := make([]string, 0, len(l))
copy(l2, l)
```
```go
l := collections.NewList[string](10)
l.Append("foo", "bar")
l.Size() // 2
l.Get(0) // foo
l2 := l.Copy()
```

### Map
```go
m := make(map[string]int, 0, 10)
m["one"] = 1
m["two"] = 2
delete(m, "one")
len(m) // 1
v, ok := m["two"]
```
```go
m := collections.NewMap[string, int](10)
m.Set("one", 1)
m.Set("two", 2)
rmd := m.Remove("one")
m.Size() // 1
v, err := m.Get("two")
```

## Contributing

Contributions are welcome. Please
[open an issue](https://github.com/barrettj12/collections/issues/new) or
[submit a pull request](https://github.com/barrettj12/collections/pulls).