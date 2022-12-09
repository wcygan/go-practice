package stack

// Stack implements a LIFO stack with peeking.
type Stack[T any] struct {
	data []T
}

// New returns an empty stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{
		data: nil,
	}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (t T) {
	if len(s.data) == 0 {
		return t
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}
