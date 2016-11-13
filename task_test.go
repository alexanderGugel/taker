package taker

import (
	"errors"
	"testing"
)

func TestWrapSuccess(t *testing.T) {
	task := Wrap(func() error {
		return nil
	})
	if err := task.Run(); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
}

func TestWrapError(t *testing.T) {
	runErr := errors.New("run err")
	task := Wrap(func() error {
		return runErr
	})
	if err := task.Run(); err != runErr {
		t.Fatalf("expected %v to be %v", err, runErr)
	}
}

func TestWrapRun(t *testing.T) {
	run := 0
	Wrap(func() error {
		run++
		return nil
	}).Run()
	if run != 1 {
		t.Fatalf("expected %v to be %v", run, 1)
	}
}
