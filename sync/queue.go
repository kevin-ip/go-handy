package sync

import (
	"sync"

	"github.com/kevin-ip/go-handy/collection"
)

// ConcurrentQueue is a queue backed by two internal queues
// so that inserting and removing an element
// can be performed concurrently.
type ConcurrentQueue[X comparable] struct {
	inLock  sync.RWMutex
	outLock sync.RWMutex

	inDeque  collection.Deque[X]
	outDeque collection.Deque[X]
}

func NewConcurrentQueue[X comparable]() *ConcurrentQueue[X] {
	return &ConcurrentQueue[X]{
		inDeque:  collection.NewSliceDeque[X](),
		outDeque: collection.NewSliceDeque[X](),
	}
}

// Enqueue adds an element to the back of the queue
func (q *ConcurrentQueue[X]) Enqueue(x X) {
	q.inLock.Lock()
	defer q.inLock.Unlock()
	q.inDeque.Enqueue(x)
}

// Dequeue removes an element from the front of the queue
func (q *ConcurrentQueue[X]) Dequeue() (X, bool) {
	// having lock acquired in methods to improve readability over performance
	x, ok := q.dequeueFromOutDeque()
	if ok {
		return x, true
	}
	q.fillOutDeque()
	return q.dequeueFromOutDeque()
}

func (q *ConcurrentQueue[X]) dequeueFromOutDeque() (X, bool) {
	q.outLock.Lock()
	defer q.outLock.Unlock()
	return q.outDeque.Dequeue()
}

func (q *ConcurrentQueue[X]) fillOutDeque() {
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
func (q *ConcurrentQueue[X]) Front() (X, bool) {
	x, ok := q.front(q.outDeque, &q.outLock)
	if ok {
		return x, true
	}
	return q.front(q.inDeque, &q.inLock)
}

func (q *ConcurrentQueue[X]) front(deque collection.Deque[X], lock *sync.RWMutex) (X, bool) {
	lock.RLock()
	defer lock.RUnlock()
	return deque.Front()
}

// Back views the last element of the queue
// a zero value and a false if the queue is empty
func (q *ConcurrentQueue[X]) Back() (X, bool) {
	// check the inDeque first as new element enqueues there
	x, ok := q.back(q.inDeque, &q.inLock)
	if ok {
		return x, true
	}
	return q.back(q.outDeque, &q.outLock)
}

func (q *ConcurrentQueue[X]) back(deque collection.Deque[X], lock *sync.RWMutex) (X, bool) {
	lock.RLock()
	defer lock.RUnlock()
	return deque.Back()
}

// Empty returns true if the queue has zero element
func (q *ConcurrentQueue[X]) Empty() bool {
	q.inLock.RLock()
	defer q.inLock.RUnlock()
	q.outLock.RLock()
	defer q.outLock.RUnlock()
	return q.inDeque.Empty() && q.outDeque.Empty()
}

// Size returns the total number of elements in the queue.
func (q *ConcurrentQueue[X]) Size() int {
	q.inLock.RLock()
	defer q.inLock.RUnlock()
	q.outLock.RLock()
	defer q.outLock.RUnlock()
	return q.inDeque.Size() + q.outDeque.Size()
}

// Clear resets the queue.
func (q *ConcurrentQueue[X]) Clear() {
	q.inLock.Lock()
	defer q.inLock.Unlock()
	q.outLock.Lock()
	defer q.outLock.Unlock()
	q.inDeque.Clear()
	q.outDeque.Clear()
}

// Contains checks if the element exists in the deque
func (q *ConcurrentQueue[X]) Contains(x X) bool {
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
func (q *ConcurrentQueue[X]) Reverse() {
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
func (q *ConcurrentQueue[X]) ToSlice() []X {
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
func (q *ConcurrentQueue[X]) Remove(x X) bool {
	if q.remove(q.inDeque, &q.inLock, x) {
		return true
	}
	return q.remove(q.outDeque, &q.outLock, x)
}

func (q *ConcurrentQueue[X]) remove(deque collection.Deque[X], lock *sync.RWMutex, x X) bool {
	lock.Lock()
	defer lock.Unlock()

	return deque.Remove(x)
}
