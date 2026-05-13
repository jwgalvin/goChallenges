package kv_store_ttl

import (
	"sync"
	"time"
)

// KVStore is a simple in-memory key-value store with optional per-key TTL.
type entry struct {
	value     string
	expiresAt time.Time
}

type KVStore struct {
	mu   sync.RWMutex
	data map[string]entry
}

// NewKVStore returns an empty store.
func NewKVStore() *KVStore {
	return &KVStore{data: make(map[string]entry)}
}

// Set stores key with value. If ttl > 0 the entry expires after that duration.
// If ttl == 0 the entry never expires. Overwriting a key resets its expiry.
func (s *KVStore) Set(key, value string, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var expiresAt time.Time

	if ttl > 0 {
		expiresAt = time.Now().Add(ttl)
	}

	s.data[key] = entry{
		value:     value,
		expiresAt: expiresAt,
	}
}

// Get returns the value and true if key exists and has not expired.
func (s *KVStore) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	e, ok := s.data[key]
	if !ok {
		return "", false
	}
	if !e.expiresAt.IsZero() && time.Now().After(e.expiresAt) {
		return "", false
	}

	return e.value, true
}

// Delete removes key immediately regardless of TTL.
func (s *KVStore) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
