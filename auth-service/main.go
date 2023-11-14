package main

import (
	"auth-service/config"
	"auth-service/routes"
	"auth-service/storage/db"
	"auth-service/storage/redis"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	apiPort = "8082"
)

func main() {

	appConfig := config.Initialize()

	//connect to database
	_, err := db.NewdbConnection(appConfig)
	if err != nil {
		logrus.Panic("Can't connect to database postgres", err)
	}

	// conn, err := db.ConnectToDB(appConfig)
	// if conn == nil {
	// 	logrus.Panic("Can't connect to database postgres", err)
	// }

	//connect to redis

	connRedis, err := redis.NewRedisClient(appConfig)
	if connRedis == nil {
		logrus.Panic("Can't connect to redis", err)
	}

	// connRedis, err := redis.ConnectToRedis(appConfig)
	// if connRedis == nil {
	// 	logrus.Panic("Can't connect to database postgres", err)
	// }

	// defer close(app.MailChan)
	// fmt.Println("Starting mail listner")
	// internal.ListenForMail()

	// if u not using method
	// db.Connection(conn)
	// redis.ConnectionRedis(connRedis)

	app := gin.Default()

	routes.SetupRoutes(app)

	app.Run("0.0.0.0:" + apiPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	logrus.Infof("Starting AUTH service on port %s\n", apiPort)
}
