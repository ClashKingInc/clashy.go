package events_test

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	clashy "clashy.go"
	"clashy.go/events"
	"clashy.go/stores"
)

func TestTrackerDispatchesOnChange(t *testing.T) {
	store := stores.NewStore()
	var calls atomic.Int32
	var value atomic.Int32

	tracker := events.NewTracker("player", nil, func(_ context.Context, tag string) (*clashy.Player, int, error) {
		return &clashy.Player{Tag: tag, Trophies: int(value.Load())}, 0, nil
	}, store)

	tracker.Group("fast", events.GroupTags("#2PP"), events.Interval(10*time.Millisecond))
	tracker.On(events.FieldChanged(
		"player_trophies_changed",
		func(player clashy.Player) int {
			return player.Trophies
		},
		func(_ context.Context, change events.FieldChange[clashy.Player, int]) error {
			calls.Add(1)
			return nil
		},
	))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := tracker.Start(ctx); err != nil {
		t.Fatalf("start tracker: %v", err)
	}

	time.Sleep(20 * time.Millisecond)
	value.Store(10)
	time.Sleep(30 * time.Millisecond)

	if calls.Load() == 0 {
		t.Fatalf("expected at least one change callback")
	}
}

func TestStaticTagSourceNormalizesTags(t *testing.T) {
	source := events.Tags("2pp", " o0o ")
	var tags []string
	if err := source.ForEach(context.Background(), func(tag string) error {
		tags = append(tags, tag)
		return nil
	}); err != nil {
		t.Fatalf("foreach: %v", err)
	}
	if len(tags) != 2 || tags[0] != "#2PP" || tags[1] != "#000" {
		t.Fatalf("unexpected normalized tags: %#v", tags)
	}
}

func TestTrackerPollOnlyDispatchesWithoutStore(t *testing.T) {
	var calls atomic.Int32

	tracker := events.NewTracker("player", nil, func(_ context.Context, tag string) (*clashy.Player, int, error) {
		return &clashy.Player{Tag: tag, Trophies: 1234}, 0, nil
	}, nil)

	tracker.Group("poll-only",
		events.GroupTags("#2PP"),
		events.Interval(10*time.Millisecond),
		events.PollOnly(),
	)

	tracker.On(events.EveryPoll("player_poll", func(_ context.Context, change events.Change[clashy.Player]) error {
		calls.Add(1)
		if change.Previous != nil {
			t.Fatalf("expected nil previous value in poll-only mode")
		}
		return nil
	}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := tracker.Start(ctx); err != nil {
		t.Fatalf("start tracker: %v", err)
	}

	time.Sleep(30 * time.Millisecond)

	if calls.Load() == 0 {
		t.Fatalf("expected poll-only callbacks to fire")
	}
}

func TestTrackerDropTagAndDeleteSnapshotOnNotFound(t *testing.T) {
	store := stores.NewStore()
	var calls atomic.Int32
	var fetches atomic.Int32

	tracker := events.NewTracker("player", nil, func(_ context.Context, tag string) (*clashy.Player, int, error) {
		fetches.Add(1)
		if fetches.Load() == 1 {
			return &clashy.Player{Tag: tag, Trophies: 100}, 0, nil
		}
		return nil, 0, &clashy.NotFound{HTTPException: &clashy.HTTPException{Status: 404, Reason: "not found"}}
	}, store)

	tracker.Group("tracked",
		events.GroupTags("#2PP"),
		events.Interval(10*time.Millisecond),
		events.OnError(func(ctx context.Context, ec events.ErrorContext) events.ErrorDecision {
			var notFound *clashy.NotFound
			if errors.As(ec.Err, &notFound) {
				return events.ErrorDecision{
					Action: events.ErrorActionDropTag | events.ErrorActionDeleteSnapshot,
				}
			}
			return events.ErrorDecision{Action: events.ErrorActionSkip}
		}),
	)

	tracker.On(events.EveryPoll("player_polled", func(_ context.Context, change events.Change[clashy.Player]) error {
		calls.Add(1)
		return nil
	}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := tracker.Start(ctx); err != nil {
		t.Fatalf("start tracker: %v", err)
	}

	time.Sleep(60 * time.Millisecond)

	if calls.Load() != 1 {
		t.Fatalf("expected exactly one successful poll before tag removal, got %d", calls.Load())
	}
	if fetches.Load() > 3 {
		t.Fatalf("expected tag to be removed after not found, got %d fetches", fetches.Load())
	}
	if _, err := store.Load(context.Background(), "player", "#2PP"); !errors.Is(err, events.ErrSnapshotNotFound) {
		t.Fatalf("expected snapshot deletion after not found, got %v", err)
	}
}

func TestTrackerOnErrorGetsPreviousSnapshot(t *testing.T) {
	store := stores.NewStore()
	var previousTrophies atomic.Int32
	var fetches atomic.Int32

	tracker := events.NewTracker("player", nil, func(_ context.Context, tag string) (*clashy.Player, int, error) {
		fetches.Add(1)
		if fetches.Load() == 1 {
			return &clashy.Player{Tag: tag, Trophies: 321}, 0, nil
		}
		return nil, 0, &clashy.NotFound{HTTPException: &clashy.HTTPException{Status: 404, Reason: "not found"}}
	}, store)

	tracker.Group(
		"tracked",
		events.GroupTags("#2PP"),
		events.Interval(10*time.Millisecond),
	)

	tracker.On(events.EveryPoll("player_polled", func(_ context.Context, change events.Change[clashy.Player]) error {
		return nil
	}))

	tracker.OnError(func(ctx context.Context, change events.ErrorChange[clashy.Player]) *events.ErrorDecision {
		var notFound *clashy.NotFound
		if !errors.As(change.Err, &notFound) {
			return nil
		}
		if change.Previous == nil {
			t.Fatalf("expected previous snapshot on not found")
		}
		previousTrophies.Store(int32(change.Previous.Trophies))
		return &events.ErrorDecision{
			Action: events.ErrorActionDropTag | events.ErrorActionDeleteSnapshot,
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := tracker.Start(ctx); err != nil {
		t.Fatalf("start tracker: %v", err)
	}

	time.Sleep(60 * time.Millisecond)

	if previousTrophies.Load() != 321 {
		t.Fatalf("expected previous trophies to be passed to OnError, got %d", previousTrophies.Load())
	}
	if _, err := store.Load(context.Background(), "player", "#2PP"); !errors.Is(err, events.ErrSnapshotNotFound) {
		t.Fatalf("expected snapshot deletion after not found, got %v", err)
	}
}

func TestTrackerOnErrorReceivesPreviousSnapshot(t *testing.T) {
	store := stores.NewStore()
	var previousName atomic.Value
	var stages atomic.Int32
	var fetches atomic.Int32

	tracker := events.NewTracker("player", nil, func(_ context.Context, tag string) (*clashy.Player, int, error) {
		fetches.Add(1)
		if fetches.Load() == 1 {
			return &clashy.Player{Tag: tag, Name: "before"}, 0, nil
		}
		return nil, 0, &clashy.NotFound{HTTPException: &clashy.HTTPException{Status: 404, Reason: "not found"}}
	}, store)

	tracker.Group("tracked",
		events.GroupTags("#2PP"),
		events.Interval(10*time.Millisecond),
	)
	tracker.OnError(func(ctx context.Context, ec events.ErrorChange[clashy.Player]) *events.ErrorDecision {
		var notFound *clashy.NotFound
		if !errors.As(ec.Err, &notFound) {
			return nil
		}
		stages.Add(1)
		if ec.Stage != events.ErrorStageFetch {
			t.Fatalf("expected fetch stage, got %s", ec.Stage)
		}
		if ec.Previous == nil {
			t.Fatalf("expected previous snapshot to be provided")
		}
		previousName.Store(ec.Previous.Name)
		return &events.ErrorDecision{
			Action: events.ErrorActionDropTag | events.ErrorActionDeleteSnapshot,
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := tracker.Start(ctx); err != nil {
		t.Fatalf("start tracker: %v", err)
	}

	time.Sleep(60 * time.Millisecond)

	if stages.Load() != 1 {
		t.Fatalf("expected one handled error, got %d", stages.Load())
	}
	if name, _ := previousName.Load().(string); name != "before" {
		t.Fatalf("expected previous snapshot name to be preserved, got %q", name)
	}
	if _, err := store.Load(context.Background(), "player", "#2PP"); !errors.Is(err, events.ErrSnapshotNotFound) {
		t.Fatalf("expected snapshot deletion after tracker error handler, got %v", err)
	}
}

func TestTrackerOnErrorCanDeferToDefaultMaintenancePause(t *testing.T) {
	var fetches atomic.Int32

	tracker := events.NewTracker("player", nil, func(_ context.Context, tag string) (*clashy.Player, int, error) {
		fetches.Add(1)
		return nil, 0, &clashy.Maintenance{HTTPException: &clashy.HTTPException{Status: 503, Reason: "maintenance"}}
	}, stores.NewStore())

	tracker.Group("tracked",
		events.GroupTags("#2PP"),
		events.Interval(5*time.Millisecond),
	)
	tracker.OnError(func(ctx context.Context, ec events.ErrorChange[clashy.Player]) *events.ErrorDecision {
		return nil
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := tracker.Start(ctx); err != nil {
		t.Fatalf("start tracker: %v", err)
	}

	time.Sleep(40 * time.Millisecond)

	if fetches.Load() > 1 {
		t.Fatalf("expected maintenance to pause the group via default fallback, got %d fetches", fetches.Load())
	}
}
