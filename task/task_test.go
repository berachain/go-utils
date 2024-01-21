package task

import (
	"testing"
	"time"
)

// TestTask tests the CancellablePeriodicTask function
// by creating a task that increments a variable every
// delay_us microseconds, and then cancelling the task
// after delay_us * 10 microseconds
// The variable should be incremented 10 times
func TestTask(t *testing.T) {
	dummyFunc := func(t *int) {
		*t += 1
	}

	v := 0
	delay_us := uint64(1000)
	_, cancel := CancellablePeriodicTask(delay_us, &v, dummyFunc)

	// Sleep for delay_us * 10
	time.Sleep(time.Duration(delay_us*10) * time.Microsecond)

	// Cancel the context
	cancel()

	// Check the value
	if !(9 <= v && v <= 11) {
		t.Errorf("Expected value in range [9, 11], got %v", v)
	}
}
