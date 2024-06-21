// views/password_view.go
package views

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PasswordResponse is the structure for the password response
type PasswordResponse struct {
	Password string `json:"password"`
}

// RenderPassword renders the password response
func RenderPassword(c echo.Context, password string) error {
	response := PasswordResponse{Password: password}
	return c.JSON(http.StatusOK, response)
}

// RenderError renders an error response
func RenderError(c echo.Context, err string) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": err})
}
