package set

type Set[T comparable] struct {
	items map[T]struct{}
}

func New[T comparable](items ...T) *Set[T] {
	s := Set[T]{
		items: map[T]struct{}{},
	}
	for _, item := range items {
		s.items[item] = struct{}{}
	}
	return &s
}

func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

func (s *Set[T]) Has(item T) bool {
	_, ok := s.items[item]
	return ok
}
