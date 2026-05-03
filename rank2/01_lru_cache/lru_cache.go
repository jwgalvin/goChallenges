package lru_cache

// LRUCache is a fixed-capacity cache that evicts the least-recently-used entry
// when full. Both Get and Put count as a use.
type LRUCache struct {
	// TODO: add fields
	// Hint: a map for O(1) lookup + a doubly linked list to track use order
}

// NewLRUCache creates a cache with the given capacity. Capacity >= 1.
func NewLRUCache(capacity int) *LRUCache {
	// TODO: implement
	return &LRUCache{}
}

// Get returns the value for key and true if found, or 0 and false if not.
// Marks key as most-recently-used.
func (c *LRUCache) Get(key int) (int, bool) {
	// TODO: implement
	return 0, false
}

// Put inserts or updates key with value, marking it as most-recently-used.
// If at capacity, the least-recently-used key is evicted first.
func (c *LRUCache) Put(key, value int) {
	// TODO: implement
}
