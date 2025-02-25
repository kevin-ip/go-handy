package sync

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkerPool_Submit_Success(t *testing.T) {
	ctx := context.Background()
	pool := NewWorkerPool(ctx, 2, 5)
	defer pool.Close()

	done := make(chan struct{})
	err := pool.Submit(func() {
		close(done)
	})

	require.NoError(t, err)

	select {
	case <-done:
		// Task executed successfully
	case <-time.After(time.Second):
		require.Fail(t, "Task did not execute in time")
	}
}

func TestWorkerPool_Submit_FullQueue(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	pool := NewWorkerPool(ctx, 1, 1) // 1 worker, 1 buffered task
	defer pool.Close()

	// Fill up the queue
	require.EventuallyWithT(t,
		func(c *assert.CollectT) {
			err := pool.Submit(func() {
				time.Sleep(100 * time.Millisecond)
			})
			require.ErrorContains(c, err, "task queue is full")
		},
		500*time.Millisecond,
		1*time.Millisecond,
	)
}

func TestWorkerPool_Close(t *testing.T) {
	ctx := context.Background()
	counter := 5
	pool := NewWorkerPool(ctx, 2, counter)
	var executed int32

	for i := 0; i < counter; i++ {
		index := i
		err := pool.Submit(func() {
			time.Sleep((time.Duration(index) * 10) * time.Millisecond)
			atomic.AddInt32(&executed, 1)
		})
		require.NoError(t, err)
	}

	pool.Close()
	require.Equal(t, int32(counter), atomic.LoadInt32(&executed), "All tasks should execute before close")
}

func TestWorkerPool_CloseImmediately(t *testing.T) {
	ctx := context.Background()
	pool := NewWorkerPool(ctx, 2, 5)

	err := pool.Submit(func() { time.Sleep(500 * time.Millisecond) })
	require.NoError(t, err)

	pool.CloseImmediately()

	// No new tasks should be accepted
	err = pool.Submit(func() {})
	require.Error(t, err)
	require.Equal(t, "worker pool has been closed", err.Error())
}

func TestWorkerPool_Submit_AfterClose(t *testing.T) {
	ctx := context.Background()
	pool := NewWorkerPool(ctx, 2, 5)
	pool.Close()

	err := pool.Submit(func() {})
	require.ErrorContains(t, err, "worker pool has been closed")
}

func TestWorkerPool_IsClosed(t *testing.T) {
	ctx := context.Background()
	pool := NewWorkerPool(ctx, 2, 5)
	require.False(t, pool.IsClosed())

	pool.Close()
	require.True(t, pool.IsClosed())
}
