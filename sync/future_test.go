package sync

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFuture(t *testing.T) {
	t.Run("Success checks if Future executes successfully", func(t *testing.T) {
		t.Parallel()
		f := NewFuture(func() (int, error) {
			return 42, nil
		})

		result, err := f.Get()
		require.NoError(t, err)
		require.Equal(t, 42, result)
	})

	// TestFuture_
	t.Run("Error ensures Future captures and returns errors properly", func(t *testing.T) {
		t.Parallel()
		f := NewFuture(func() (int, error) {
			return 0, errors.New("something went wrong")
		})

		result, err := f.Get()
		require.ErrorContains(t, err, "something went wrong")
		require.Equal(t, 0, result)
	})

	// TestFuture_ImmediateGet ensures calling Get() multiple times doesn't block
	t.Run("ImmediateGet ensures calling Get() multiple times doesn't block", func(t *testing.T) {
		t.Parallel()
		f := NewFuture(func() (int, error) {
			return 42, nil
		})

		// Call Get multiple times
		for i := 0; i < 5; i++ {
			result, err := f.Get()
			require.NoError(t, err)
			require.Equal(t, 42, result)
		}
	})

	t.Run("ConcurrentGet tests if multiple concurrent Get() calls work", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		goRoutineCount := 3
		inChan := make(chan struct{}, goRoutineCount)

		f := NewFuture(func() (int, error) {
			for count := 0; count < goRoutineCount; {
				select {
				case <-ctx.Done():
					return 0, errors.New("timeout")
				case <-inChan:
					count++
				}
			}
			return 42, nil
		})

		var wg sync.WaitGroup
		wg.Add(goRoutineCount)

		for i := 0; i < goRoutineCount; i++ {
			go func() {
				defer wg.Done()
				// Signal the future that this go routine is running
				inChan <- struct{}{}

				result, err := f.Get()
				require.NoError(t, err)
				require.Equal(t, 42, result)
			}()
		}

		wg.Wait()
	})

	// TestFuture_ZeroValue tests Future with a zero-value type
	t.Run("ZeroValue tests Future with a zero-value type", func(t *testing.T) {
		t.Parallel()
		f := NewFuture(func() (struct{}, error) {
			return struct{}{}, nil
		})

		result, err := f.Get()
		require.NoError(t, err)
		require.Equal(t, struct{}{}, result)
	})
}

func TestFutureWithContext(t *testing.T) {
	t.Run("Success checks if Future executes successfully", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		f := NewFutureWithContext(ctx, func(_ context.Context) (int, error) {
			return 42, nil
		})

		result, err := f.Get()
		require.NoError(t, err)
		require.Equal(t, 42, result)
	})

	t.Run("Error ensures Future captures and returns errors properly", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		f := NewFutureWithContext(ctx, func(_ context.Context) (int, error) {
			return 0, errors.New("something went wrong")
		})

		result, err := f.Get()
		require.ErrorContains(t, err, "something went wrong")
		require.Equal(t, 0, result)
	})

	t.Run("ImmediateGet ensures calling Get() multiple times doesn't block", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		f := NewFutureWithContext(ctx, func(_ context.Context) (int, error) {
			return 42, nil
		})

		// Call Get multiple times
		for i := 0; i < 5; i++ {
			result, err := f.Get()
			require.NoError(t, err)
			require.Equal(t, 42, result)
		}
	})

	t.Run("ConcurrentGet tests if multiple concurrent Get() calls work", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		goRoutineCount := 3
		inChan := make(chan struct{}, goRoutineCount)

		f := NewFutureWithContext(ctx, func(ctx context.Context) (int, error) {
			for count := 0; count < goRoutineCount; {
				select {
				case <-ctx.Done():
					return 0, errors.New("timeout")
				case <-inChan:
					count++
				}
			}
			return 42, nil
		})

		var wg sync.WaitGroup
		wg.Add(goRoutineCount)

		for i := 0; i < goRoutineCount; i++ {
			go func() {
				defer wg.Done()
				// Signal the future that this go routine is running
				inChan <- struct{}{}

				result, err := f.Get()
				require.NoError(t, err)
				require.Equal(t, 42, result)
			}()
		}

		wg.Wait()
	})

	t.Run("ZeroValue tests Future with a zero-value type", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		f := NewFutureWithContext(ctx, func(_ context.Context) (struct{}, error) {
			return struct{}{}, nil
		})

		result, err := f.Get()
		require.NoError(t, err)
		require.Equal(t, struct{}{}, result)
	})

	t.Run("CancelImmediately tests the func should not been invoked if context has been canceled", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithCancel(context.Background())
		// cancel immediately
		cancel()

		f := NewFutureWithContext(ctx, func(ctx context.Context) (int, error) {
			return 42, nil
		})

		result, err := f.Get()
		require.ErrorContains(t, err, "context canceled")
		require.Equal(t, 0, result)
	})
}
