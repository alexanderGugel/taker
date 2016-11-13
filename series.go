package taker

// Series runs the supplied tasks in series.
// If a task returns an error, the remaining tasks won't be run and the error
// returned.
func Series(tasks ...Task) error {
	for _, t := range tasks {
		if err := t.Run(); err != nil {
			return err
		}
	}
	return nil
}
