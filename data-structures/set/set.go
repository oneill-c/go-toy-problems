package datastructures

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

func (s *Set[T]) Add(v T) {
	s.data[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.data, v)
}

func (s *Set[T]) Has(v T) bool {
	_, exists := s.data[v]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Values() []T {
	if s == nil || len(s.data) == 0 {
		return []T{}
	}
	out := make([]T, 0, len(s.data))
	for k := range s.data {
		out = append(out, k)
	}
	return out
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	out := NewSet[T]()
	if s != nil {
		for k := range s.data {
			out.Add(k)
		}
	}
	if other != nil {
		for k := range other.data {
			out.Add(k)
		}
	}
	return out
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	out := NewSet[T]()
	for k := range s.data {
		if other.Has(k) {
			out.Add(k)
		}
	}
	return out
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	out := NewSet[T]()
	for k := range s.data {
		if !other.Has(k) {
			out.Add(k)
		}
	}
	return out
}