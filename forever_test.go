package taker

import (
	"testing"
)

func TestForever(t *testing.T) {
	task := NewTestCounterTask(2)
	err := Forever(task)
	if err == nil {
		t.Fatal("expected err")
	}
	if task.Counter != 0 {
		t.Fatalf("expected %v to be %v", task.Counter, 0)
	}
}
