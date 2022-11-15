package set

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"golang.org/x/exp/constraints"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestNew(t *testing.T) {
	threadsafe := New[int](true)
	_, ok := threadsafe.(*ThreadsafeSet[int])
	if !ok {
		t.Fatalf("expected new with true to return ThreadsafeSet but got %T\n", threadsafe)
	}
	standard := New[int]()
	_, ok = standard.(*StandardSet[int])
	if !ok {
		t.Fatalf("expected new with no bool to return StandardSet but got %T\n", standard)
	}
}

func SlicesMatch[T constraints.Ordered](first, second []T) bool {
	if len(first) != len(second) {
		return false
	}
	for i, t := range first {
		if second[i] != t {
			return false
		}
	}
	return true
}

func TestStandard(t *testing.T) {
	var demo = []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := NewStandardSet[int]()
	if slice := s.Slice(); !SlicesMatch(slice, []int{}) {
		t.Fatalf("error: expected empty StandardSet[int] after construction but it contains %v\n", slice)
	}
	s.Add(demo...)
	if slice := s.Slice(); !SlicesMatch(slice, demo) {
		t.Fatalf("error: expected slices to match after Add but they do not\n")
	}
	s.Drop(demo...)
	if slice := s.Slice(); !SlicesMatch(slice, []int{}) {
		t.Fatalf("error: expected empty slice after full Drop but found %v", slice)
	}
	s.Add(demo...)
	s.Add(demo...)
	if slice := s.Clear(); !SlicesMatch(slice, demo) {
		t.Fatalf("error: Clear() slice %v does not match input slice %v\n", slice, demo)
	}
	if s.Len() != 0 {
		t.Fatalf("error: Clear() failed to reset the count: expected 0, got %d\n", s.Len())
	}
}

func TestThreadsafe(t *testing.T) {
	var demo = []int{1, 2, 3, 4, 5, 6, 7, 8}
	s := NewThreadsafeSet[int]()
	var wg sync.WaitGroup
	for _, n := range demo {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(n)
	}
	for _, n := range demo {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Add(i)
		}(n)
	}
	wg.Wait()
	if slice := s.Slice(); !SlicesMatch(slice, demo) {
		t.Fatalf("error: expected slices to match after concurrent Adds but they do not\n")
	}
	for _, n := range demo {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Drop(i)
		}(n)
	}
	wg.Wait()
	if slice := s.Slice(); !SlicesMatch(slice, []int{}) {
		t.Fatalf("error: expected empty slice after concurrent Drops but found %v", slice)
	}
}

const RandomNumberSampleSize = 1000000

var RandomNumbers = func() []int {
	results := make([]int, 0, RandomNumberSampleSize)
	for i := 0; i < RandomNumberSampleSize; i++ {
		results = append(results, rand.Int())
	}
	return results
}()

func GetRandomNumber() int {
	return RandomNumbers[rand.Intn(RandomNumberSampleSize)]
}

func BenchmarkStandardNewFromSlice(b *testing.B) {
	sample := func() []int {
		results := make([]int, 0, 100)
		for i := 0; i < 100; i++ {
			results = append(results, GetRandomNumber())
		}
		return results
	}()
	for i := 0; i < b.N; i++ {
		NewStandardSetFromSlice(sample)
	}
}

func BenchmarkThreadsafeNewFromSlice(b *testing.B) {
	sample := func() []int {
		results := make([]int, 0, 100)
		for i := 0; i < 100; i++ {
			results = append(results, GetRandomNumber())
		}
		return results
	}()
	for i := 0; i < b.N; i++ {
		NewThreadsafeSetFromSlice(sample)
	}
}

func BenchmarkStandardSetAdd(b *testing.B) {
	set := NewStandardSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(GetRandomNumber())
	}
}

func BenchmarkThreadsafeSetAdd(b *testing.B) {
	set := NewThreadsafeSet[int]()
	for i := 0; i < b.N; i++ {
		set.Add(GetRandomNumber())
	}
}

func BenchmarkStandardContains(b *testing.B) {
	sample := func() []int {
		results := make([]int, 0, 100)
		for i := 0; i < 100; i++ {
			results = append(results, GetRandomNumber())
		}
		return results
	}()
	set := NewStandardSetFromSlice(sample)
	for i := 0; i < b.N; i++ {
		set.Contains(GetRandomNumber())
	}
}

func BenchmarkThreadsafeSetContains(b *testing.B) {
	sample := func() []int {
		results := make([]int, 0, 100)
		for i := 0; i < 100; i++ {
			results = append(results, GetRandomNumber())
		}
		return results
	}()
	set := NewThreadsafeSetFromSlice(sample)
	for i := 0; i < b.N; i++ {
		set.Contains(GetRandomNumber())
	}
}
