package main

import (
	"gateway/routes"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	apiPort = "9090"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// 	os.Exit(1)
	// }

	app := gin.Default()

	routes.SetupRoutes(app)

	app.Run("0.0.0.0:" + apiPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	logrus.Infof("Starting gateway  on port %s\n", apiPort)
}
