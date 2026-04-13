package clashy_test

import (
	"sync/atomic"
	"testing"
	"time"
)

func waitForCount(t *testing.T, counter *atomic.Int32, want int32, timeout time.Duration) {
	t.Helper()

	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if counter.Load() >= want {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatalf("timed out waiting for count %d, got %d", want, counter.Load())
}

func setMax(dst *atomic.Int32, current int32) {
	for {
		currentMax := dst.Load()
		if current <= currentMax {
			return
		}
		if dst.CompareAndSwap(currentMax, current) {
			return
		}
	}
}
