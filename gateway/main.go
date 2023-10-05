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
	app := gin.Default()

	routes.SetupRoutes(app)

	app.Run("localhost:" + apiPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	logrus.Infof("Starting gateway  on port %s\n", apiPort)
}
