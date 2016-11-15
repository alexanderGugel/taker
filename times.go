package taker

// Times runs the passed in task n times. If the task returns an error, the
// execution will be stopped and the error returned.
func Times(n int, task Task) error {
	for i := 0; i < n; i++ {
		err := task.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
