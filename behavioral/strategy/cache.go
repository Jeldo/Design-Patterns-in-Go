package main

type Cache struct {
	storage           map[string]string
	evictionAlgorithm EvictionAlgorithm
	capacity          int
	maxCapacity       int
}

func (c *Cache) SetEvictionAlgorithm(evictionAlgorithm EvictionAlgorithm) {
	c.evictionAlgorithm = evictionAlgorithm
}

func (c *Cache) Evict() {
	c.evictionAlgorithm.Evict(c)
	c.capacity--
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.Evict()
	}

	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) string {
	defer delete(c.storage, key)
	return c.storage[key]
}

func NewCache(
	evictionPolicy EvictionAlgorithm,
) *Cache {
	return &Cache{
		storage:           make(map[string]string),
		evictionAlgorithm: evictionPolicy,
		capacity:          0,
		maxCapacity:       2,
	}
}

type EvictionAlgorithm interface {
	Evict(cache *Cache)
}
