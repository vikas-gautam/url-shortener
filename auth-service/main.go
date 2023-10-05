package main

import (
	"auth-service/routes"
	"auth-service/storage/db"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	apiPort = "8082"
)

const DSN = "host=postgres port=5432 user=postgres password=password dbname=shortener sslmode=disable timezone=UTC connect_timeout=5"

func main() {
	os.Setenv("DSN", "host=localhost port=5432 user=postgres password=password dbname=shortener sslmode=disable timezone=UTC connect_timeout=5")

	//connect to database
	conn, err := db.ConnectToDB()
	if conn == nil {
		logrus.Panic("Can't connect to database postgres", err)
	}
	// if u not using method
	db.Connection(conn)

	app := gin.Default()

	routes.SetupRoutes(app)

	app.Run("localhost:" + apiPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	logrus.Infof("Starting AUTH service on port %s\n", apiPort)
}
