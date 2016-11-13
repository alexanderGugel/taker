package taker

import "github.com/hashicorp/go-multierror"

// ForceSeries runs the supplied tasks in series.
// All tasks will be run, regardless of any errors.
// If there is an error, it will be of type *multierror.Error.
func ForceSeries(tasks ...Task) error {
	var result *multierror.Error
	for _, t := range tasks {
		if err := t.Run(); err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result.ErrorOrNil()
}
