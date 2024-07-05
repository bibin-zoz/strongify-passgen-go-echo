// main.go
package main

import (
	"fmt"
	"os"
	"strongify-passgen-go-echo/database"
	"strongify-passgen-go-echo/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())



	routes.InitRoutes(e)
	database.InitDB()
	secret := os.Getenv("secret")
	fmt.Println(secret)

	e.Logger.Fatal(e.Start(":8080"))
}
