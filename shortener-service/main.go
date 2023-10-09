package main

import (
	"os"
	"shortener-service/routes"
	"shortener-service/storage/db"
	"shortener-service/storage/redis"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	apiPort = "8081"
)

// const DSN = "host=postgres port=5432 user=postgres password=password dbname=shortener sslmode=disable timezone=UTC connect_timeout=5"

func main() {

	//setting up required env
	os.Setenv("DSN", "host=localhost port=5432 user=postgres password=password dbname=shortener sslmode=disable timezone=UTC connect_timeout=5")
	os.Setenv("REDIS_ENDPOINT", "localhost")

	//connect to database
	connDB, err := db.ConnectToDB()
	if connDB == nil {
		logrus.Panic("Can't connect to database postgres", err)
	}
	//connect to database
	connRedis, err := redis.ConnectToRedis()
	if connRedis == nil {
		logrus.Panic("Can't connect to database postgres", err)
	}

	// if u not using method
	db.Connection(connDB)
	redis.ConnectionRedis(connRedis)

	app := gin.Default()

	routes.SetupRoutes(app)

	app.Run("localhost:" + apiPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	logrus.Infof("Starting shortner service on port %s\n", apiPort)
}
