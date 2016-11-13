package taker

import (
	"testing"
)

func TestInvalidLimitErrorError(t *testing.T) {
	limit := 123
	err := &InvalidLimitError{limit}
	if err.Limit != limit {
		t.Fatalf("expected %v to be %v", err.Limit, limit)
	}
	wantMsg := "invalid limit 123"
	if gotMsg := err.Error(); gotMsg != wantMsg {
		t.Fatalf("expected %v to be %v", gotMsg, wantMsg)
	}
}

func TestParallelLimitInvalidLimitError(t *testing.T) {
	task0 := NewLockedTestTask(nil, 0)
	task1 := NewLockedTestTask(nil, 1)
	task2 := NewLockedTestTask(nil, 2)
	tasks := []Task{task0, task1, task2}

	wantErr := &InvalidLimitError{0}

	if err := ParallelLimit(0, tasks...); *err.(*InvalidLimitError) != *wantErr {
		t.Fatalf("expected %v to be %v", err, wantErr)
	}
}

func TestParallelLimit(t *testing.T) {
	task0 := NewLockedTestTask(nil, 0)
	task1 := NewLockedTestTask(nil, 1)
	task2 := NewLockedTestTask(nil, 2)
	tasks := []Task{task0, task1, task2}

	go func() {
		task1.Mutex.Unlock()
		task0.Mutex.Unlock()
		task2.Mutex.Unlock()
	}()

	if err := ParallelLimit(1, tasks...); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
}
