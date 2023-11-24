package main

import (
	"auth-service/config"
	"auth-service/routes"
	"auth-service/routes/handlers"
	"auth-service/storage/db"
	"auth-service/storage/redis"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	configInfo := Initialize()

	app := gin.Default()

	routes.SetupRoutes(app)

	if configInfo.APP_PORT == "" {
		configInfo.APP_PORT = "8082"
	}

	app.Run("0.0.0.0:" + configInfo.APP_PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	logrus.Infof("Starting AUTH service on port %s\n", configInfo.APP_PORT)
}

func Initialize() config.Config {

	configInfo := config.Initialize()

	//connect to database
	dbConn, err := db.NewdbConnection(configInfo)
	if dbConn == nil {
		logrus.Panic("Can't connect to database postgres", err)
	}

	//connect to redis
	connRedis, err := redis.NewRedisClient(configInfo)
	if connRedis == nil {
		logrus.Panic("Can't connect to redis", err)
	}

	// Create a store dependency with the db AND redis connection
	RedisStoreValue := redis.NewRedisStore(connRedis)
	DBStoreValue := db.NewDBStore(dbConn)

	//################################################################

	service := &handlers.Service{
		RedisStore: RedisStoreValue,
		DbStore:    DBStoreValue,
	}

	handlers.NewRepo(service)

	return configInfo

	// defer close(app.MailChan)
	// fmt.Println("Starting mail listner")
	// internal.ListenForMail()

	// if u r not using method
	// db.Connection(conn)
	// redis.ConnectionRedis(connRedis)

	//common func to take these all connections out of main
	// db.Connectiondb(appConf)

}
