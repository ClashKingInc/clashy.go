package events

import (
	"context"

	clashy "github.com/clashkinginc/clashy.go"
)

type Engine struct {
	client   *clashy.Client
	store    SnapshotStore
	trackers []starter
}

type starter interface {
	Start(context.Context) error
}

func NewEngine(client *clashy.Client, store SnapshotStore) *Engine {
	return &Engine{client: client, store: store}
}

func (e *Engine) Start(ctx context.Context) error {
	for _, tracker := range e.trackers {
		if err := tracker.Start(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (e *Engine) Clans() *Tracker[clashy.Clan] {
	tracker := newTracker("clan", e.client, func(ctx context.Context, tag string) (*clashy.Clan, int, error) {
		clan, err := e.client.GetClan(ctx, tag)
		retry := 0
		if clan != nil {
			retry = clan.RetryAfter()
		}
		return clan, retry, err
	}, e.store)
	e.trackers = append(e.trackers, tracker)
	return tracker
}

func (e *Engine) Players() *Tracker[clashy.Player] {
	tracker := newTracker("player", e.client, func(ctx context.Context, tag string) (*clashy.Player, int, error) {
		player, err := e.client.GetPlayer(ctx, tag)
		retry := 0
		if player != nil {
			retry = player.RetryAfter()
		}
		return player, retry, err
	}, e.store)
	e.trackers = append(e.trackers, tracker)
	return tracker
}

func (e *Engine) Wars() *Tracker[clashy.ClanWar] {
	tracker := newTracker("war", e.client, func(ctx context.Context, tag string) (*clashy.ClanWar, int, error) {
		war, err := e.client.GetCurrentWar(ctx, tag)
		retry := 0
		if war != nil {
			retry = war.RetryAfter()
		}
		return war, retry, err
	}, e.store)
	e.trackers = append(e.trackers, tracker)
	return tracker
}

func NewTracker[T any](kind string, client *clashy.Client, fetch FetchFunc[T], store SnapshotStore) *Tracker[T] {
	return newTracker(kind, client, fetch, store)
}
