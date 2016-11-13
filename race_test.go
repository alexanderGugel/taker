package taker

import (
	"errors"
	"sync"
	"testing"
	"time"
)

func TestRaceError(t *testing.T) {
	task0 := NewLockedTestTask(errors.New("task0"), 0)
	task1 := NewLockedTestTask(errors.New("task1"), 1)
	task2 := NewLockedTestTask(errors.New("task2"), 2)
	tasks := []Task{task0, task1, task2}

	go func() {
		time.Sleep(time.Millisecond * 10)
		task1.Mutex.Unlock()
		time.Sleep(time.Millisecond * 10)
		task0.Mutex.Unlock()
		time.Sleep(time.Millisecond * 10)
		task2.Mutex.Unlock()
	}()

	err := Race(tasks...)
	if err != task1.Err {
		t.Fatalf("expected %v to be %v", err, task1.Err)
	}
}

func TestRaceSuccess(t *testing.T) {
	task0 := &TestTask{nil, 0, false, sync.Mutex{}}
	task1 := &TestTask{nil, 1, false, sync.Mutex{}}
	task2 := &TestTask{nil, 2, false, sync.Mutex{}}
	tasks := []Task{task0, task1, task2}

	task0.Mutex.Lock()
	task1.Mutex.Lock()
	task2.Mutex.Lock()

	go func() {
		time.Sleep(time.Millisecond * 10)
		task1.Mutex.Unlock()
		time.Sleep(time.Millisecond * 10)
		task0.Mutex.Unlock()
		time.Sleep(time.Millisecond * 10)
		task2.Mutex.Unlock()
	}()

	if err := Race(tasks...); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
}
