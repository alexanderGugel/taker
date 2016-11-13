package taker

// DoWhilst runs the task as long as the test passes. This is the post-check
// version of whilst. This means the task will be run at least once.
func DoWhilst(task Task, test Test) error {
	for {
		err := task.Run()
		if err != nil {
			return err
		}
		if !test() {
			break
		}
	}
	return nil
}
