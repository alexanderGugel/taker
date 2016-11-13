package taker

import "sync"

type TestTask struct {
  Err   error
  ID    int
  Done  bool
  Mutex sync.Mutex
}

func (t *TestTask) Run() error {
  t.Mutex.Lock()
  defer t.Mutex.Unlock()
  t.Done = true
  return t.Err
}

func NewTestTask(err error, id int) *TestTask {
  task := &TestTask{
    Err:   err,
    ID:    0,
    Done:  false,
    Mutex: sync.Mutex{},
  }
  return task
}

func NewLockedTestTask(err error, id int) *TestTask {
  task := NewTestTask(err, id)
  task.Mutex.Lock()
  return task
}
