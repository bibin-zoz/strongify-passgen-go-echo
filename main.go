// main.go
package main

import (
	"strongify-passgen-go-echo/database"
	"strongify-passgen-go-echo/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize routes

	routes.InitRoutes(e)
	database.InitDB()

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
