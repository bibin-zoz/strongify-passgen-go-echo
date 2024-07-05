package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	db "strongify-passgen-go-echo/database"
	"strongify-passgen-go-echo/domain"
	"strongify-passgen-go-echo/helpers"
	models "strongify-passgen-go-echo/model"
	views "strongify-passgen-go-echo/view"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error {
	signupData := new(domain.User)

	if err := c.Bind(signupData); err != nil {
		return views.RenderError(c, "Invalid request")
	}

	if err := helpers.ValidateUser(*signupData); err != nil {
		return views.RenderError(c, "Invalid request")
	}

	if signupData.Email == "" {
		return views.RenderError(c, "Invalid email")
	}

	emailPattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailPattern).MatchString(signupData.Email) {
		return views.RenderError(c, "Invalid email format")
	}

	var emailCount int64
	if err := db.DB.Table("users").Where("email = ?", signupData.Email).Count(&emailCount).Error; err != nil {
		return views.RenderError(c, "Database error")
	}
	if emailCount > 0 {
		return views.RenderError(c, "Email already exists")
	}

	if signupData.Password == "" {
		return views.RenderError(c, "Password should not be empty")
	}
	if len(signupData.Password) < 6 {
		return views.RenderError(c, "Password length should be at least 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
	if err != nil {
		return views.RenderError(c, "Failed to hash password")
	}
	signupData.Password = string(hashedPassword)

	if err := db.DB.Create(&signupData).Error; err != nil {
		return views.RenderError(c, "Failed to create user")
	}

	return views.RenderSuccess(c, "Signup successful")
}

func LoginPost(c echo.Context) error {
	newMail := c.FormValue("email")
	newPassword := c.FormValue("password")

	var compare domain.User

	if newMail == "" {
		return views.RenderError(c, "Email should not be empty")
	}
	if newPassword == "" {
		return views.RenderError(c, "Password should not be empty")
	}

	// Query user from database
	if err := db.DB.Raw("SELECT id, password, user_name, email FROM users WHERE email=?", newMail).Scan(&compare).Error; err != nil {
		fmt.Println("Error querying the database:", err)
		return views.RenderError(c, "An error occurred while querying the database")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(compare.Password), []byte(newPassword)); err != nil {
		return views.RenderError(c, "Incorrect password")
	}

	claims := models.UserDetails{
		ID:        compare.ID,
		Firstname: compare.UserName,
		Email:     compare.Email,
	}

	accessToken, err := helpers.GenerateAccessToken(claims)
	if err != nil {
		fmt.Println("Error generating access token:", err)
		return views.RenderError(c, "Error generating access token")
	}
	refreshToken, err := helpers.GenerateRefreshToken(claims)
	if err != nil {
		fmt.Println("Error generating refresh token:", err)
		return views.RenderError(c, "Error generating refresh token")
	}

	response := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	return c.JSON(http.StatusOK, response)
}
