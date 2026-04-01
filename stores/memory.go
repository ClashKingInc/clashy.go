package stores

import (
	"context"
	"sync"

	"clashy.go/events"
)

type MemoryStore struct {
	mu    sync.RWMutex
	items map[string][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{items: make(map[string][]byte)}
}

func NewStore() *MemoryStore {
	return NewMemoryStore()
}

func memorySnapshotKey(kind, key string) string {
	return kind + ":" + key
}

func (s *MemoryStore) Load(_ context.Context, kind, key string) (events.Snapshot, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value := s.items[memorySnapshotKey(kind, key)]
	if len(value) == 0 {
		return events.Snapshot{}, events.ErrSnapshotNotFound
	}
	return events.Snapshot{
		Kind: kind,
		Key:  key,
		Data: append([]byte(nil), value...),
	}, nil
}

func (s *MemoryStore) Save(_ context.Context, snapshot events.Snapshot) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[memorySnapshotKey(snapshot.Kind, snapshot.Key)] = append([]byte(nil), snapshot.Data...)
	return nil
}

func (s *MemoryStore) Delete(_ context.Context, kind, key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, memorySnapshotKey(kind, key))
	return nil
}
