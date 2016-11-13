package taker

import (
	"errors"
	"testing"
)

func TestTimesSuccess(t *testing.T) {
	task := NewTestTask(nil, 0)
	err := Times(5, task)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
}

func TestTimesError(t *testing.T) {
	task := NewTestTask(errors.New("task err"), 0)
	err := Times(5, task)
	if err != task.Err {
		t.Fatalf("expected %v to be %v", err, task.Err)
	}
}
