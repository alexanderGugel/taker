package taker

// Times runs the passed in task n times or as long as it succeeds.
func Times(n int, task Task) error {
	for i := 0; i < n; i++ {
		err := task.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
