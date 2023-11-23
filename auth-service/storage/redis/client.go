package redis

import (
	"auth-service/config"
	"context"
	"fmt"
	"time"

	redis "github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()
var counts int64

//#####################################################################

func RedisClient(endpoint string) (*redis.Client, error) {
	fmt.Println("Go Redis Client")
	redisEndpoint := endpoint

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisEndpoint + ":6379",
		Password: "",
		DB:       0,
	})

	err := redisClient.Ping(ctx).Err()
	if err != nil {
		logrus.Errorln("Failed to connect to redis:", err)
		return nil, err
	}

	logrus.Info("Connected to Redis")
	return redisClient, nil
}

func NewRedisClient(appConfig config.Config) (*redis.Client, error) {

	for {
		connection, err := RedisClient(appConfig.REDIS_ENDPOINT)
		if err != nil {
			logrus.Warnln("Redis is not ready yet...", err)
			counts++

		} else {
			logrus.Infof("Connected to Redis")
			return connection, nil
		}

		if counts > 10 {
			logrus.Error(err)
			return connection, err
		}

		logrus.Infof("Backing off for 2 seconds ..")
		time.Sleep(2 * time.Second)
		continue

	}
}
