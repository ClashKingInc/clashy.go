package stores

import (
	"context"
	"errors"

	"clashy.go/events"
)

type MongoDocumentStore interface {
	FindSnapshot(ctx context.Context, key string) ([]byte, error)
	UpsertSnapshot(ctx context.Context, key string, value []byte) error
	DeleteSnapshot(ctx context.Context, key string) error
}

type MongoCursorStore interface {
	ScanTags(ctx context.Context, cursor string, limit int) (tags []string, nextCursor string, err error)
}

type MongoStore struct {
	client MongoDocumentStore
}

func NewMongoStore(client MongoDocumentStore) *MongoStore {
	return &MongoStore{client: client}
}

func (s *MongoStore) Load(ctx context.Context, kind, key string) (events.Snapshot, error) {
	if s == nil || s.client == nil {
		return events.Snapshot{}, errors.New("mongo document store is nil")
	}
	value, err := s.client.FindSnapshot(ctx, kind+":"+key)
	if err != nil || len(value) == 0 {
		if err == nil {
			err = events.ErrSnapshotNotFound
		}
		return events.Snapshot{}, err
	}
	return events.Snapshot{Kind: kind, Key: key, Data: value}, nil
}

func (s *MongoStore) Save(ctx context.Context, snapshot events.Snapshot) error {
	if s == nil || s.client == nil {
		return errors.New("mongo document store is nil")
	}
	return s.client.UpsertSnapshot(ctx, snapshot.Kind+":"+snapshot.Key, snapshot.Data)
}

func (s *MongoStore) Delete(ctx context.Context, kind, key string) error {
	if s == nil || s.client == nil {
		return errors.New("mongo document store is nil")
	}
	return s.client.DeleteSnapshot(ctx, kind+":"+key)
}

type TagSource struct {
	client MongoCursorStore
}

func NewMongoTagSource(client MongoCursorStore) events.TagSource {
	return &TagSource{client: client}
}

func (s *TagSource) ForEach(ctx context.Context, fn func(tag string) error) error {
	if s == nil || s.client == nil {
		return errors.New("mongo cursor store is nil")
	}
	cursor := ""
	for {
		tags, nextCursor, err := s.client.ScanTags(ctx, cursor, 1000)
		if err != nil {
			return err
		}
		for _, tag := range tags {
			if err := fn(tag); err != nil {
				return err
			}
		}
		if nextCursor == "" {
			return nil
		}
		cursor = nextCursor
	}
}
