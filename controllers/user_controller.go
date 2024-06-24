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

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error {
	signupData := new(domain.User)

	// Bind the request body to signupData
	if err := c.Bind(signupData); err != nil {
		return views.RenderError(c, "Invalid request")
	}

	// Validate user data
	if err := helpers.ValidateUser(*signupData); err != nil {
		return views.RenderError(c, "Invalid request")
	}

	// Check if email is provided
	if signupData.Email == "" {
		return views.RenderError(c, "Invalid email")
	}

	// Check if email format is valid
	emailPattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailPattern).MatchString(signupData.Email) {
		return views.RenderError(c, "Invalid email format")
	}

	// Check if email already exists
	var emailCount int64
	if err := db.DB.Table("users").Where("email = ?", signupData.Email).Count(&emailCount).Error; err != nil {
		return views.RenderError(c, "Database error")
	}
	if emailCount > 0 {
		return views.RenderError(c, "Email already exists")
	}

	// Validate password
	if signupData.Password == "" {
		return views.RenderError(c, "Password should not be empty")
	}
	if len(signupData.Password) < 6 {
		return views.RenderError(c, "Password length should be at least 6 characters")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupData.Password), bcrypt.DefaultCost)
	if err != nil {
		return views.RenderError(c, "Failed to hash password")
	}
	signupData.Password = string(hashedPassword)

	// Create the user
	if err := db.DB.Create(&signupData).Error; err != nil {
		return views.RenderError(c, "Failed to create user")
	}

	// Return success response
	return views.RenderSuccess(c, "Signup successful")
}
func LoginPost(c *gin.Context) {
	newMail := c.PostForm("email")
	newPassword := c.PostForm("password")

	var compare domain.User

	// Validate email and password presence
	if newMail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email should not be empty"})
		return
	}
	if newPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password should not be empty"})
		return
	}

	// Query user from database
	if err := db.DB.Raw("SELECT ID, password, username, email FROM users WHERE email=?", newMail).Scan(&compare).Error; err != nil {
		fmt.Println("Error querying the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while querying the database"})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(compare.Password), []byte(newPassword)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}

	// Generate access token
	claims := models.Claims{
		ID:       compare.ID,
		Username: compare.UserName,
		Email:    compare.Email,
	}

	accessToken, err := helpers.GenerateAccessToken(claims)
	if err != nil {
		fmt.Println("Error generating access token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating access token"})
		return
	}

	// Create JSON response with token
	response := gin.H{
		"access_token": accessToken,
	}

	// Return JSON response with token
	c.JSON(http.StatusOK, response)
}
