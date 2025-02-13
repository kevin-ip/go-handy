package datastructures

type Stack[X any] struct {
	data []X
}

func NewStack[X any]() *Stack[X] {
	return &Stack[X]{}
}

func (s *Stack[X]) Push(x X) {
	s.data = append(s.data, x)
}

func (s *Stack[X]) Pop() (X, bool) {
	if len(s.data) == 0 {
		var zero X
		return zero, false
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, true
}

func (s *Stack[X]) Peek() (X, bool) {
	if len(s.data) == 0 {
		var zero X
		return zero, false
	}
	x := s.data[len(s.data)-1]
	return x, true
}

func (s *Stack[X]) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack[X]) Size() int {
	return len(s.data)
}
