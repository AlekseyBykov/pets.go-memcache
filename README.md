# Go In-Memory Cache

Simple in-memory cache written in Go with TTL (time-to-live) support and thread-safe operations.

## Features
- In-memory key-value storage
- Per-item TTL (expiration)
- Thread-safe (via RWMutex)
- Core `Set`, `Get`, and `Delete` operations

## Usage

```go
import "github.com/AlekseyBykov/pets.go-memcache/internal/cache/storage"

cache := storage.NewCache(5 * time.Second) // default TTL

cache.Set("key1", "value1", 5 * time.Second)
value, err := cache.Get("key1")
cache.Delete("key1")
```

## How TTL works
Each item in the cache can have its own TTL. After TTL expires, the item will be considered expired and won't be returned by Get().

## Requirements
Go 1.20+
