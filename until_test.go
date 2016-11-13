package taker

import (
	"errors"
	"testing"
)

func TestUntilFalseTest(t *testing.T) {
	test := func() bool {
		return true
	}
	task := NewTestTask(errors.New("task err"), 0)
	err := Until(test, task)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if task.Done {
		t.Fatal("expected task not to be run")
	}
}
