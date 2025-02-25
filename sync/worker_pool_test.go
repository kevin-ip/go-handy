package sync

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkerPool_Submit(t *testing.T) {
	t.Run("task submitted to empty queue should execute successfully", func(t *testing.T) {
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
	})

	t.Run("full queue should return an error", func(t *testing.T) {
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
	})

	t.Run("submit after close should return an error", func(t *testing.T) {
		ctx := context.Background()
		pool := NewWorkerPool(ctx, 2, 5)
		pool.Close()

		err := pool.Submit(func() {})
		require.ErrorContains(t, err, "worker pool has been closed")
	})
}

func TestWorkerPool_Close(t *testing.T) {
	t.Run("all tasks should be completed after close", func(t *testing.T) {
		t.Parallel()
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
	})

	t.Run("closing multiple times should be fine", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		pool := NewWorkerPool(ctx, 2, 2)

		closeCount := 5
		wg := &sync.WaitGroup{}
		wg.Add(closeCount)
		for i := 0; i < closeCount; i++ {
			go func() {
				defer wg.Done()
				pool.Close()
				require.True(t, pool.IsClosed())
			}()
		}

		wg.Wait()
		require.True(t, pool.IsClosed())
	})
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

func TestWorkerPool_IsClosed(t *testing.T) {
	ctx := context.Background()
	pool := NewWorkerPool(ctx, 2, 5)
	require.False(t, pool.IsClosed())

	pool.Close()
	require.True(t, pool.IsClosed())
}

func TestWorkerPool_Resize(t *testing.T) {
	t.Run("increase should create more workers", func(t *testing.T) {
		t.Parallel()
		taskBuffer := 5
		ctx := context.Background()
		pool := NewWorkerPool(ctx, 2, taskBuffer)
		defer pool.Close()

		require.Equal(t, 2, pool.Workers())

		wg := &sync.WaitGroup{}
		var executed int32

		wg.Add(taskBuffer)
		for i := 0; i < taskBuffer; i++ {
			index := i
			err := pool.Submit(func() {
				defer wg.Done()
				time.Sleep((time.Duration(index) * 10) * time.Millisecond)
				atomic.AddInt32(&executed, 1)
			})
			require.NoError(t, err)
		}

		go func() {
			// keep increasing the pool size
			for i := 10; i < 20; i++ {
				err := pool.Resize(i)
				require.NoError(t, err)
				require.Equal(t, i, pool.Workers())
			}
		}()

		wg.Wait()
		require.Equal(t, int32(taskBuffer), atomic.LoadInt32(&executed), "All tasks should execute during resize")
	})

	t.Run("decrease should create stop some workers", func(t *testing.T) {
		t.Parallel()
		taskBuffer := 10
		ctx := context.Background()
		pool := NewWorkerPool(ctx, 10, taskBuffer)
		defer pool.Close()

		require.Equal(t, 10, pool.Workers())

		wg := &sync.WaitGroup{}
		var executed int32

		wg.Add(taskBuffer)
		for i := 0; i < taskBuffer; i++ {
			index := i
			err := pool.Submit(func() {
				defer wg.Done()
				time.Sleep((time.Duration(index) * 10) * time.Millisecond)
				atomic.AddInt32(&executed, 1)
			})
			require.NoError(t, err)
		}

		go func() {
			// keep decreasing the pool size
			for i := 9; i > 2; i-- {
				err := pool.Resize(i)
				require.NoError(t, err)
				require.Equal(t, i, pool.Workers())
			}
		}()

		wg.Wait()
		require.Equal(t, int32(taskBuffer), atomic.LoadInt32(&executed), "All tasks should execute during resize")
	})

	t.Run("resize after close should return an error", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		pool := NewWorkerPool(ctx, 2, 5)
		pool.Close()

		err := pool.Resize(10)
		require.ErrorContains(t, err, "worker pool has been closed")
	})
}
