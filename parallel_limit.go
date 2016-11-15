package taker

import (
	"github.com/hashicorp/go-multierror"
	"strconv"
)

// See http://www.golangpatterns.info/concurrency/semaphores
type empty struct{}

// Semaphore implements a semaphore using an empty channel with specified buffer
// size.
type Semaphore chan empty

// NewSemaphore creates a new semaphore of the specified capacity.
func NewSemaphore(n int) Semaphore {
	return make(chan empty, n)
}

// Up acquires n resources.
func (s Semaphore) Up(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

// Down releases n resources.
func (s Semaphore) Down(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

// InvalidLimitError represents a runtime error that occurs when passing an
// invalid limit to ParallelLimit.
type InvalidLimitError struct {
	Limit int
}

// Error returns an error string.
func (e *InvalidLimitError) Error() string {
	return "invalid limit " + strconv.Itoa(e.Limit)
}

// ParallelLimit runs the supplied tasks in parallel.
// The function returns once all tasks have been run.
// If there is an error, it will be of type *multierror.Error.
func ParallelLimit(limit int, tasks ...Task) error {
	if limit <= 0 {
		return &InvalidLimitError{limit}
	}

	errs := make(chan error)
	defer close(errs)

	sema := NewSemaphore(limit)
	sema.Up(limit)

	for _, t := range tasks {
		go func(t Task) {
			sema.Down(1)
			errs <- t.Run()
			sema.Up(1)
		}(t)
	}

	var result *multierror.Error
	for i := 0; i < len(tasks); i++ {
		if err := <-errs; err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result.ErrorOrNil()
}
