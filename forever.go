package taker

// Forever runs the passed in task until it returns an error. If it never
// returns an error, it will be run indefinitely.
func Forever(task Task) error {
	for {
		if err := task.Run(); err != nil {
			return err
		}
	}
}
