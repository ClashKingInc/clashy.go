package clashy

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

// ErrInvalidLimit is returned when a limiter is created or used without a
// positive requests-per-second limit.
var ErrInvalidLimit = errors.New("clashy: requests per second must be greater than zero")

// Limiter gates request starts with a strict rolling one-second window and
// also caps concurrent in-flight work.
type Limiter struct {
	rps         int
	maxInFlight int

	mu       sync.Mutex
	starts   []time.Time
	head     int
	count    int
	inFlight int
	wake     chan struct{}
}

// NewLimiter creates a limiter. If maxInFlight is zero or negative, rps is
// used as the in-flight limit.
func NewLimiter(rps, maxInFlight int) (*Limiter, error) {
	if rps <= 0 {
		return nil, ErrInvalidLimit
	}
	if maxInFlight <= 0 {
		maxInFlight = rps
	}
	return &Limiter{
		rps:         rps,
		maxInFlight: maxInFlight,
		starts:      make([]time.Time, rps),
		wake:        make(chan struct{}, maxInFlight),
	}, nil
}

// Acquire waits until starting one more operation stays within both the RPS
// and in-flight limits. The returned release function must be called once the
// operation finishes.
func (l *Limiter) Acquire(ctx context.Context) (func(), error) {
	if l == nil || l.rps <= 0 {
		return nil, ErrInvalidLimit
	}
	if ctx == nil {
		ctx = context.Background()
	}

	for {
		l.mu.Lock()
		now := time.Now()
		l.trimStarts(now)
		if l.inFlight < l.maxInFlight && l.count < l.rps {
			l.recordStart(now)
			l.inFlight++
			l.mu.Unlock()

			var released atomic.Bool
			return func() {
				if released.CompareAndSwap(false, true) {
					l.release()
				}
			}, nil
		}

		if l.count == l.rps {
			wait := l.starts[l.head].Add(time.Second).Sub(now)
			if wait < 0 {
				wait = 0
			}
			l.mu.Unlock()
			if wait == 0 {
				continue
			}
			timer := time.NewTimer(wait)
			select {
			case <-ctx.Done():
				stopTimer(timer)
				return nil, ctx.Err()
			case <-timer.C:
			}
			continue
		}
		l.mu.Unlock()

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-l.wake:
		}
	}
}

func (l *Limiter) recordStart(now time.Time) {
	index := (l.head + l.count) % len(l.starts)
	l.starts[index] = now
	l.count++
}

func (l *Limiter) release() {
	l.mu.Lock()
	if l.inFlight > 0 {
		l.inFlight--
	}
	l.mu.Unlock()

	select {
	case l.wake <- struct{}{}:
	default:
	}
}

func (l *Limiter) trimStarts(now time.Time) {
	cutoff := now.Add(-time.Second)
	for l.count > 0 && !l.starts[l.head].After(cutoff) {
		l.starts[l.head] = time.Time{}
		l.head = (l.head + 1) % len(l.starts)
		l.count--
	}
}

func stopTimer(timer *time.Timer) {
	if !timer.Stop() {
		select {
		case <-timer.C:
		default:
		}
	}
}
