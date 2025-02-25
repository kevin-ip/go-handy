package sync

import (
	"context"
	"errors"
	"sync"
)

type Task func()

// WorkerPool provides a simple worker pool implementation,
// allowing tasks to be executed concurrently with a fixed number of worker goroutines.
// It supports graceful shutdowns and immediate termination of workers.
type WorkerPool struct {
	tasks      chan Task
	wg         sync.WaitGroup
	ctx        context.Context
	cancelFunc context.CancelFunc
	isClosed   bool
	lock       sync.RWMutex
	numWorkers int
	doneChan   chan struct{}
}

// NewWorkerPool creates a worker pool.
func NewWorkerPool(ctx context.Context, numWorkers int, taskBuffer int) *WorkerPool {
	ctx, cancel := context.WithCancel(ctx)

	pool := &WorkerPool{
		tasks:      make(chan Task, taskBuffer),
		ctx:        ctx,
		cancelFunc: cancel,
		numWorkers: numWorkers,
		doneChan:   make(chan struct{}, numWorkers),
	}
	pool.start(numWorkers)
	return pool
}

// start kicks off the fixed number of worker goroutines.
func (p *WorkerPool) start(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go p.worker()
	}
}

func (p *WorkerPool) worker() {
	for {
		select {
		case <-p.ctx.Done():
			return
		case <-p.doneChan:
			return
		case task, ok := <-p.tasks:
			if !ok {
				return
			}
			defer p.wg.Done()
			task()
		}
	}
}

// Submit submits a task into the pool
// If the task queue is full, Submit returns an error instead of blocking.
// Client can retry some time later or report an error.
// If the pool is closed, Submit returns an error.
func (p *WorkerPool) Submit(task Task) error {
	if p.IsClosed() {
		return errors.New("worker pool has been closed")
	}

	p.wg.Add(1)
	select {
	case p.tasks <- task:
		return nil
	default:
		// If the channel is full, return an error.
		// Manually call Done if we can't add to the channel.
		p.wg.Done()
		return errors.New("task queue is full")
	}
}

func (p *WorkerPool) Resize(numWorkers int) error {
	if p.IsClosed() {
		return errors.New("worker pool has been closed")
	}

	p.lock.Lock()
	defer p.lock.Unlock()
	if p.isClosed {
		return errors.New("worker pool has been closed")
	}

	if p.numWorkers < numWorkers {
		// increase worker go-routines
		for i := 0; i < (numWorkers - p.numWorkers); i++ {
			go p.worker()
		}
	} else if p.numWorkers > numWorkers {
		// decrease worker go-routines
		for i := 0; i < (p.numWorkers - numWorkers); i++ {
			p.doneChan <- struct{}{}
		}
	}

	p.numWorkers = numWorkers
	return nil
}

func (p *WorkerPool) Workers() int {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.numWorkers
}

// IsClosed returns whether this pool is closed
func (p *WorkerPool) IsClosed() bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.isClosed
}

// Close closes the worker pool while waiting for the tasks
// in progress to complete
func (p *WorkerPool) Close() {
	if !p.close() {
		return
	}

	p.wg.Wait()
}

// CloseImmediately closes the worker pool without waiting
// for the tasks in progress to complete
func (p *WorkerPool) CloseImmediately() {
	p.cancelFunc()
	p.close()
}

// close closes the worker pool
// true if closed successfully, false otherwise
func (p *WorkerPool) close() bool {
	if p.IsClosed() {
		return false
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	if p.isClosed {
		return false
	}

	p.isClosed = true
	close(p.tasks)
	return true
}
