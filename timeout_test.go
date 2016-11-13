package taker

import (
	"errors"
	"testing"
	"time"
)

func TestTimeoutErrorError(t *testing.T) {
	limit := 0 * time.Second
	err := &TimeoutError{limit}
	if err.Limit != limit {
		t.Fatalf("expected %v to be %v", err.Limit, limit)
	}
	// In newer versions of Go, 1 second will be "1s", in older versions 1 second
	// will be "1".
	wantMsg := "exceeded " + limit.String() + " timeout"
	if gotMsg := err.Error(); gotMsg != wantMsg {
		t.Fatalf("expected %v to be %v", gotMsg, wantMsg)
	}
}

func TestTimeoutSuccess(t *testing.T) {
	task := NewLockedTestTask(nil, 0)

	go func() {
		time.Sleep(time.Millisecond * 10)
		task.Mutex.Unlock()
	}()

	if err := Timeout(task, time.Millisecond*500); err != nil {
		t.Fatalf("unexpected err", err)
	}
}

func TestTimeoutExceeded(t *testing.T) {
	task := NewLockedTestTask(nil, 0)

	go func() {
		time.Sleep(time.Millisecond * 100)
		task.Mutex.Unlock()
	}()

	if err := Timeout(task, time.Millisecond*50); err == nil {
		t.Fatalf("expected task to timeout")
	}
}

func TestTimeoutError(t *testing.T) {
	task := NewLockedTestTask(errors.New("task err"), 0)

	go func() {
		time.Sleep(time.Millisecond * 10)
		task.Mutex.Unlock()
	}()

	if err := Timeout(task, time.Millisecond*500); err != task.Err {
		t.Fatalf("expected %v to be %v", err, task.Err)
	}
}
