package taker

// Until is the inverse of While. It runs the specified task until it returns
// true.
func Until(test Test, task Task) error {
	return Whilst(test.Negate(), task)
}
