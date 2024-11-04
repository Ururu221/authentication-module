package routes

import (
	"github.com/labstack/echo/v4"
	"structured/handlers"
	"structured/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/register", handlers.RegisterUser)
	e.POST("/login", handlers.LoginAndCreateToken)

	protected := e.Group("/protected")
	protected.Use(middleware.JWTMiddleware)
	protected.GET("", handlers.SecretRoute)
}
