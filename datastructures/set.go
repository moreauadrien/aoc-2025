package datastructures

import "iter"

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, found := s.elements[value]
	return found
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) List() []T {
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for key := range s.elements {
			if !yield(key) {
				return
			}
		}
	}
}
