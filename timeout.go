package taker

import "time"

// A TimeoutError is a description of a timeout error that will be returned once
// a specified time duration has been exceeded.
type TimeoutError struct {
	Limit time.Duration
}

func (e *TimeoutError) Error() string {
	return "exceeded " + string(e.Limit) + " timeout"
}

type TimeoutTask struct {
	Limit time.Duration
}

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
