package taker

// Test is a truth test that will be performed before each execution of the
// underlying task.
type Test func() bool

// Negate returns the inverse of the underlying test function. If the test
// returns true, Negate would return false and vice-versa.
func (t Test) Negate() Test {
	return func() bool {
		return !t()
	}
}

// Whilst repeatedly runs the task as long as test returns true and the task
// succeeds. If the task returns an error, the execution will be stopped and the
// error returned.
func Whilst(test Test, task Task) error {
	for test() {
		err := task.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
