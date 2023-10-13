package main

import (
	"log"
	"os"
	"shortener-service/routes"
	"shortener-service/storage/db"
	"shortener-service/storage/redis"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	apiPort = "8081"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}

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

	app.Run("0.0.0.0:" + apiPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	logrus.Infof("Starting shortner service on port %s\n", apiPort)
}
