package collection

// Deque stands for double ended queue
type Deque[X comparable] struct {
	data []X
}

// NewDeque creates a new stackqueue
func NewDeque[X comparable]() *Deque[X] {
	return &Deque[X]{}
}

// Stack specifics

// Push adds an element to the top of the stackqueue
func (s *Deque[X]) Push(x X) {
	s.data = append(s.data, x)
}

// Pop removes the top element from the stackqueue
func (s *Deque[X]) Pop() (X, bool) {
	if len(s.data) == 0 {
		var zero X
		return zero, false
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, true
}

// Peek views the top element from the stackqueue
func (s *Deque[X]) Peek() (X, bool) {
	if len(s.data) == 0 {
		var zero X
		return zero, false
	}
	x := s.data[len(s.data)-1]
	return x, true
}

// Top views the top element from the stackqueue
// This is an alias for Peek but provides semantic clarity for stack
// use cases.
func (s *Deque[X]) Top() (X, bool) {
	return s.Peek()
}

// Queue specifics

// Enqueue adds an element to the end of the stackqueue
func (s *Deque[X]) Enqueue(x X) {
	s.data = append(s.data, x)
}

// Dequeue removes an element from the front of the stackqueue
func (s *Deque[X]) Dequeue() (X, bool) {
	if len(s.data) == 0 {
		var zero X
		return zero, false
	}
	x := s.data[0]
	s.data = s.data[1:]
	return x, true
}

// Front views the first element of the stackqueue
func (s *Deque[X]) Front() (X, bool) {
	if len(s.data) == 0 {
		var zero X
		return zero, false
	}
	return s.data[0], true
}

// Front views the last element of the stackqueue
func (s *Deque[X]) Back() (X, bool) {
	return s.Peek()
}

// Others

// Empty returns if the stackqueue has at least one element
func (s *Deque[X]) Empty() bool {
	return len(s.data) == 0
}

// Size returns the total elements in the stackqueue
func (s *Deque[X]) Size() int {
	return len(s.data)
}

// Clear removes all elements from the stackqueue
func (s *Deque[X]) Clear() {
	s.data = nil
}

// Contains checks if the element exists in the stackqueue
func (s *Deque[X]) Contains(x X) bool {
	for _, v := range s.data {
		if v == x {
			return true
		}
	}
	return false
}

// Reverse reverses the stackqueue
func (s *Deque[X]) Reverse() {
	for i, j := 0, len(s.data)-1; i < j; i, j = i+1, j-1 {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
}

func (s *Deque[X]) ToSlice() []X {
	// Return a copy to prevent mutation
	return append([]X(nil), s.data...)
}

// IndexOf finds the index of the first occurrence of an element in the Deque.
func (s *Deque[X]) IndexOf(x X) (int, bool) {
	for i, v := range s.data {
		if v == x {
			return i, true
		}
	}
	return -1, false
}

// RemoveAt removes the element at a specific index
func (s *Deque[X]) RemoveAt(index int) (X, bool) {
	if index < 0 || index >= len(s.data) {
		var zero X
		return zero, false
	}
	x := s.data[index]
	s.data = append(s.data[:index], s.data[index+1:]...)
	return x, true
}
