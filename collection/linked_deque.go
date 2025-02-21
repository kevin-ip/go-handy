package collection

type linkedNode[X comparable] struct {
	x    X
	next *linkedNode[X]
	prev *linkedNode[X]
}

type LinkedDeque[X comparable] struct {
	head *linkedNode[X]
	last *linkedNode[X]
	size int
}

func NewLinkedDeque[X comparable]() Deque[X] {
	return &LinkedDeque[X]{}
}

// Push adds an element to the top of the deque
func (d *LinkedDeque[X]) Push(x X) {
	node := &linkedNode[X]{x: x}

	d.size += 1
	if d.head == nil {
		d.head = node
		d.last = node
		return
	}

	currLast := d.last
	node.prev = currLast
	currLast.next = node

	d.last = node
}

// Pop removes the top element from the deque
func (d *LinkedDeque[X]) Pop() (X, bool) {
	if d.last == nil {
		var zero X
		return zero, false
	}

	node := d.last
	if node.prev != nil {
		d.last = node.prev
		d.last.next = nil
	} else {
		d.last = nil
	}
	d.size -= 1

	return node.x, true
}

// Peek views the top element from the deque
func (d *LinkedDeque[X]) Peek() (X, bool) {
	if d.last == nil {
		var zero X
		return zero, false
	}

	return d.last.x, true
}

// Top views the top element from the deque
// This is an alias for Peek but provides semantic clarity for stack
// use cases.
func (d *LinkedDeque[X]) Top() (X, bool) {
	return d.Peek()
}

// Enqueue adds an element to the end of the deque
func (d *LinkedDeque[X]) Enqueue(x X) {
	d.Push(x)
}

// Dequeue removes an element from the front of the deque
func (d *LinkedDeque[X]) Dequeue() (X, bool) {
	if d.head == nil {
		var zero X
		return zero, false
	}

	currHead := d.head
	d.head = d.head.next
	d.head.prev = nil
	d.size -= 1

	return currHead.x, true
}

// Front views the first element of the deque
func (d *LinkedDeque[X]) Front() (X, bool) {
	if d.head == nil {
		var zero X
		return zero, false
	}

	return d.head.x, true
}

// Back views the last element of the deque
func (d *LinkedDeque[X]) Back() (X, bool) {
	if d.last == nil {
		var zero X
		return zero, false
	}

	return d.last.x, true
}

// Empty returns if the deque has at least one element
func (d *LinkedDeque[X]) Empty() bool {
	return d.head == nil
}

// Size returns the total elements in the deque
func (d *LinkedDeque[X]) Size() int {
	return d.size
}

// Clear removes all elements from the deque
func (d *LinkedDeque[X]) Clear() {
	d.head = nil
	d.last = nil
	d.size = 0
}

// Contains checks if the element exists in the deque
func (d *LinkedDeque[X]) Contains(x X) bool {
	for node := d.head; node != nil; node = node.next {
		if node.x == x {
			return true
		}
	}
	return false
}

// Reverse reverses the deque
func (d *LinkedDeque[X]) Reverse() {
	if d.last == nil {
		return
	}

	currHead, currLast := d.head, d.last
	for currHead != currLast && currHead.prev != currLast {
		currHead.x, currLast.x = currLast.x, currHead.x
		currHead, currLast = currHead.next, currLast.prev
	}
}

func (d *LinkedDeque[X]) ToSlice() []X {
	slice := []X{}
	for node := d.head; node != nil; node = node.next {
		slice = append(slice, node.x)
	}
	return slice
}

// Remove removes the first occurrence of an element in the deque
func (d *LinkedDeque[X]) Remove(x X) bool {
	for node := d.head; node != nil; node = node.next {
		if node.x == x {
			node.prev.next = node.next
			node.next.prev = node.prev
			d.size -= 1
			return true
		}
	}
	return false
}
