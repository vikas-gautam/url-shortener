package routes

import (
	"gateway/middleware"
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

	unprotected := app.Group("/api/v1")
	protected := app.Group("/api/v1")

	protected.Use(middleware.Authentication())

	app.GET("/:url", shortener.ResolveURL)
	unprotected.GET("/health", handlers.HealthCheck)
	unprotected.POST("/signup", auth.Signup)
	unprotected.POST("/login", auth.Login)
	protected.POST("/shorturl", shortener.ShortURL)
	unprotected.POST("/reset", auth.GenerateResetToken)
	unprotected.POST("/reset/:token", auth.ResetPassword)

}
