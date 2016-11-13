package taker

import (
	"testing"
	"time"
)

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
