package taker

import (
	"errors"
	"github.com/hashicorp/go-multierror"
	"reflect"
	"testing"
)

func TestForceSeriesSuccess(t *testing.T) {
	tasks := []Task{
		NewTestTask(nil, 0),
		NewTestTask(nil, 1),
		NewTestTask(nil, 2),
	}

	if err := ForceSeries(tasks...); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	for _, task := range tasks {
		if !task.(*TestTask).Done {
			t.Fatalf("expected %v to be done", task)
		}
	}
}

func TestForceSeriesError(t *testing.T) {
	task0 := NewTestTask(nil, 0)
	task1 := NewTestTask(errors.New("task1 err"), 1)
	task2 := NewTestTask(errors.New("task2 err"), 2)

	tasks := []Task{task0, task1, task2}
	expectedErrs := []error{task1.Err, task2.Err}

	err := ForceSeries(tasks...)

	errs := err.(*multierror.Error).Errors
	if !reflect.DeepEqual(errs, expectedErrs) {
		t.Fatalf("expected %v to be %v", errs, expectedErrs)
	}

	if !task0.Done {
		t.Fatalf("expected %v to be done", task0)
	}

	if !task1.Done {
		t.Fatalf("expected %v to be done", task1)
	}

	if !task2.Done {
		t.Fatalf("expected %v to be done", task2)
	}
}
