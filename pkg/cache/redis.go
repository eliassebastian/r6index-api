package cache

import (
	"context"
	"time"

	"github.com/eliassebastian/r6index-api/cmd/api/models"
	"github.com/eliassebastian/r6index-api/pkg/utils"
	"github.com/redis/go-redis/v9"
	"github.com/shamaton/msgpackgen/msgpack"
	"golang.org/x/sync/singleflight"
)

type CacheStore struct {
	redis *redis.Client
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

	return &CacheStore{
		redis: rdb,
	}, nil
}

func (c *CacheStore) Set(ctx context.Context, key string, value *models.ProfileCache, ttl time.Duration) error {
	err, _, _ := c.group.Do(key, func() (interface{}, error) {

		b, err := msgpack.Marshal(value)
		if err != nil {
			return err, nil
		}

		err = c.redis.Set(ctx, key, b, ttl).Err()
		if err != nil {
			return err, nil
		}

		return err, nil
	})

	//err=interface conversion
	if err == nil {
		return nil
	}

	return err.(error)
}

func (c *CacheStore) Get(ctx context.Context, key string, value *models.ProfileCache) error {
	resp, err := c.redis.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	err = msgpack.Unmarshal(resp, value)
	if err != nil {
		return err
	}

	return nil
}

func (c *CacheStore) Close() {
	c.redis.Close()
}
