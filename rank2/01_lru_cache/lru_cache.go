package lru_cache

// LRUCache is a fixed-capacity cache that evicts the least-recently-used entry
// when full. Both Get and Put count as a use.

type node struct {
	key, val   int
	prev, next *node
}
type LRUCache struct {
	capacity int
	cache    map[int]*node
	head     *node
	tail     *node
}

// NewLRUCache creates a cache with the given capacity. Capacity >= 1.
func NewLRUCache(capacity int) *LRUCache {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node),
		head:     head,
		tail:     tail,
	}
}

// Get returns the value for key and true if found, or 0 and false if not.
// Marks key as most-recently-used.
func (c *LRUCache) Get(key int) (int, bool) {
	n, ok := c.cache[key]
	if !ok {
		return 0, false
	}

	c.moveToFront(n)
	return n.val, true
}

// Put inserts or updates key with value, marking it as most-recently-used.
// If at capacity, the least-recently-used key is evicted first.
func (c *LRUCache) Put(key, value int) {
	if n, ok := c.cache[key]; ok {
		n.val = value
		c.moveToFront(n)
		return
	}

	n := &node{key: key, val: value}
	c.cache[key] = n
	c.insertAtFront(n)

	if len(c.cache) > c.capacity {
		lru := c.tail.prev
		c.removeNode(lru)
		delete(c.cache, lru.key)
	}
}

func (c *LRUCache) insertAtFront(n *node) {
	n.next = c.head.next
	n.prev = c.head
	c.head.next.prev = n
	c.head.next = n
}

func (c *LRUCache) removeNode(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRUCache) moveToFront(n *node) {
	c.removeNode(n)
	c.insertAtFront(n)
}
