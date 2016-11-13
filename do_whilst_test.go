package taker

import (
	"errors"
	"testing"
)

func TestDoWhilstError(t *testing.T) {
	test := func() bool {
		return true
	}
	task := NewTestTask(errors.New("task err"), 0)
	err := DoWhilst(task, test)
	if err == nil {
		t.Fatal("expected err")
	}
}

func TestDoWhilstFalseTest(t *testing.T) {
	test := func() bool {
		return false
	}
	task := NewTestTask(errors.New("task err"), 0)
	err := DoWhilst(task, test)
	if err == nil {
		t.Fatalf("expected err")
	}
	if !task.Done {
		t.Fatal("expected task to be run")
	}
}
