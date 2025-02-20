package sync

import (
	"context"
	"errors"
	"runtime"
	"sync"
)

type FanOutResult[Y any] struct {
	Result Y
	Err    error
}

// FanOut spawns a fixed number of workers to process tasks concurrently
func FanOut[X any, Y any](
	ctx context.Context,
	tasks <-chan X,
	workers int,
	workerFunc func(context.Context, X) (Y, error),
) <-chan FanOutResult[Y] {
	resultChan := make(chan FanOutResult[Y], 10)

	wg := &sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ctx, wg, tasks, resultChan, workerFunc)
	}

	// Wait for all go routines to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}

// FanIn merges multiple input channels into a single one
func FanIn[X any](ctx context.Context, channels ...<-chan X) <-chan X {
	outputChan := make(chan X, len(channels))
	wg := &sync.WaitGroup{}

	for _, channel := range channels {
		wg.Add(1)
		go func(ch <-chan X) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case val, ok := <-ch:
					if !ok {
						return
					}
					outputChan <- val
				}
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(outputChan)
	}()

	return outputChan
}

type concurrentMapSettings struct {
	goRoutineCount int
}

func newSettings(options []ConcurrentMapOption) *concurrentMapSettings {
	// Default settings
	settings := &concurrentMapSettings{
		goRoutineCount: runtime.NumCPU(),
	}
	for _, option := range options {
		option.Apply(settings)
	}
	return settings
}

type ConcurrentMapOption interface {
	Apply(*concurrentMapSettings)
}

func WithGoRoutineCount(count int) ConcurrentMapOption {
	return withGoRoutineCount(count)
}

type withGoRoutineCount int

func (w withGoRoutineCount) Apply(settings *concurrentMapSettings) {
	settings.goRoutineCount = int(w)
}

// ConcurrentMap applies the work function concurrently to each value
// in the input. The ordering of the result list is not guaranteed
// to be the same as the ordering of the input.
// Partial result may be returned if some of the work are able
// to complete successfully. If there is an error from one of the work,
// the result from the work will not be included in the return value.
func ConcurrentMap[X any, Y any](
	ctx context.Context,
	inputs []X,
	work func(context.Context, X) (Y, error),
	options ...ConcurrentMapOption,
) ([]Y, error) {
	settings := newSettings(options)

	inputChan := make(chan X, len(inputs))
	resultChan := FanOut[X, Y](
		ctx,
		inputChan,
		settings.goRoutineCount,
		work,
	)

	// Send each input to the input channel
	go func() {
		defer close(inputChan)
		for _, input := range inputs {
			select {
			case <-ctx.Done():
				return
			default:
				inputChan <- input
			}
		}
	}()

	results := make([]Y, 0, len(resultChan))
	errs := make([]error, 0, len(resultChan))
outputLoop:
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case result, ok := <-resultChan:
			// The chanel has been closed
			if !ok {
				break outputLoop
			}

			if result.Err != nil {
				errs = append(errs, result.Err)
			} else {
				results = append(results, result.Result)
			}
		}
	}

	if len(errs) > 0 {
		err := errors.Join(errs...)
		return results, err
	}
	return results, nil
}

func worker[X any, Y any](
	ctx context.Context,
	wg *sync.WaitGroup,
	inputChan <-chan X,
	resultChan chan<- FanOutResult[Y],
	work func(context.Context, X) (Y, error),
) {
	defer wg.Done()
	for input := range inputChan {
		select {
		case <-ctx.Done():
			return
		default:
			output, err := work(ctx, input)
			resultChan <- FanOutResult[Y]{Result: output, Err: err}
		}
	}
}
