package set

import "golang.org/x/exp/constraints"

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
