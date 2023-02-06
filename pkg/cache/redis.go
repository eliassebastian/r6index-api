package cache

import (
	"context"
	"time"

	"github.com/eliassebastian/r6index-api/pkg/utils"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type CacheStore struct {
	redis *redis.Client
	cache *cache.Cache
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

func (c *CacheStore) Once(key string, set interface{}, obj interface{}) error {
	return c.cache.Once(&cache.Item{
		Key:   key,
		Value: set,
		Do: func(*cache.Item) (interface{}, error) {
			return obj, nil
		},
	})
}

func (c *CacheStore) SetOnce(ctx context.Context, key string, set interface{}) error {
	return c.cache.Once(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: set,
	})
}

func (c *CacheStore) GetOnce(key string, set interface{}) error {
	return c.cache.Once(&cache.Item{
		Key:   key,
		Value: set,
	})
}

func (c *CacheStore) GetCache(ctx context.Context, key string, value interface{}) error {
	return c.cache.Get(ctx, key, value)
}

// func (c *CacheStore) Set(ctx context.Context, key string, value interface{}) error {
// 	return c.cache.Set(&cache.Item{
// 		Ctx:   ctx,
// 		Key:   key,
// 		Value: value,
// 		//TTL:   time.Hour,
// 	})
// }

func (c *CacheStore) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	return c.redis.Set(ctx, key, value, ttl).Err()
}

func (c *CacheStore) Get(ctx context.Context, key string, obj interface{}) (string, error) {
	return c.redis.Get(ctx, key).Result()
}

func (c *CacheStore) Close() {
	c.redis.Close()
}
