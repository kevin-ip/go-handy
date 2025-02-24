package sync

import (
	"context"
	"sync"
)

// Future represents an async computation.
type Future[T any] struct {
	result T
	err    error
	done   chan struct{}
	mu     sync.Mutex
}

// NewFuture runs a function asynchronously and returns a Future.
func NewFuture[T any](fn func() (T, error)) *Future[T] {
	f := &Future[T]{done: make(chan struct{})}
	go func() {
		f.mu.Lock()
		defer f.mu.Unlock()
		f.result, f.err = fn()
		close(f.done)
	}()
	return f
}

func NewFutureWithContext[T any](
	ctx context.Context,
	fn func(ctx context.Context) (T, error),
) *Future[T] {
	f := &Future[T]{done: make(chan struct{})}

	go func() {
		var res T
		var err error

		done := make(chan struct{})
		go func() {
			res, err = fn(ctx)
			close(done)
		}()

		select {
		case <-ctx.Done(): // Context canceled or timed out
			f.mu.Lock()
			defer f.mu.Unlock()
			f.err = ctx.Err()
			close(f.done)
		case <-done: // Function finished execution
			f.mu.Lock()
			defer f.mu.Unlock()
			f.result = res
			f.err = err
			close(f.done)
		}
	}()

	return f
}

// Get waits for the result.
func (f *Future[T]) Get() (T, error) {
	<-f.done
	return f.result, f.err
}
