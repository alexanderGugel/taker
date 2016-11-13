package taker

import (
	"errors"
	"testing"
)

func TestWhilstError(t *testing.T) {
	test := func() bool {
		return true
	}
	task := NewTestTask(errors.New("task err"), 0)
	err := Whilst(test, task)
	if err == nil {
		t.Fatal("expected err")
	}
}

func TestWhilstFalseTest(t *testing.T) {
	test := func() bool {
		return false
	}
	task := NewTestTask(errors.New("task err"), 0)
	err := Whilst(test, task)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if task.Done {
		t.Fatal("expected task not to be run")
	}
}

func TestWhilstEventualTrueTest(t *testing.T) {
	counter := 0
	test := func() bool {
		counter++
		return counter != 3
	}
	task := NewTestTask(nil, 0)
	err := Whilst(test, task)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if !task.Done {
		t.Fatal("expected task to be run")
	}
	if counter != 3 {
		t.Fatalf("expected %v to be %v", counter, 3)
	}
}
