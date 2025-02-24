package sync

import (
	"context"
)

// Future represents an async computation.
type Future[T any] struct {
	result T
	err    error
	done   chan struct{}
}

// NewFuture runs a function asynchronously and returns a Future.
func NewFuture[T any](fn func() (T, error)) *Future[T] {
	return NewFutureWithContext(
		context.Background(),
		func(_ context.Context) (T, error) {
			return fn()
		},
	)
}

// NewFutureWithContext runs a function asynchronously and returns a Future.
func NewFutureWithContext[T any](
	ctx context.Context,
	fn func(ctx context.Context) (T, error),
) *Future[T] {
	f := &Future[T]{done: make(chan struct{})}

	go func() {
		select {
		// Context canceled before function finishes
		case <-ctx.Done():
			f.err = ctx.Err()
			close(f.done)
			return
		default:
			f.result, f.err = fn(ctx)
			close(f.done)
			return
		}
	}()

	return f
}

// Get waits for the result.
func (f *Future[T]) Get() (T, error) {
	<-f.done
	return f.result, f.err
}
