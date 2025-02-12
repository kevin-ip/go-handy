package sync

import (
	"context"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func timesTwo(_ context.Context, input int) (int, error) {
	return input * 2, nil
}

func funkyTimesTwo(_ context.Context, input int) (int, error) {
	if input%2 == 0 {
		return input * 2, nil
	} else {
		return 0, fmt.Errorf("some-error: %v", input)
	}
}

func totalFailure(_ context.Context, input int) (int, error) {
	return 0, fmt.Errorf("some-error: %v", input)
}

func TestConcurrentMap(t *testing.T) {
	t.Run("concurrent multiplication should work", func(t *testing.T) {
		ctx := context.Background()
		actual, err := ConcurrentMap[int, int](ctx, []int{1, 2, 3, 4, 5}, timesTwo)
		require.NoError(t, err)
		sort.Ints(actual)
		require.Equal(t, []int{2, 4, 6, 8, 10}, actual)
	})

	t.Run("concurrent multiplication with partial error", func(t *testing.T) {
		ctx := context.Background()
		actual, err := ConcurrentMap[int, int](ctx, []int{1, 2, 3, 4, 5}, funkyTimesTwo)
		require.Error(t, err)
		sort.Ints(actual)
		require.Equal(t, []int{4, 8}, actual)
	})

	t.Run("concurrent with all errors", func(t *testing.T) {
		ctx := context.Background()
		actual, err := ConcurrentMap[int, int](ctx, []int{1, 2, 3, 4, 5}, totalFailure)
		require.Error(t, err)
		require.Len(t, actual, 0)
	})

	t.Run("context cancel before running should work", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		// cancel immediate before everything starts
		cancel()

		inputs := []int{}
		for i := 0; i < 10; i++ {
			inputs = append(inputs, i)
		}

		out, err := ConcurrentMap[int, int](
			ctx,
			inputs,
			timesTwo,
			WithGoRoutineCount(len(inputs)),
		)
		require.ErrorContainsf(t, err, context.Canceled.Error(), "out: %v; err: %v", out, err)
	})

	t.Run("context cancel after running should work", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		inputs := []int{}
		for i := 0; i < 10; i++ {
			inputs = append(inputs, i)
		}
		isRunningChan := make(chan interface{}, len(inputs))

		go func() {
			for {
				select {
				case <-isRunningChan:
					cancel()
					return
				}
			}
		}()

		out, err := ConcurrentMap[int, int](
			ctx,
			inputs,
			func(_ context.Context, input int) (int, error) {
				isRunningChan <- struct{}{}
				time.Sleep(1 * time.Second)
				return input * 2, nil
			},
			WithGoRoutineCount(len(inputs)),
		)
		require.ErrorContainsf(t, err, context.Canceled.Error(), "out: %v; err: %v", out, err)
	})
}
