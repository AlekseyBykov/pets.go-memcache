package storage

import (
	"fmt"
	"github.com/AlekseyBykov/pets.go-memcache/internal/cache/validation"
	"github.com/AlekseyBykov/pets.go-memcache/internal/utils"
	"sync"
	"time"
)

type Cache struct {
	sync    sync.RWMutex
	ttl     time.Duration
	storage map[string]cacheItem
}

func NewCache(lifeTime time.Duration) *Cache {
	return &Cache{
		ttl:     lifeTime,
		storage: make(map[string]cacheItem),
	}
}

func (c *Cache) Set(key string, value any, ttl time.Duration) error {
	if err := validation.ValidateKey(key); err != nil {
		return err
	}

	if err := validation.ValidateValue(value); err != nil {
		return err
	}

	if err := validation.ValidateItemTtl(ttl); err != nil {
		return err
	}

	c.sync.Lock()
	defer c.sync.Unlock()

	c.storage[key] = cacheItem{
		value:   value,
		expired: utils.GetExpirationTime(ttl),
	}

	return nil
}

func (c *Cache) Get(key string) (any, error) {
	if err := validation.ValidateKey(key); err != nil {
		return nil, err
	}

	c.sync.RLock()
	defer c.sync.RUnlock()

	item, found := c.storage[key]
	if !found {
		return nil, validation.GetKeyNotFoundError(key)
	}

	if item.IsExpired() {
		return nil, validation.GetItemExpiredError(key)
	}

	return item.value, nil
}

func (c *Cache) Delete(key string) error {
	if err := validation.ValidateKey(key); err != nil {
		return err
	}

	c.sync.Lock()
	defer c.sync.Unlock()

	if _, found := c.storage[key]; !found {
		return validation.GetKeyNotFoundError(key)
	}

	delete(c.storage, key)

	return nil
}

func (c *Cache) StartJanitor(stop <-chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("[Janitor] Running cleanup...")
			c.cleanupExpiredItems()
		case <-stop:
			fmt.Printf("[Janitor] Stopping janitor")
			return
		}
	}
}

func (c *Cache) cleanupExpiredItems() {
	c.sync.Lock()
	defer c.sync.Unlock()

	for key, item := range c.storage {
		if item.IsExpired() {
			fmt.Printf("[Janitor] Removing expired key: %s\n", key)
			delete(c.storage, key)
		}
	}
}
