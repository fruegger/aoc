package ds

import "slices"

type Set[T comparable] struct {
	elements []T
}

func (s *Set[T]) Add(el T) {
	if !slices.Contains(s.elements, el) {
		s.elements = append(s.elements, el)
	}
}

func (s *Set[T]) Remove(el T) {
	p := slices.Index(s.elements, el)
	if p >= 0 {
		s.elements = append(s.elements[:p], s.elements[p+1:]...)
	}
}

func (s *Set[T]) Intersection(b Set[T]) Set[T] {
	var r []T
	for _, v := range s.elements {
		if slices.Contains(b.elements, v) {
			r = append(r, v)
		}
	}
	return Set[T]{elements: r}
}

func (s *Set[T]) Union(b Set[T]) Set[T] {
	var r = Set[T]{elements: s.elements}
	for _, v := range b.elements {
		r.Add(v)
	}
	return r
}

func (s *Set[T]) Elements() *[]T {
	return &s.elements
}

func (s *Set[T]) Size() int { return len(s.elements) }
func (s *Set[T]) Contains(el T) bool {
	return slices.Contains(s.elements, el)
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
