# package set
package set implements a Set interface for types defined by constraints.Ordered (https://pkg.go.dev/golang.org/x/exp/constraints).
it also includes both a standard and threadsafe Set types.

## Important Types and Functions
```go
type Set interface {
    Add(...T) // adds a variadic slice of T
	Drop(...T) // drops a variadic slice of T
	Contains(T) bool // checks for T, returns true if found
	Len() int // returns len(contents)
	Slice() []T // returns a sorted slice of contents
	Clear() []T // stores a slice, resets the contents then returns the previous slice
	Reserve(int) // creates a buffer with the specified reserved size, fills it with the contents then swaps them
}

/* takes a variadic bool to indicate whether or not the returned Set should be threadsafe */
func New[T constraints.Ordered](doThreadsafe ...bool) Set[T] {}
/* takes a slice of T and a variadic bool to indicate whether or not the returned Set should be threadsafe */
func NewFromSlice[T constraints.Ordered](slice []T, doThreadsafe ...bool) Set[T] {}
```