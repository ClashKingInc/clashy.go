package clashy

import (
	"context"
)

type rateLimitContextKey struct{}

type requestLimiter struct {
	slots chan struct{}
}

func newRequestLimiter(limit int) *requestLimiter {
	if limit <= 0 {
		return nil
	}
	return &requestLimiter{slots: make(chan struct{}, limit)}
}

func (l *requestLimiter) Acquire(ctx context.Context) (func(), error) {
	if l == nil {
		return func() {}, nil
	}
	ctx = contextOrBackground(ctx)

	select {
	case l.slots <- struct{}{}:
		return func() {
			<-l.slots
		}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func WithoutRateLimit(ctx context.Context) context.Context {
	return context.WithValue(contextOrBackground(ctx), rateLimitContextKey{}, true)
}

func rateLimitDisabled(ctx context.Context) bool {
	ctx = contextOrBackground(ctx)
	disabled, _ := ctx.Value(rateLimitContextKey{}).(bool)
	return disabled
}
