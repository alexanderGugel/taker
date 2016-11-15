package taker

import "time"

// TimeoutError is a description of a timeout error that will be returned once
// a specified time duration has been exceeded.
type TimeoutError struct {
	Limit time.Duration
}

// Error returns an error string.
func (e *TimeoutError) Error() string {
	return "exceeded " + e.Limit.String() + " timeout"
}

// TimeoutTask represents a task that will return a TimeoutError once a
// specified time limit has been exceeded.
type TimeoutTask struct {
	// Limit is the duration that the task is allowed to take.
	Limit time.Duration
}

// Run starts the execution of the task and returns an error once the time limit
// has passed.
func (t *TimeoutTask) Run() error {
	time.Sleep(t.Limit)
	return &TimeoutError{t.Limit}
}

// Timeout sets a time limit on an asynchronous task.
// If the task takes longer than the specified duration, an error of type
// TimeoutError will be returned.
func Timeout(t Task, limit time.Duration) error {
	return Race(t, &TimeoutTask{limit})
}
