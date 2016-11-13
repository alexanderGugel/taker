package taker

import (
	"errors"
	"testing"
)

func TestSeriesSuccess(t *testing.T) {
	tasks := []Task{
		NewTestTask(nil, 0),
		NewTestTask(nil, 1),
		NewTestTask(nil, 2),
	}

	if err := Series(tasks...); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	for _, task := range tasks {
		if !task.(*TestTask).Done {
			t.Fatalf("expected %v to be done", task)
		}
	}
}

func TestSeriesError(t *testing.T) {
	task1Err := errors.New("task1 err")

	task0 := NewTestTask(nil, 0)
	task1 := NewTestTask(task1Err, 1)
	task2 := NewTestTask(nil, 2)
	tasks := []Task{task0, task1, task2}

	if err := Series(tasks...); err != task1Err {
		t.Fatal("expected %v to be %v", err, task1Err)
	}

	if !task0.Done {
		t.Fatalf("expected %v to be done", task0)
	}

	if !task1.Done {
		t.Fatalf("expected %v to be done", task1)
	}

	if task2.Done {
		t.Fatalf("expected %v not to be done", task2)
	}
}
