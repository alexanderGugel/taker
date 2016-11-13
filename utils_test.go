package taker

import (
	"errors"
	"sync"
)

type TestTask struct {
	Err   error
	ID    int
	Done  bool
	Mutex sync.Mutex
}

func (t *TestTask) Run() error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	t.Done = true
	return t.Err
}

func NewTestTask(err error, id int) *TestTask {
	task := &TestTask{
		Err:   err,
		ID:    0,
		Done:  false,
		Mutex: sync.Mutex{},
	}
	return task
}

func NewLockedTestTask(err error, id int) *TestTask {
	task := NewTestTask(err, id)
	task.Mutex.Lock()
	return task
}

type TestRetryTask struct {
	Max int
}

func (t *TestRetryTask) Run() error {
	if t.Max == 0 {
		return nil
	}
	t.Max--
	return errors.New("retry " + string(t.Max))
}

func NewTestRetryTask(max int) *TestRetryTask {
	task := &TestRetryTask{max}
	return task
}

type TestCounterTask struct {
	Counter int
}

func (t *TestCounterTask) Run() error {
	if t.Counter == 0 {
		return errors.New("counter")
	}
	t.Counter--
	return nil
}

func NewTestCounterTask(counter int) *TestCounterTask {
	task := &TestCounterTask{counter}
	return task
}
