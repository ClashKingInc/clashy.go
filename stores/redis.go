package stores

import (
	"context"
	"errors"
	"time"

	"clashy.go/events"
)

type RedisKVClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
	Del(ctx context.Context, key string) error
}

type RedisScanClient interface {
	SScan(ctx context.Context, key, cursor string, count int) (members []string, nextCursor string, err error)
}

type RedisStore struct {
	client RedisKVClient
	ttl    time.Duration
}

func NewRedisStore(client RedisKVClient, ttl time.Duration) *RedisStore {
	return &RedisStore{client: client, ttl: ttl}
}

func redisSnapshotKey(kind, key string) string {
	return kind + ":" + key
}

func (s *RedisStore) Load(ctx context.Context, kind, key string) (events.Snapshot, error) {
	if s == nil || s.client == nil {
		return events.Snapshot{}, errors.New("redis store client is nil")
	}
	value, err := s.client.Get(ctx, redisSnapshotKey(kind, key))
	if err != nil || value == "" {
		if err == nil {
			err = events.ErrSnapshotNotFound
		}
		return events.Snapshot{}, err
	}
	return events.Snapshot{Kind: kind, Key: key, Data: []byte(value)}, nil
}

func (s *RedisStore) Save(ctx context.Context, snapshot events.Snapshot) error {
	if s == nil || s.client == nil {
		return errors.New("redis store client is nil")
	}
	return s.client.Set(ctx, redisSnapshotKey(snapshot.Kind, snapshot.Key), string(snapshot.Data), s.ttl)
}

func (s *RedisStore) Delete(ctx context.Context, kind, key string) error {
	if s == nil || s.client == nil {
		return errors.New("redis store client is nil")
	}
	return s.client.Del(ctx, redisSnapshotKey(kind, key))
}

type SetTagSource struct {
	client RedisScanClient
	key    string
}

func NewRedisSetTagSource(client RedisScanClient, key string) events.TagSource {
	return &SetTagSource{client: client, key: key}
}

func (s *SetTagSource) ForEach(ctx context.Context, fn func(tag string) error) error {
	if s == nil || s.client == nil {
		return errors.New("redis scan client is nil")
	}
	cursor := "0"
	for {
		members, nextCursor, err := s.client.SScan(ctx, s.key, cursor, 1000)
		if err != nil {
			return err
		}
		for _, member := range members {
			if err := fn(member); err != nil {
				return err
			}
		}
		if nextCursor == "" || nextCursor == "0" {
			return nil
		}
		cursor = nextCursor
	}
}
