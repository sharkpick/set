package set

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type ThreadsafeSet[T constraints.Ordered] struct {
	contents *StandardSet[T]
	mutex    sync.RWMutex
}

func NewThreadsafeSet[T constraints.Ordered]() *ThreadsafeSet[T] {
	return &ThreadsafeSet[T]{
		contents: NewStandardSet[T](),
	}
}

func NewThreadsafeSetFromSlice[T constraints.Ordered](slice []T) *ThreadsafeSet[T] {
	return &ThreadsafeSet[T]{
		contents: NewStandardSetFromSlice(slice),
	}
}

func (s *ThreadsafeSet[T]) Reserve(size int) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.contents.Reserve(size)
}

func (s *ThreadsafeSet[T]) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.contents.Len()
}

func (s *ThreadsafeSet[T]) Add(slice ...T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.contents.Add(slice...)
}

func (s *ThreadsafeSet[T]) Drop(slice ...T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.contents.Drop(slice...)
}

func (s *ThreadsafeSet[T]) Contains(t T) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.contents.Contains(t)
}

func (s *ThreadsafeSet[T]) Slice() []T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.contents.Slice()
}

func (s *ThreadsafeSet[T]) Clear() []T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.contents.Clear()
}
