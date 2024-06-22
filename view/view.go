// views/hash_view.go
package views

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HashResponse is the structure for the hash response
type HashResponse struct {
	Hash string `json:"hash"`
}

// RenderHash renders the hash response
func RenderHash(c echo.Context, hash string) error {
	response := HashResponse{Hash: hash}
	return c.JSON(http.StatusOK, response)
}

// RenderError renders an error response
func RenderError(c echo.Context, err string) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": err})
}
