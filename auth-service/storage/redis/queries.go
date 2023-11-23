package redis

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var (
	redisExpiration = 10 * time.Minute
)

// func ConnectionRedis(conn *redis.Client) {
// 	redisClient = conn
// }

//using DI ****************************************************************

type RedisStore interface {
	SetData(string, string) error
	GetData(string) (string, error)
	DelKey(string) error
}

func NewRedisStore(r *redis.Client) RedisStore {
	return &store{r}
}

// The actual store would contain some state. In this case it's the sql.db instance, that holds the connection to our database
type store struct {
	redis *redis.Client
}

//######################################################################

func (r *store) SetData(key, value string) error {

	err := r.redis.Set(context.Background(), key, value, redisExpiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *store) GetData(key string) (string, error) {
	val, err := r.redis.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *store) DelKey(key string) error {
	_, err := r.redis.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return nil
}
