package set

import (
	"golang.org/x/exp/constraints"
)

type Set[T constraints.Ordered] interface {
	Add(...T)
	Drop(...T)
	Contains(T) bool
	Len() int
	Slice() []T
	Clear() []T
	Reserve(int)
}

func New[T constraints.Ordered](doThreadsafe ...bool) Set[T] {
	if len(doThreadsafe) == 1 && doThreadsafe[0] {
		return NewThreadsafeSet[T]()
	}
	return NewStandardSet[T]()
}

func NewFromSlice[T constraints.Ordered](slice []T, doThreadsafe ...bool) Set[T] {
	if len(doThreadsafe) == 1 && doThreadsafe[0] {
		return NewThreadsafeSetFromSlice(slice)
	}
	return NewStandardSetFromSlice(slice)
}
