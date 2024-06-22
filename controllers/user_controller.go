package controllers

import (
	db "strongify-passgen-go-echo/database"
	"strongify-passgen-go-echo/domain"
	"strongify-passgen-go-echo/helpers"
	views "strongify-passgen-go-echo/view"

	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) error {
	signupData := new(domain.User)
	if err := c.Bind(&signupData); err != nil {
		return views.RenderError(c, "Invalid request")
	}
	if err := helpers.ValidateUser(*signupData); err != nil { 
		return views.RenderError(c, "Invalid request")
	}
	db.DB.Create(&signupData)
	return nil
}
