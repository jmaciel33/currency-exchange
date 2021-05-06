package persistence

import (
	"github.com/go-redis/redis"
	"github.com/jmaciel33/currency-exchange/currency-exchange-api/internal/business"
	"time"
)

type CacheRepositoryImpl struct {
	redis *redis.Client
}
func NewCacheRepository(redis *redis.Client) business.CacheRepository  {
	return CacheRepositoryImpl{redis}
}

func (c CacheRepositoryImpl) Set(key string, data string, expiration time.Duration) (string, error)  {
	return c.redis.Set(key, data, expiration).Result()
}

func (c CacheRepositoryImpl) Get(key string) (string, error) {
	return c.redis.Get(key).Result()
}




