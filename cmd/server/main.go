package main

import (
	"github.com/labstack/echo/v4"
	"structured/config"
	"structured/models"
	"structured/routes"
)

func main() {
	e := echo.New()
	db := config.InitDB()
	models.MigrateDB(db)

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":1488"))
}
