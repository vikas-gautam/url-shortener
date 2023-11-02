package redis

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var (
	redisClient     *redis.Client
	redisExpiration = 10 * time.Minute
)

func ConnectionRedis(conn *redis.Client) {
	redisClient = conn
}

func SetData(key, value string) error {
	err := redisClient.Set(context.Background(), key, value, redisExpiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetData(key string) (string, error) {
	val, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func DelKey(key string) error {
	_, err := redisClient.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return nil
}
