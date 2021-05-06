package business

import "time"

type CacheRepository interface {
	Set(key string, data string, expiration time.Duration) (string, error)
	Get(key string) (string, error)}