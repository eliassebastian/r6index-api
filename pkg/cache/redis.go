package cache

import (
	"context"
	"time"

	"github.com/eliassebastian/r6index-api/pkg/utils"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/singleflight"
)

type CacheStore struct {
	redis *redis.Client
	cache *cache.Cache
	group singleflight.Group
}

func New(ctx context.Context) (*CacheStore, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     utils.GetEnv("REDIS_URL", "localhost:6379"),
		Password: utils.GetEnv("REDIS_PASSWORD", ""),
		DB:       0,
	})

	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	cdb := cache.New(&cache.Options{
		Redis: rdb,
	})

	return &CacheStore{
		redis: rdb,
		cache: cdb,
	}, nil
}

func (c *CacheStore) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	err, _, _ := c.group.Do(key, func() (interface{}, error) {
		err := c.cache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   key,
			Value: value,
			TTL:   ttl,
		})

		return err, nil
	})

	//err=interface conversion
	if err == nil {
		return nil
	}

	return err.(error)
}

func (c *CacheStore) Get(ctx context.Context, key string, value interface{}) error {
	return c.cache.Get(ctx, key, value)
}

func (c *CacheStore) Close() {
	c.redis.Close()
}
