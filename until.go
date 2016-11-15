package taker

// Until is the inverse of While. It runs the specified task until it returns
// true or the task fails. If the task returns an error, the execution will be
// stopped and the error returned.
func Until(test Test, task Task) error {
	return Whilst(test.Negate(), task)
}
