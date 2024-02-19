# package set
package set implements a Set interface for types defined by constraints.Ordered (https://pkg.go.dev/golang.org/x/exp/constraints)

## Important Types and Functions
```go   
type Set[T constraints.Ordered] interface {
	Add(T) bool
	AddSlice([]T) []bool
	Contains(T) bool
	ContainsSlice([]T) []bool
	Drop(T) bool
	DropSlice([]T) []bool
	Len() int
	Reset()
	Slice() []T
}