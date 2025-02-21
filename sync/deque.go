package sync

import (
	"sync"

	"github.com/kevin-ip/go-handy/collection"
)

// ConcurrentDeque is a queue backed by two internal queues
// so that inserting and removing an element
// can be performed concurrently.
type ConcurrentDeque[X comparable] struct {
	inLock  sync.RWMutex
	outLock sync.RWMutex

	inDeque  collection.Deque[X]
	outDeque collection.Deque[X]
}

// ConcurrentDeque backed by slice deque
func NewConcurrentSliceDeque[X comparable]() collection.Deque[X] {
	return &ConcurrentDeque[X]{
		inDeque:  collection.NewSliceDeque[X](),
		outDeque: collection.NewSliceDeque[X](),
	}
}

// ConcurrentDeque backed by linked deque
func NewConcurrentLinkedDeque[X comparable]() collection.Deque[X] {
	return &ConcurrentDeque[X]{
		inDeque:  collection.NewLinkedDeque[X](),
		outDeque: collection.NewLinkedDeque[X](),
	}
}

func (q *ConcurrentDeque[X]) Push(x X) {
	q.inLock.Lock()
	defer q.inLock.Unlock()

	q.inDeque.Push(x)
}

func (q *ConcurrentDeque[X]) Pop() (X, bool) {
	value, ok := q.pop(q.inDeque, &q.inLock)
	if ok {
		return value, true
	}
	return q.pop(q.outDeque, &q.outLock)
}

func (q *ConcurrentDeque[X]) pop(deque collection.Deque[X], lock *sync.RWMutex) (X, bool) {
	lock.Lock()
	defer lock.Unlock()

	return deque.Pop()
}

func (q *ConcurrentDeque[X]) Peek() (X, bool) {
	value, ok := q.peek(q.inDeque, &q.inLock)
	if ok {
		return value, true
	}
	return q.peek(q.outDeque, &q.outLock)
}

func (q *ConcurrentDeque[X]) peek(deque collection.Deque[X], lock *sync.RWMutex) (X, bool) {
	lock.RLock()
	defer lock.RUnlock()

	return deque.Peek()
}

func (q *ConcurrentDeque[X]) Top() (X, bool) {
	return q.Peek()
}

// Enqueue adds an element to the back of the queue
func (q *ConcurrentDeque[X]) Enqueue(x X) {
	q.inLock.Lock()
	defer q.inLock.Unlock()
	q.inDeque.Enqueue(x)
}

// Dequeue removes an element from the front of the queue
func (q *ConcurrentDeque[X]) Dequeue() (X, bool) {
	// having lock acquired in methods to improve readability over performance
	x, ok := q.dequeueFromOutDeque()
	if ok {
		return x, true
	}
	q.fillOutDeque()
	return q.dequeueFromOutDeque()
}

func (q *ConcurrentDeque[X]) dequeueFromOutDeque() (X, bool) {
	q.outLock.Lock()
	defer q.outLock.Unlock()
	return q.outDeque.Dequeue()
}

func (q *ConcurrentDeque[X]) fillOutDeque() {
	q.inLock.Lock()
	elements := q.inDeque.ToSlice()
	q.inDeque.Clear()
	q.inLock.Unlock()

	q.outLock.Lock()
	defer q.outLock.Unlock()
	for _, value := range elements {
		q.outDeque.Enqueue(value)
	}
}

// Front views the first element of the queue
// a zero value and a false if the queue is empty
func (q *ConcurrentDeque[X]) Front() (X, bool) {
	x, ok := q.front(q.outDeque, &q.outLock)
	if ok {
		return x, true
	}
	return q.front(q.inDeque, &q.inLock)
}

func (q *ConcurrentDeque[X]) front(deque collection.Deque[X], lock *sync.RWMutex) (X, bool) {
	lock.RLock()
	defer lock.RUnlock()
	return deque.Front()
}

// Back views the last element of the queue
// a zero value and a false if the queue is empty
func (q *ConcurrentDeque[X]) Back() (X, bool) {
	// check the inDeque first as new element enqueues there
	x, ok := q.back(q.inDeque, &q.inLock)
	if ok {
		return x, true
	}
	return q.back(q.outDeque, &q.outLock)
}

func (q *ConcurrentDeque[X]) back(deque collection.Deque[X], lock *sync.RWMutex) (X, bool) {
	lock.RLock()
	defer lock.RUnlock()
	return deque.Back()
}

// Empty returns true if the queue has zero element
func (q *ConcurrentDeque[X]) Empty() bool {
	q.inLock.RLock()
	defer q.inLock.RUnlock()
	q.outLock.RLock()
	defer q.outLock.RUnlock()
	return q.inDeque.Empty() && q.outDeque.Empty()
}

// Size returns the total number of elements in the queue.
func (q *ConcurrentDeque[X]) Size() int {
	q.inLock.RLock()
	defer q.inLock.RUnlock()
	q.outLock.RLock()
	defer q.outLock.RUnlock()
	return q.inDeque.Size() + q.outDeque.Size()
}

// Clear resets the queue.
func (q *ConcurrentDeque[X]) Clear() {
	q.inLock.Lock()
	defer q.inLock.Unlock()
	q.outLock.Lock()
	defer q.outLock.Unlock()
	q.inDeque.Clear()
	q.outDeque.Clear()
}

// Contains checks if the element exists in the deque
func (q *ConcurrentDeque[X]) Contains(x X) bool {
	q.inLock.RLock()
	defer q.inLock.RUnlock()
	if result := q.inDeque.Contains(x); result {
		return true
	}

	q.outLock.RLock()
	defer q.outLock.RUnlock()
	return q.outDeque.Contains(x)
}

// Reverse reverses the queue
func (q *ConcurrentDeque[X]) Reverse() {
	// acquiring both inLock and outLock to ensure
	// no write operation is interfering the reverse
	q.inLock.Lock()
	defer q.inLock.Unlock()
	q.outLock.Lock()
	defer q.outLock.Unlock()

	for _, value := range q.inDeque.ToSlice() {
		q.outDeque.Enqueue(value)
	}
	q.inDeque.Clear()

	q.outDeque.Reverse()
}

// ToSlice creates a snapshot of all the elements in the queue
func (q *ConcurrentDeque[X]) ToSlice() []X {
	q.inLock.RLock()
	inSlice := q.inDeque.ToSlice()
	q.inLock.RUnlock()

	q.outLock.RLock()
	outSlice := q.outDeque.ToSlice()
	q.outLock.RUnlock()

	return append(inSlice, outSlice...)
}

// Remove removes the element from the queue
// true if removed successfully, false otherwise
func (q *ConcurrentDeque[X]) Remove(x X) bool {
	if q.remove(q.inDeque, &q.inLock, x) {
		return true
	}
	return q.remove(q.outDeque, &q.outLock, x)
}

func (q *ConcurrentDeque[X]) remove(deque collection.Deque[X], lock *sync.RWMutex, x X) bool {
	lock.Lock()
	defer lock.Unlock()

	return deque.Remove(x)
}
