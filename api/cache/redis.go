package cache

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
)

var redisClient *redis.Client

func init() {
	// Initializing Redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "redis:6379"
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}

// RedisCache is a struct which holds some defaults for cache
// entries, such as the ttl and a prefix for each key. The idea here is that
// you have can have multiple instances of this struct and they all use the
// one client in this pacakge.
type RedisCache struct {
	keyPrefix string
	ttl time.Duration
}

// Add adds a key-value pair to the cache. This is specifically made to
// implement the PersistedQueryCache of gqlgen, so you can pass an instance
// of RedisCache to gqlgen. If you are wanting to manipulate the redis cache
// more directly, you may wish to use Set instead.
func (c *RedisCache) Add(ctx context.Context, key string, value interface{}) {
	redisClient.Set(c.keyPrefix + key, value, c.ttl)
}

// Set adds a key-value pair to the cache. Main difference between this and
// Add is that it returns an error, which makes it more useful than Add outside
// of using it with gqlgen.
func (c *RedisCache) Set(key string, value interface{}) error {
	return redisClient.Set(c.keyPrefix + key, value, c.ttl).Err()
}

// Gets a key-value pair from the cache. Part of the implementation of gqlgen's
// PersistedQueryCache so you can pass an isntance of RedisCache to a gqlgen
// handler.
func (c *RedisCache) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := redisClient.Get(c.keyPrefix + key).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}

var ErrRedisCacheNotInstantiated = errors.New("redis client not instantiated, call NewRedisCacheInstance first")

func NewRedisCacheInstance(keyPrefix string, ttl time.Duration) (*RedisCache, error) {
	if redisClient == nil {
		return nil, ErrRedisCacheNotInstantiated
	}

	err := redisClient.Ping().Err()
	if err != nil {
		return nil, err
	}

	return &RedisCache{keyPrefix: keyPrefix, ttl: ttl}, nil
}