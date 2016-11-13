package taker

// Test implements a truth test to perform before each execution of the
// underlying task.
type Test func() bool

// Negate returns the inverse of the underlying test function.
func (t Test) Negate() Test {
	return func() bool {
		return !t()
	}
}

// Whilst runs the task as long as test passes and task succeeds.
func Whilst(test Test, task Task) error {
	for test() {
		err := task.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
