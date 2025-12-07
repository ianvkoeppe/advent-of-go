package aoc

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) Add(elem T) {
	s.m[elem] = struct{}{}
}

func (s *Set[T]) Remove(elem T) {
	delete(s.m, elem)
}

func (s *Set[T]) Contains(elem T) bool {
	_, ok := s.m[elem]
	return ok
}

func (s *Set[T]) Empty() bool {
	return s.Size() == 0
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, s.Size())
	for elem := range s.m {
		elements = append(elements, elem)
	}
	return elements
}
