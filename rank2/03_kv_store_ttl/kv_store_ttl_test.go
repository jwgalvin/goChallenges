package kv_store_ttl

import (
	"testing"
	"time"
)

func TestKVStore_SetGet(t *testing.T) {
	s := NewKVStore()
	s.Set("name", "alice", 0)
	v, ok := s.Get("name")
	if !ok || v != "alice" {
		t.Errorf("Get(%q) = %q, %v; want %q, true", "name", v, ok, "alice")
	}
}

func TestKVStore_Miss(t *testing.T) {
	s := NewKVStore()
	if _, ok := s.Get("missing"); ok {
		t.Error("expected miss for unknown key")
	}
}

func TestKVStore_Delete(t *testing.T) {
	s := NewKVStore()
	s.Set("key", "val", 0)
	s.Delete("key")
	if _, ok := s.Get("key"); ok {
		t.Error("expected miss after delete")
	}
}

func TestKVStore_DeleteMissingKeyIsNoop(t *testing.T) {
	s := NewKVStore()
	s.Delete("nonexistent") // should not panic
}

func TestKVStore_TTLExpiry(t *testing.T) {
	s := NewKVStore()
	s.Set("key", "val", 50*time.Millisecond)

	if _, ok := s.Get("key"); !ok {
		t.Fatal("key should exist before TTL expires")
	}

	time.Sleep(100 * time.Millisecond)

	if _, ok := s.Get("key"); ok {
		t.Error("key should be expired after TTL")
	}
}

func TestKVStore_ZeroTTLNeverExpires(t *testing.T) {
	s := NewKVStore()
	s.Set("key", "val", 0)
	time.Sleep(50 * time.Millisecond)
	if _, ok := s.Get("key"); !ok {
		t.Error("key with zero TTL should never expire")
	}
}

func TestKVStore_OverwriteResetsExpiry(t *testing.T) {
	s := NewKVStore()
	s.Set("key", "v1", 50*time.Millisecond)
	time.Sleep(30 * time.Millisecond)
	s.Set("key", "v2", 200*time.Millisecond) // reset with a longer TTL

	time.Sleep(40 * time.Millisecond) // original TTL would have expired, new one hasn't
	v, ok := s.Get("key")
	if !ok || v != "v2" {
		t.Errorf("Get() = %q, %v; want %q, true", v, ok, "v2")
	}
}
