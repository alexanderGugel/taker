package taker

import (
  "errors"
  "github.com/hashicorp/go-multierror"
  "reflect"
  "testing"
  "time"
)

func TestParallelSequentialSuccess(t *testing.T) {
  task0 := NewTestTask(nil, 0)
  task1 := NewTestTask(nil, 1)
  task2 := NewTestTask(nil, 2)
  tasks := []Task{task0, task1, task2}
  testTasks := []*TestTask{task0, task1, task2}

  if err := Parallel(tasks...); err != nil {
    t.Fatalf("unexpected err: %v", err)
  }

  for _, task := range testTasks {
    if !task.Done {
      t.Fatalf("expected %v to be done", task)
    }
  }
}

func TestParallelConcurrentSuccess(t *testing.T) {
  task0 := NewLockedTestTask(nil, 0)
  task1 := NewLockedTestTask(nil, 1)
  task2 := NewLockedTestTask(nil, 2)
  tasks := []Task{task0, task1, task2}
  testTasks := []*TestTask{task0, task1, task2}

  go func() {
    time.Sleep(time.Millisecond * 10)
    task1.Mutex.Unlock()
    time.Sleep(time.Millisecond * 10)
    task0.Mutex.Unlock()
    time.Sleep(time.Millisecond * 10)
    task2.Mutex.Unlock()
  }()

  if err := Parallel(tasks...); err != nil {
    t.Fatalf("unexpected err: %v", err)
  }

  for _, task := range testTasks {
    if !task.Done {
      t.Fatalf("expected %v to be done", task)
    }
  }
}

func TestParallelConcurrentError(t *testing.T) {
  task0 := NewLockedTestTask(nil, 0)
  task1 := NewLockedTestTask(errors.New("task1"), 1)
  task2 := NewLockedTestTask(nil, 2)
  tasks := []Task{task0, task1, task2}
  testTasks := []*TestTask{task0, task1, task2}

  go func() {
    time.Sleep(time.Millisecond * 10)
    task1.Mutex.Unlock()
    time.Sleep(time.Millisecond * 10)
    task0.Mutex.Unlock()
    time.Sleep(time.Millisecond * 10)
    task2.Mutex.Unlock()
  }()

  expectedErrs := []error{task1.Err}

  err := Parallel(tasks...)
  errs := err.(*multierror.Error).Errors
  if !reflect.DeepEqual(errs, expectedErrs) {
    t.Fatalf("expected %v to be %v", errs, expectedErrs)
  }

  for _, task := range testTasks {
    if !task.Done {
      t.Fatalf("expected %v to be done", task)
    }
  }
}

func TestParallelConcurrentErrors(t *testing.T) {
  task0 := NewLockedTestTask(errors.New("task0"), 0)
  task1 := NewLockedTestTask(errors.New("task1"), 1)
  task2 := NewLockedTestTask(nil, 2)
  tasks := []Task{task0, task1, task2}
  testTasks := []*TestTask{task0, task1, task2}

  go func() {
    time.Sleep(time.Millisecond * 10)
    task1.Mutex.Unlock()
    time.Sleep(time.Millisecond * 10)
    task0.Mutex.Unlock()
    time.Sleep(time.Millisecond * 10)
    task2.Mutex.Unlock()
  }()

  // Unlock task1 before task0.
  expectedErrs := []error{task1.Err, task0.Err}

  err := Parallel(tasks...)
  errs := err.(*multierror.Error).Errors
  if !reflect.DeepEqual(errs, expectedErrs) {
    t.Fatalf("expected %v to be %v", errs, expectedErrs)
  }

  for _, task := range testTasks {
    if !task.Done {
      t.Fatalf("expected %v to be done", task)
    }
  }
}
