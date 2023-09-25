package main

import (
	"os"
	"shortener-service/handlers"
	"shortener-service/storage/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	apiPort = "8081"
)

// const DSN = "host=postgres port=5432 user=postgres password=password dbname=shortener sslmode=disable timezone=UTC connect_timeout=5"

func main() {
	os.Setenv("DSN", "host=localhost port=5432 user=postgres password=password dbname=shortener sslmode=disable timezone=UTC connect_timeout=5")

	//connect to database
	conn, err := db.ConnectToDB()
	if conn == nil {
		logrus.Panic("Can't connect to database postgres", err)
	}
	// if u not using method
	db.Connection(conn)

	router := gin.Default()

	router.GET("/:url", handlers.ResolveURL)
	router.GET("/health", handlers.HealthCheck)

	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)
	router.POST("/shorturl", handlers.ShortURL)

	router.Run("localhost:" + apiPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	logrus.Infof("Starting shortner service on port %s\n", apiPort)
}
