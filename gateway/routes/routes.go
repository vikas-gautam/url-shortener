package routes

import (
	"gateway/routes/handlers"
	"gateway/routes/handlers/auth"
	"gateway/routes/handlers/shortener"

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

	api.GET("/:url", shortener.ResolveURL)
	api.GET("/health", handlers.HealthCheck)
	api.POST("/signup", auth.Signup)
	api.POST("/login", auth.Login)
	api.POST("/shorturl", shortener.ShortURL)

}
