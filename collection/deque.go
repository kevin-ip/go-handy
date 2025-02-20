package collection

type Deque[X comparable] interface {

	// Push adds an element to the top of the deque
	Push(x X)

	// Pop removes the top element from the deque
	Pop() (X, bool)

	// Peek views the top element from the deque
	Peek() (X, bool)

	// Top views the top element from the deque
	// This is an alias for Peek but provides semantic clarity for stack
	// use cases.
	Top() (X, bool)

	// Queue specifics

	// Enqueue adds an element to the end of the deque
	Enqueue(x X)

	// Dequeue removes an element from the front of the deque
	Dequeue() (X, bool)

	// Front views the first element of the deque
	Front() (X, bool)

	// Front views the last element of the deque
	Back() (X, bool)

	// Others

	// Empty returns if the deque has at least one element
	Empty() bool

	// Size returns the total elements in the deque
	Size() int

	// Clear removes all elements from the deque
	Clear()

	// Contains checks if the element exists in the deque
	Contains(x X) bool

	// Reverse reverses the deque
	Reverse()

	ToSlice() []X

	// IndexOf finds the index of the first occurrence of an element in the sliceDeque.
	IndexOf(x X) (int, bool)

	// RemoveAt removes the element at a specific index
	RemoveAt(index int) (X, bool)
}
