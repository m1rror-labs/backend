package multisync_test

import (
	"mirror-backend/pkg/dependencies/multisync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	// This is a simple test to ensure that the mutex allows the correct number of threads
	// and that it releases them properly.

	m := multisync.NewMutex(3)

	// Acquire 3 threads
	m.Acquire()
	m.Acquire()
	m.Acquire()

	// Ensure we can acquire a fourth thread
	ch4 := m.Acquire()

	// Release the first three threads
	m.Release()
	m.Release()
	m.Release()

	// Now we should be able to acquire the fourth thread
	select {
	case <-ch4:
		// Successfully acquired the fourth thread
	case <-time.After(time.Second):
		t.Fatal("failed to acquire the fourth thread")
	}
}
