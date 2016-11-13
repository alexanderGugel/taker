package taker

import "testing"

func TestRetrySuccess(t *testing.T) {
	task := NewTestRetryTask(2)
	err := Retry(task, 3)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
}

func TestRetryError(t *testing.T) {
	task := NewTestRetryTask(3)
	err := Retry(task, 3)
	if err == nil {
		t.Fatal("expected err")
	}
}
