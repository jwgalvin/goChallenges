package lru_cache

import "testing"

func TestLRUCache_Miss(t *testing.T) {
	c := NewLRUCache(2)
	if _, ok := c.Get(99); ok {
		t.Error("expected cache miss on empty cache")
	}
}

func TestLRUCache_PutThenGet(t *testing.T) {
	c := NewLRUCache(3)
	c.Put(1, 10)
	c.Put(2, 20)
	v, ok := c.Get(1)
	if !ok || v != 10 {
		t.Errorf("Get(1) = %d, %v; want 10, true", v, ok)
	}
}

func TestLRUCache_UpdateExistingKey(t *testing.T) {
	c := NewLRUCache(2)
	c.Put(1, 10)
	c.Put(1, 99)
	v, ok := c.Get(1)
	if !ok || v != 99 {
		t.Errorf("after update: Get(1) = %d, %v; want 99, true", v, ok)
	}
}

func TestLRUCache_EvictsLRU(t *testing.T) {
	c := NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3) // evicts key 1 (least recently used)

	if _, ok := c.Get(1); ok {
		t.Error("key 1 should have been evicted")
	}
	if v, ok := c.Get(2); !ok || v != 2 {
		t.Errorf("Get(2) = %d, %v; want 2, true", v, ok)
	}
	if v, ok := c.Get(3); !ok || v != 3 {
		t.Errorf("Get(3) = %d, %v; want 3, true", v, ok)
	}
}

func TestLRUCache_GetPromotesToFront(t *testing.T) {
	c := NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Get(1)    // key 1 is now most-recently-used
	c.Put(3, 3) // should evict key 2, not key 1

	if _, ok := c.Get(2); ok {
		t.Error("key 2 should have been evicted")
	}
	if _, ok := c.Get(1); !ok {
		t.Error("key 1 should still be present")
	}
}

func TestLRUCache_PutPromotesToFront(t *testing.T) {
	c := NewLRUCache(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(1, 11) // updating key 1 makes it most-recently-used
	c.Put(3, 3)  // should evict key 2

	if _, ok := c.Get(2); ok {
		t.Error("key 2 should have been evicted")
	}
	if v, ok := c.Get(1); !ok || v != 11 {
		t.Errorf("Get(1) = %d, %v; want 11, true", v, ok)
	}
}

func TestLRUCache_CapacityOne(t *testing.T) {
	c := NewLRUCache(1)
	c.Put(1, 1)
	c.Put(2, 2) // evicts 1

	if _, ok := c.Get(1); ok {
		t.Error("key 1 should have been evicted")
	}
	if v, ok := c.Get(2); !ok || v != 2 {
		t.Errorf("Get(2) = %d, %v; want 2, true", v, ok)
	}
}
