package sync

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestFuture_Success checks if Future executes successfully
func TestFuture_Success(t *testing.T) {
	f := NewFuture(func() (int, error) {
		return 42, nil
	})

	result, err := f.Get()
	require.NoError(t, err)
	require.Equal(t, 42, result)
}

// TestFuture_Error ensures Future captures and returns errors properly
func TestFuture_Error(t *testing.T) {
	f := NewFuture(func() (int, error) {
		return 0, errors.New("something went wrong")
	})

	result, err := f.Get()
	require.ErrorContains(t, err, "something went wrong")
	require.Equal(t, 0, result)
}

// TestFuture_ImmediateGet ensures calling Get() multiple times doesn't block
func TestFuture_ImmediateGet(t *testing.T) {
	f := NewFuture(func() (int, error) {
		return 42, nil
	})

	// Call Get multiple times
	for i := 0; i < 5; i++ {
		result, err := f.Get()
		require.NoError(t, err)
		require.Equal(t, 42, result)
	}
}

// TestFuture_ConcurrentGet tests if multiple concurrent Get() calls work
func TestFuture_ConcurrentGet(t *testing.T) {
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
}

// TestFuture_ZeroValue tests Future with a zero-value type
func TestFuture_ZeroValue(t *testing.T) {
	f := NewFuture(func() (struct{}, error) {
		return struct{}{}, nil
	})

	result, err := f.Get()
	require.NoError(t, err)
	require.Equal(t, struct{}{}, result)
}

// TestFutureWithContext_Success checks if Future executes successfully
func TestFutureWithContext_Success(t *testing.T) {
	ctx := context.Background()
	f := NewFutureWithContext(ctx, func(_ context.Context) (int, error) {
		return 42, nil
	})

	result, err := f.Get()
	require.NoError(t, err)
	require.Equal(t, 42, result)
}

// TestFutureWithContext_Error ensures Future captures and returns errors properly
func TestFutureWithContext_Error(t *testing.T) {
	ctx := context.Background()
	f := NewFutureWithContext(ctx, func(_ context.Context) (int, error) {
		return 0, errors.New("something went wrong")
	})

	result, err := f.Get()
	require.ErrorContains(t, err, "something went wrong")
	require.Equal(t, 0, result)
}

// TestFutureWithContext_ImmediateGet ensures calling Get() multiple times doesn't block
func TestFutureWithContext_ImmediateGet(t *testing.T) {
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
}

// TestFutureWithContext_ConcurrentGet tests if multiple concurrent Get() calls work
func TestFutureWithContext_ConcurrentGet(t *testing.T) {
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
}

// TestFutureWithContext_ZeroValue tests Future with a zero-value type
func TestFutureWithContext_ZeroValue(t *testing.T) {
	ctx := context.Background()
	f := NewFutureWithContext(ctx, func(_ context.Context) (struct{}, error) {
		return struct{}{}, nil
	})

	result, err := f.Get()
	require.NoError(t, err)
	require.Equal(t, struct{}{}, result)
}

func TestFutureWithContext_CancelImmediately(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	f := NewFutureWithContext(ctx, func(ctx context.Context) (int, error) {
		return 42, nil
	})

	result, err := f.Get()
	require.Error(t, err)
	require.Equal(t, 0, result)
}
