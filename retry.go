package taker

// Retry runs the passed in task until it succeeds, but a total maximum of max
// times.
func Retry(task Task, max int) error {
	var err error
	for i := 0; i < max; i++ {
		err = task.Run()
		if err == nil {
			return nil
		}
	}
	return err
}
