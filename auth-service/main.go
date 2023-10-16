package main

import (
	"auth-service/routes"
	"auth-service/storage/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	apiPort = "8082"
)

func main() {

	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// 	os.Exit(1)
	// }

	//connect to database
	conn, err := db.ConnectToDB()
	if conn == nil {
		logrus.Panic("Can't connect to database postgres", err)
	}
	// if u not using method
	db.Connection(conn)

	app := gin.Default()

	routes.SetupRoutes(app)

	app.Run("0.0.0.0:" + apiPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	logrus.Infof("Starting AUTH service on port %s\n", apiPort)
}
