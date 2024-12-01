package cache

import (
	"WBL0/internal/messages"
	"sync"
)

type Cache interface {
	Set(key string, value messages.Order) error
	Get(key string) (interface{}, bool)
}

type cache struct {
	mutex sync.Mutex
	store map[string]messages.Order
}

func NewCache() Cache {
	return &cache{store: make(map[string]messages.Order)}
}

func (c *cache) Set(key string, value messages.Order) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.store[key] = value
	return nil
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, ok := c.store[key]
	return value, ok
}
