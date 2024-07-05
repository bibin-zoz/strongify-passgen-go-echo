// views/hash_view.go
package views

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HashResponse struct {
	Hash string `json:"hash"`
}

func RenderHash(c echo.Context, hash string) error {
	response := HashResponse{Hash: hash}
	return c.JSON(http.StatusOK, response)
}

func RenderError(c echo.Context, err string) error {
	return c.JSON(http.StatusBadRequest, map[string]string{"error": err})
}

func RenderSuccess(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, map[string]string{"message": message})
}
