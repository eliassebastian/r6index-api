package cache

import (
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type CacheStore struct {
	redis *redis.Client
	cache *cache.Cache
}

func New() (*CacheStore, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr: "172.0.0.1:6379",
	})

	cdb := cache.New(&cache.Options{
		Redis: rdb,
	})

	return &CacheStore{
		redis: rdb,
		cache: cdb,
	}, nil
}
