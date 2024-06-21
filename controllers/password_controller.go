// controllers/password_controller.go
package controllers

import (
	"math/rand"
	models "strongify-passgen-go-echo/model"
	views "strongify-passgen-go-echo/view"
	"time"

	"github.com/labstack/echo/v4"
)

// GeneratePassword generates a password based on the given options
func GeneratePassword(c echo.Context) error {
	options := new(models.PasswordOptions)
	if err := c.Bind(options); err != nil {
		return views.RenderError(c, "Invalid request")

	}

	password := generatePassword(*options)
	return views.RenderPassword(c, password)
}

// Helper function to generate the password
func generatePassword(options models.PasswordOptions) string {
	var chars string
	if options.IncludeLowercase {
		chars += "abcdefghijklmnopqrstuvwxyz"
	}
	if options.IncludeUppercase {
		chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if options.IncludeNumbers {
		chars += "0123456789"
	}
	if options.IncludeSpecials {
		chars += "!@#$%^&*()_+-=[]{}|;:,.<>?/"
	}

	password := make([]byte, options.Length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))]
	}

	return string(password)
}

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
}
