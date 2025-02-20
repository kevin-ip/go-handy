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

func TestFanOut(t *testing.T) {
	t.Run("concurrent multiplication should work", func(t *testing.T) {
		ctx := context.Background()
		tasks := make(chan int, 10)

		responseChan := FanOut[int, int](ctx, tasks, 2, timesTwo)
		go func() {
			defer close(tasks)
			for i := 0; i < 10; i++ {
				tasks <- i
			}
		}()

		result, errors := gatherResponse(responseChan)
		sort.Ints(result)
		require.Equal(t, []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}, result)
		require.Empty(t, errors)
	})

	t.Run("concurrent multiplication with partial failiure should work", func(t *testing.T) {
		ctx := context.Background()
		tasks := make(chan int, 10)

		responseChan := FanOut[int, int](ctx, tasks, 2, funkyTimesTwo)
		go func() {
			defer close(tasks)
			for i := 0; i < 10; i++ {
				tasks <- i
			}
		}()

		result, errors := gatherResponse(responseChan)
		sort.Ints(result)
		require.Equal(t, []int{0, 4, 8, 12, 16}, result)
		require.Len(t, errors, 5)
	})
}

func gatherResponse(responseChan <-chan FanOutResult[int]) ([]int, []error) {
	result := []int{}
	errors := []error{}
	for response := range responseChan {
		if response.Err != nil {
			errors = append(errors, response.Err)
		} else {
			result = append(result, response.Result)
		}

	}
	return result, errors
}

func TestFanIn(t *testing.T) {
	t.Run("fan-in should merge multiple channels into one", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		channels := make([]chan int, 10)
		for i := range channels {
			channels[i] = make(chan int)
		}

		// Start sending data to each channel
		for i, channel := range channels {
			go func(ch chan int, id int) {
				defer close(ch)
				for j := 0; j < i; j++ {
					ch <- j
				}
			}(channel, i)
		}

		readOnlyChannels := make([]<-chan int, len(channels))
		for i, channel := range channels {
			readOnlyChannels[i] = channel
		}

		outChan := FanIn[int](ctx, readOnlyChannels...)

		result := 0
		for out := range outChan {
			result = result + out
		}
		require.Equal(t, 120, result)
	})

	t.Run("context cancel before running should stop the operation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		// Cancel immediately
		cancel()

		channels := make([]chan int, 10)
		for i := range channels {
			channels[i] = make(chan int)
		}

		// Start sending data to each channel
		for i, channel := range channels {
			go func(ch chan int, id int) {
				defer close(ch)
				for j := 0; j < i; j++ {
					ch <- j
				}
			}(channel, i)
		}

		readOnlyChannels := make([]<-chan int, len(channels))
		for i, channel := range channels {
			readOnlyChannels[i] = channel
		}

		outChan := FanIn[int](ctx, readOnlyChannels...)

		result := 0
		for out := range outChan {
			result = result + out
		}
		require.Equal(t, 0, result)
	})

	t.Run("should return if no input channel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		outChan := FanIn[int](ctx)

		result := 0
		for out := range outChan {
			result = result + out
		}
		require.Equal(t, 0, result)
	})
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
			for range isRunningChan {
				cancel()
				return
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
