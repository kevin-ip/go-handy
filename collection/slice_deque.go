package collection

// sliceDeque stands for double ended queue
type sliceDeque[X comparable] struct {
	data []X
}

// NewSliceDeque creates a new deque backed by a slice
func NewSliceDeque[X comparable]() Deque[X] {
	return &sliceDeque[X]{}
}

// Stack specifics

// Push adds an element to the top of the deque
func (s *sliceDeque[X]) Push(x X) {
	s.data = append(s.data, x)
}

// Pop removes the top element from the deque
func (s *sliceDeque[X]) Pop() (X, bool) {
	if len(s.data) == 0 {
		var zero X
		return zero, false
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, true
}

// Peek views the top element from the deque
func (s *sliceDeque[X]) Peek() (X, bool) {
	if len(s.data) == 0 {
		var zero X
		return zero, false
	}
	x := s.data[len(s.data)-1]
	return x, true
}

// Top views the top element from the deque
// This is an alias for Peek but provides semantic clarity for stack
// use cases.
func (s *sliceDeque[X]) Top() (X, bool) {
	return s.Peek()
}

// Queue specifics

// Enqueue adds an element to the end of the deque
func (s *sliceDeque[X]) Enqueue(x X) {
	s.data = append(s.data, x)
}

// Dequeue removes an element from the front of the deque
func (s *sliceDeque[X]) Dequeue() (X, bool) {
	if len(s.data) == 0 {
		var zero X
		return zero, false
	}
	x := s.data[0]
	s.data = s.data[1:]
	return x, true
}

// Front views the first element of the deque
func (s *sliceDeque[X]) Front() (X, bool) {
	if len(s.data) == 0 {
		var zero X
		return zero, false
	}
	return s.data[0], true
}

// Back views the last element of the deque
func (s *sliceDeque[X]) Back() (X, bool) {
	return s.Peek()
}

// Others

// Empty returns if the deque has at least one element
func (s *sliceDeque[X]) Empty() bool {
	return len(s.data) == 0
}

// Size returns the total elements in the deque
func (s *sliceDeque[X]) Size() int {
	return len(s.data)
}

// Clear removes all elements from the deque
func (s *sliceDeque[X]) Clear() {
	s.data = nil
}

// Contains checks if the element exists in the deque
func (s *sliceDeque[X]) Contains(x X) bool {
	for _, v := range s.data {
		if v == x {
			return true
		}
	}
	return false
}

// Reverse reverses the deque
func (s *sliceDeque[X]) Reverse() {
	for i, j := 0, len(s.data)-1; i < j; i, j = i+1, j-1 {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
}

func (s *sliceDeque[X]) ToSlice() []X {
	// Return a copy to prevent mutation
	return append([]X(nil), s.data...)
}

// Remove removes the first occurrence of an element in the deque
func (s *sliceDeque[X]) Remove(x X) bool {
	for i, v := range s.data {
		if v == x {
			s.data = append(s.data[:i], s.data[i+1:]...)
			return true
		}
	}

	return false
}
