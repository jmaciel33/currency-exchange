package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func ReadConfig() error {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	return nil
}

func RedisClient() *redis.Client {
	ReadConfig()
	redisAddr := fmt.Sprintf("%s:%s", viper.Get("REDIS.REDIS_HOST"), viper.Get("REDIS.REDIS_PORT"))
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	return client
}