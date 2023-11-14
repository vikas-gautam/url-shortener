package redis

import (
	"context"
	"time"
)

var (
	redisExpiration = 10 * time.Minute
)

// func ConnectionRedis(conn *redis.Client) {
// 	redisClient = conn
// }

func (r *RedisInfo) SetData(key, value string) error {

	err := r.redisClient.Set(context.Background(), key, value, redisExpiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisInfo) GetData(key string) (string, error) {
	val, err := r.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisInfo) DelKey(key string) error {
	_, err := r.redisClient.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return nil
}
