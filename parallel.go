package taker

import "github.com/hashicorp/go-multierror"

// Parallel runs the supplied tasks in parallel.
// The function returns once all tasks have been run.
// If there is an error, it will be of type *multierror.Error.
func Parallel(tasks ...Task) error {
	errs := make(chan error)
	defer close(errs)

	for _, t := range tasks {
		go func(t Task) { errs <- t.Run() }(t)
	}

	var result *multierror.Error
	for i := 0; i < len(tasks); i++ {
		if err := <-errs; err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result.ErrorOrNil()
}
