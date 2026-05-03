package kv_store_ttl

import "time"

// KVStore is a simple in-memory key-value store with optional per-key TTL.
type KVStore struct {
	// TODO: add fields
}

// NewKVStore returns an empty store.
func NewKVStore() *KVStore {
	// TODO: implement
	return &KVStore{}
}

// Set stores key with value. If ttl > 0 the entry expires after that duration.
// If ttl == 0 the entry never expires. Overwriting a key resets its expiry.
func (s *KVStore) Set(key, value string, ttl time.Duration) {
	// TODO: implement
}

// Get returns the value and true if key exists and has not expired.
func (s *KVStore) Get(key string) (string, bool) {
	// TODO: implement
	return "", false
}

// Delete removes key immediately regardless of TTL.
func (s *KVStore) Delete(key string) {
	// TODO: implement
}
