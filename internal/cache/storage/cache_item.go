package storage

import (
	"github.com/AlekseyBykov/pets.go-memcache/internal/utils"
	"time"
)

type cacheItem struct {
	value   any
	expired time.Time
}

func (c cacheItem) IsExpired() bool {
	return utils.GetCurrentTime().After(c.expired)
}
