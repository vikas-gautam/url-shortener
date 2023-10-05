package routes

import (
	"auth-service/routes/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET,POST,PUT,DELETE,OPTIONS"},
		AllowHeaders:     []string{"Accept,Authorization,Content-Type,X-CSRF-TOKEN"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	api := app.Group("/api/v1")

	api.GET("/health", handlers.HealthCheck)

	api.POST("/signup", handlers.Signup)
	api.POST("/login", handlers.Login)

}
