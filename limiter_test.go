package clashy

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestLimiterCapsStartsAndInFlight(t *testing.T) {
	limiter, err := NewLimiter(3, 2)
	if err != nil {
		t.Fatalf("NewLimiter: %v", err)
	}

	var active atomic.Int32
	var peak atomic.Int32
	gate := make(chan struct{})
	var wg sync.WaitGroup
	for range 3 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			release, acquireErr := limiter.Acquire(context.Background())
			if acquireErr != nil {
				t.Errorf("Acquire: %v", acquireErr)
				return
			}
			current := active.Add(1)
			for current > peak.Load() && !peak.CompareAndSwap(peak.Load(), current) {
			}
			<-gate
			active.Add(-1)
			release()
		}()
	}

	deadline := time.Now().Add(250 * time.Millisecond)
	for peak.Load() < 2 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if got := peak.Load(); got != 2 {
		t.Fatalf("peak in-flight = %d, want 2", got)
	}
	close(gate)
	wg.Wait()
}

func TestLimiterUsesStrictRollingWindow(t *testing.T) {
	limiter, err := NewLimiter(2, 2)
	if err != nil {
		t.Fatalf("NewLimiter: %v", err)
	}
	for range 2 {
		release, acquireErr := limiter.Acquire(context.Background())
		if acquireErr != nil {
			t.Fatalf("Acquire: %v", acquireErr)
		}
		release()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if release, acquireErr := limiter.Acquire(ctx); !errors.Is(acquireErr, context.DeadlineExceeded) {
		if acquireErr == nil {
			release()
		}
		t.Fatalf("Acquire error = %v, want deadline exceeded", acquireErr)
	}
}

func TestLimiterWakesOneWaiterPerReleasedSlot(t *testing.T) {
	limiter, err := NewLimiter(10, 3)
	if err != nil {
		t.Fatalf("NewLimiter: %v", err)
	}

	releases := make([]func(), 0, 3)
	for range 3 {
		release, acquireErr := limiter.Acquire(context.Background())
		if acquireErr != nil {
			t.Fatalf("Acquire: %v", acquireErr)
		}
		releases = append(releases, release)
	}

	started := make(chan func(), 3)
	for range 3 {
		go func() {
			release, acquireErr := limiter.Acquire(context.Background())
			if acquireErr == nil {
				started <- release
			}
		}()
	}
	for _, release := range releases {
		release()
	}

	for range 3 {
		select {
		case release := <-started:
			release()
		case <-time.After(250 * time.Millisecond):
			t.Fatal("waiter did not wake after an in-flight slot was released")
		}
	}
}
