package datastruct

type Queue[T any] struct {
	items []T
}

func (s *Queue[T]) Enqueue(item T) {
	s.items = append(s.items, item)
}

func (s *Queue[T]) Dequeue() T {
	v := s.items[0]
	s.items = s.items[1:]
	return v
}

func (s *Queue[T]) Empty() bool {
	return len(s.items) == 0
}
