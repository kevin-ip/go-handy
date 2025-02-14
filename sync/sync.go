package sync

import (
	"context"
	"errors"
	"runtime"
	"sync"
)

func FanOut[X any, Y any](
	ctx context.Context,
	tasks <-chan X,
	workers int,
	workerFunc func(context.Context, X) (Y, error),
) (<-chan Y, <-chan error) {
	outputChan := make(chan Y, 10)
	errorChan := make(chan error, 10)

	wg := &sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ctx, wg, tasks, outputChan, errorChan, workerFunc)
	}

	// Wait for all go routines to finish
	go func() {
		wg.Wait()
		close(outputChan)
		close(errorChan)
	}()

	return outputChan, errorChan
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
// to complete successfully.
func ConcurrentMap[X any, Y any](
	ctx context.Context,
	inputs []X,
	work func(context.Context, X) (Y, error),
	options ...ConcurrentMapOption,
) ([]Y, error) {
	settings := newSettings(options)

	inputChan := make(chan X, len(inputs))
	outputChan, errorChan := FanOut[X, Y](
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

	result := make([]Y, 0, len(outputChan))
outputLoop:
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case output, ok := <-outputChan:
			// The chanel has been closed
			if !ok {
				break outputLoop
			}
			result = append(result, output)
		}
	}

	errs := make([]error, 0, len(errorChan))
errorLoop:
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case err, ok := <-errorChan:
			// The chanel has been closed
			if !ok {
				break errorLoop
			}
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		err := errors.Join(errs...)
		return result, err
	}
	return result, nil
}

func worker[X any, Y any](
	ctx context.Context,
	wg *sync.WaitGroup,
	inputChan <-chan X,
	outputChan chan<- Y,
	errorChan chan<- error,
	work func(context.Context, X) (Y, error),
) {
	defer wg.Done()
	for input := range inputChan {
		select {
		case <-ctx.Done():
			return
		default:
			output, err := work(ctx, input)
			if err != nil {
				errorChan <- err
			} else {
				outputChan <- output
			}
		}
	}
}
