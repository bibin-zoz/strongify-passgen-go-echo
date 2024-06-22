// routes/routes.go
package routes

import (
	"strongify-passgen-go-echo/controllers"

	"github.com/labstack/echo/v4"
)

// InitRoutes initializes the routes for the application
func InitRoutes(e *echo.Echo) {
	e.POST("/generate-hash", controllers.GenerateHash)
}
