package events

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	clashy "github.com/clashkinginc/clashy.go"
)

var ErrSnapshotNotFound = errors.New("snapshot not found")

type Snapshot struct {
	Kind      string
	Key       string
	Data      []byte
	UpdatedAt time.Time
}

type SnapshotStore interface {
	Load(ctx context.Context, kind, key string) (Snapshot, error)
	Save(ctx context.Context, snapshot Snapshot) error
	Delete(ctx context.Context, kind, key string) error
}

type TagSource interface {
	ForEach(ctx context.Context, fn func(tag string) error) error
}

type StaticTagSource struct {
	mu   sync.Mutex
	tags []string
}

func Tags(tags ...string) TagSource {
	normalized := make([]string, 0, len(tags))
	for _, tag := range tags {
		normalized = append(normalized, clashy.CorrectTag(tag))
	}
	return &StaticTagSource{tags: normalized}
}

func (s *StaticTagSource) ForEach(_ context.Context, fn func(tag string) error) error {
	s.mu.Lock()
	tags := append([]string(nil), s.tags...)
	s.mu.Unlock()
	for _, tag := range tags {
		if err := fn(tag); err != nil {
			return err
		}
	}
	return nil
}

func (s *StaticTagSource) RemoveTag(tag string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	tag = clashy.CorrectTag(tag)
	filtered := s.tags[:0]
	for _, existing := range s.tags {
		if existing != tag {
			filtered = append(filtered, existing)
		}
	}
	s.tags = filtered
}

type FuncTagSource func(ctx context.Context, fn func(tag string) error) error

func (f FuncTagSource) ForEach(ctx context.Context, fn func(tag string) error) error {
	return f(ctx, fn)
}

type Change[T any] struct {
	Group    string
	Kind     string
	Tag      string
	Event    string
	Previous *T
	Current  *T
}

type Handler[T any] func(context.Context, Change[T]) error

type Spec[T any] interface {
	Name() string
	Evaluate(context.Context, Change[T]) error
}

type specFunc[T any] struct {
	name string
	fn   func(context.Context, Change[T]) error
}

func (s specFunc[T]) Name() string { return s.name }
func (s specFunc[T]) Evaluate(ctx context.Context, change Change[T]) error {
	change.Event = s.name
	return s.fn(ctx, change)
}

func EveryPoll[T any](name string, handler Handler[T]) Spec[T] {
	return specFunc[T]{
		name: name,
		fn: func(ctx context.Context, change Change[T]) error {
			if change.Current == nil {
				return nil
			}
			return handler(ctx, change)
		},
	}
}

func Custom[T any](name string, match func(oldValue, newValue *T) bool, handler Handler[T]) Spec[T] {
	return specFunc[T]{
		name: name,
		fn: func(ctx context.Context, change Change[T]) error {
			if change.Previous == nil || change.Current == nil {
				return nil
			}
			if match(change.Previous, change.Current) {
				return handler(ctx, change)
			}
			return nil
		},
	}
}

type FieldChange[T any, V comparable] struct {
	Change[T]
	OldValue V
	NewValue V
}

func FieldChanged[T any, V comparable](name string, selector func(T) V, handler func(context.Context, FieldChange[T, V]) error) Spec[T] {
	return specFunc[T]{
		name: name,
		fn: func(ctx context.Context, change Change[T]) error {
			if change.Previous == nil || change.Current == nil {
				return nil
			}
			oldValue := selector(*change.Previous)
			newValue := selector(*change.Current)
			if oldValue == newValue {
				return nil
			}
			return handler(ctx, FieldChange[T, V]{
				Change:   change,
				OldValue: oldValue,
				NewValue: newValue,
			})
		},
	}
}

type ErrorAction uint8

const (
	ErrorActionSkip           ErrorAction = 0
	ErrorActionDeleteSnapshot ErrorAction = 1 << iota
	ErrorActionDropTag
	ErrorActionPauseGroup
	ErrorActionStopTracker
)

type ErrorContext struct {
	Group string
	Kind  string
	Tag   string
	Err   error
	Store SnapshotStore
}

type ErrorStage string

const (
	ErrorStageFetch   ErrorStage = "fetch"
	ErrorStageMarshal ErrorStage = "marshal"
	ErrorStageStore   ErrorStage = "store"
)

type ErrorChange[T any] struct {
	Group    string
	Kind     string
	Tag      string
	Stage    ErrorStage
	Err      error
	Store    SnapshotStore
	Previous *T
	Current  *T
}

type ErrorDecision struct {
	Action   ErrorAction
	PauseFor time.Duration
}

type ErrorHandler func(context.Context, ErrorContext) ErrorDecision
type TrackerErrorHandler[T any] func(context.Context, ErrorChange[T]) *ErrorDecision

func defaultErrorHandler(ctx context.Context, ec ErrorContext) ErrorDecision {
	_ = ctx
	var maintenance *clashy.Maintenance
	if errors.As(ec.Err, &maintenance) {
		return ErrorDecision{Action: ErrorActionPauseGroup, PauseFor: 30 * time.Second}
	}
	return ErrorDecision{Action: ErrorActionSkip}
}

type GroupConfig struct {
	Name             string
	Interval         time.Duration
	RateLimit        int
	Store            SnapshotStore
	TagSource        TagSource
	CompareSnapshots bool
	ErrorHandler     ErrorHandler
}

func defaultGroupConfig(name string) GroupConfig {
	return GroupConfig{
		Name:             name,
		Interval:         30 * time.Second,
		TagSource:        Tags(),
		CompareSnapshots: true,
		ErrorHandler:     defaultErrorHandler,
	}
}

type GroupOption func(*GroupConfig)

func Interval(interval time.Duration) GroupOption {
	return func(cfg *GroupConfig) { cfg.Interval = interval }
}

func RateLimit(limit int) GroupOption {
	return func(cfg *GroupConfig) {
		if limit >= 0 {
			cfg.RateLimit = limit
		}
	}
}

func Store(store SnapshotStore) GroupOption {
	return func(cfg *GroupConfig) { cfg.Store = store }
}

func Source(source TagSource) GroupOption {
	return func(cfg *GroupConfig) { cfg.TagSource = source }
}

func GroupTags(tags ...string) GroupOption {
	return func(cfg *GroupConfig) { cfg.TagSource = Tags(tags...) }
}

func PollOnly() GroupOption {
	return func(cfg *GroupConfig) { cfg.CompareSnapshots = false }
}

func OnError(handler ErrorHandler) GroupOption {
	return func(cfg *GroupConfig) {
		if handler != nil {
			cfg.ErrorHandler = handler
		}
	}
}

type FetchFunc[T any] func(context.Context, string) (*T, int, error)

type Tracker[T any] struct {
	kind         string
	client       *clashy.Client
	fetch        FetchFunc[T]
	defaultStore SnapshotStore

	mu       sync.RWMutex
	groups   map[string]GroupConfig
	specs    []Spec[T]
	onError  TrackerErrorHandler[T]
	leasesMu sync.Mutex
	leases   map[string]*sync.Mutex

	stopOnce sync.Once
	stopCh   chan struct{}
}

func newTracker[T any](kind string, client *clashy.Client, fetch FetchFunc[T], store SnapshotStore) *Tracker[T] {
	return &Tracker[T]{
		kind:         kind,
		client:       client,
		fetch:        fetch,
		defaultStore: store,
		groups:       make(map[string]GroupConfig),
		leases:       make(map[string]*sync.Mutex),
		stopCh:       make(chan struct{}),
	}
}

func (t *Tracker[T]) Group(name string, options ...GroupOption) *Tracker[T] {
	t.mu.Lock()
	defer t.mu.Unlock()
	cfg := defaultGroupConfig(name)
	for _, option := range options {
		option(&cfg)
	}
	if cfg.Store == nil {
		cfg.Store = t.defaultStore
	}
	if cfg.TagSource == nil {
		cfg.TagSource = Tags()
	}
	t.groups[name] = cfg
	return t
}

func (t *Tracker[T]) On(spec Spec[T]) *Tracker[T] {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.specs = append(t.specs, spec)
	return t
}

func (t *Tracker[T]) OnMany(specs ...Spec[T]) *Tracker[T] {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.specs = append(t.specs, specs...)
	return t
}

func (t *Tracker[T]) OnError(handler TrackerErrorHandler[T]) *Tracker[T] {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.onError = handler
	return t
}

func (t *Tracker[T]) Start(ctx context.Context) error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, group := range t.groups {
		group := group
		go t.runGroup(ctx, group)
	}
	return nil
}

func (t *Tracker[T]) runGroup(ctx context.Context, group GroupConfig) {
	limiter := newGroupLimiter(group.RateLimit)
	for {
		delay, stop := t.runGroupOnce(ctx, group, limiter)
		if stop {
			return
		}
		if delay <= 0 {
			delay = group.Interval
		}
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-t.stopCh:
			timer.Stop()
			return
		case <-timer.C:
		}
	}
}

func (t *Tracker[T]) runGroupOnce(ctx context.Context, group GroupConfig, limiter *groupLimiter) (time.Duration, bool) {
	if group.TagSource == nil || (group.CompareSnapshots && group.Store == nil) {
		return group.Interval, false
	}
	var wg sync.WaitGroup
	var pauseFor time.Duration
	var shouldStop bool
	var actionMu sync.Mutex
	_ = group.TagSource.ForEach(ctx, func(tag string) error {
		tag = clashy.CorrectTag(tag)
		wg.Add(1)
		go func(tag string) {
			defer wg.Done()
			decision := t.processTag(ctx, group, tag, limiter)
			actionMu.Lock()
			defer actionMu.Unlock()
			if decision.PauseFor > pauseFor {
				pauseFor = decision.PauseFor
			}
			if decision.Action&ErrorActionStopTracker != 0 {
				shouldStop = true
				t.stop()
			}
		}(tag)
		return nil
	})
	wg.Wait()
	if pauseFor > 0 {
		return pauseFor, shouldStop
	}
	return group.Interval, shouldStop
}

func (t *Tracker[T]) processTag(ctx context.Context, group GroupConfig, tag string, limiter *groupLimiter) ErrorDecision {
	lock := t.tagLease(tag)
	lock.Lock()
	defer lock.Unlock()

	if limiter != nil {
		release, err := limiter.Acquire(ctx)
		if err != nil {
			return ErrorDecision{Action: ErrorActionSkip}
		}
		defer release()
		ctx = clashy.WithoutRateLimit(ctx)
	}

	current, retry, err := t.fetch(ctx, tag)
	if err != nil || current == nil {
		if err == nil {
			err = errors.New("nil current value returned without error")
		}
		decision := t.handleError(ctx, group, tag, ErrorStageFetch, err, nil, current)
		t.applyErrorDecision(ctx, group, tag, decision)
		return decision
	}
	currentBytes, err := json.Marshal(current)
	if err != nil {
		decision := t.handleError(ctx, group, tag, ErrorStageMarshal, err, nil, current)
		t.applyErrorDecision(ctx, group, tag, decision)
		return decision
	}

	if !group.CompareSnapshots {
		change := Change[T]{
			Group:    group.Name,
			Kind:     t.kind,
			Tag:      tag,
			Previous: nil,
			Current:  current,
		}
		t.mu.RLock()
		specs := append([]Spec[T](nil), t.specs...)
		t.mu.RUnlock()
		for _, spec := range specs {
			_ = spec.Evaluate(ctx, change)
		}
		return ErrorDecision{Action: ErrorActionSkip}
	}

	var previous *T
	if snapshot, err := group.Store.Load(ctx, t.kind, tag); err == nil {
		var oldValue T
		if json.Unmarshal(snapshot.Data, &oldValue) == nil {
			previous = &oldValue
		}
	}

	change := Change[T]{
		Group:    group.Name,
		Kind:     t.kind,
		Tag:      tag,
		Previous: previous,
		Current:  current,
	}

	t.mu.RLock()
	specs := append([]Spec[T](nil), t.specs...)
	t.mu.RUnlock()
	for _, spec := range specs {
		_ = spec.Evaluate(ctx, change)
	}

	updatedAt := time.Now()
	if retry > 0 {
		updatedAt = updatedAt.Add(time.Duration(retry) * time.Second)
	}
	if err := group.Store.Save(ctx, Snapshot{
		Kind:      t.kind,
		Key:       tag,
		Data:      currentBytes,
		UpdatedAt: updatedAt,
	}); err != nil {
		decision := t.handleError(ctx, group, tag, ErrorStageStore, err, previous, current)
		t.applyErrorDecision(ctx, group, tag, decision)
		return decision
	}
	return ErrorDecision{Action: ErrorActionSkip}
}

func (t *Tracker[T]) tagLease(tag string) *sync.Mutex {
	t.leasesMu.Lock()
	defer t.leasesMu.Unlock()
	if lease, ok := t.leases[tag]; ok {
		return lease
	}
	lease := &sync.Mutex{}
	t.leases[tag] = lease
	return lease
}

func (t *Tracker[T]) stop() {
	t.stopOnce.Do(func() {
		close(t.stopCh)
	})
}

type removableTagSource interface {
	RemoveTag(tag string)
}

func (t *Tracker[T]) applyErrorDecision(ctx context.Context, group GroupConfig, tag string, decision ErrorDecision) {
	if decision.Action&ErrorActionDeleteSnapshot != 0 && group.Store != nil {
		_ = group.Store.Delete(ctx, t.kind, tag)
	}
	if decision.Action&ErrorActionDropTag != 0 {
		if source, ok := group.TagSource.(removableTagSource); ok {
			source.RemoveTag(tag)
		}
	}
}

func (t *Tracker[T]) handleError(ctx context.Context, group GroupConfig, tag string, stage ErrorStage, err error, previous, current *T) ErrorDecision {
	if previous == nil && group.CompareSnapshots && group.Store != nil {
		if snapshot, loadErr := group.Store.Load(ctx, t.kind, tag); loadErr == nil {
			var oldValue T
			if json.Unmarshal(snapshot.Data, &oldValue) == nil {
				previous = &oldValue
			}
		}
	}

	t.mu.RLock()
	onError := t.onError
	t.mu.RUnlock()
	if onError != nil {
		if decision := onError(ctx, ErrorChange[T]{
			Group:    group.Name,
			Kind:     t.kind,
			Tag:      tag,
			Stage:    stage,
			Err:      err,
			Store:    group.Store,
			Previous: previous,
			Current:  current,
		}); decision != nil {
			return *decision
		}
	}

	return group.ErrorHandler(ctx, ErrorContext{
		Group: group.Name,
		Kind:  t.kind,
		Tag:   tag,
		Err:   err,
		Store: group.Store,
	})
}

type groupLimiter struct {
	slots chan struct{}
}

func newGroupLimiter(limit int) *groupLimiter {
	if limit <= 0 {
		return nil
	}
	return &groupLimiter{slots: make(chan struct{}, limit)}
}

func (l *groupLimiter) Acquire(ctx context.Context) (func(), error) {
	if l == nil {
		return func() {}, nil
	}
	if ctx == nil {
		ctx = context.Background()
	}

	select {
	case l.slots <- struct{}{}:
		return func() {
			<-l.slots
		}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
