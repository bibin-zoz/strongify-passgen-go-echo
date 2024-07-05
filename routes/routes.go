// routes/routes.go
package routes

import (
	"strongify-passgen-go-echo/controllers"

	"github.com/labstack/echo/v4"
)

// InitRoutes initializes the routes for the application
func InitRoutes(e *echo.Echo) {
	e.POST("/signup", controllers.Signup)
	e.POST("/login", controllers.LoginPost)

	e.POST("/generate-hash", controllers.GenerateHash)

	e.GET("/phrase", controllers.GetPhrases)
	e.POST("/phrase", controllers.AddPhrase)
}
