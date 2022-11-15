package set

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type StandardSet[T constraints.Ordered] struct {
	contents map[T]any
}

func NewStandardSet[T constraints.Ordered]() *StandardSet[T] {
	return &StandardSet[T]{
		contents: make(map[T]any),
	}
}

func NewStandardSetFromSlice[T constraints.Ordered](slice []T) *StandardSet[T] {
	set := &StandardSet[T]{
		contents: make(map[T]any, len(slice)),
	}
	for _, t := range slice {
		set.Add(t)
	}
	return set
}

func (s *StandardSet[T]) Reserve(size int) {
	buffer := make(map[T]any, size)
	for t := range s.contents {
		buffer[t] = struct{}{}
	}
	s.contents = buffer
}

func (s *StandardSet[T]) Len() int {
	return len(s.contents)
}

func (s *StandardSet[T]) Add(slice ...T) {
	for _, t := range slice {
		s.contents[t] = struct{}{}
	}
}

func (s *StandardSet[T]) Drop(slice ...T) {
	for _, t := range slice {
		delete(s.contents, t)
	}
}

func (s *StandardSet[T]) Contains(t T) bool {
	_, found := s.contents[t]
	return found
}

func (s *StandardSet[T]) Slice() []T {
	results := make([]T, 0, len(s.contents))
	for t := range s.contents {
		results = append(results, t)
	}
	sort.Slice(results, func(i, j int) bool { return results[i] < results[j] })
	return results
}

func (s *StandardSet[T]) Clear() []T {
	slice := s.Slice()
	s.contents = make(map[T]any)
	return slice
}
