// controllers/hash_controller.go
package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	models "strongify-passgen-go-echo/model"
	views "strongify-passgen-go-echo/view"
	"time"

	"github.com/labstack/echo/v4"
)

// GenerateHash generates a hash based on the given text, secret, length, number of symbols, and numbers
func GenerateHash(c echo.Context) error {
	request := new(models.HashRequest)
	if err := c.Bind(request); err != nil {
		return views.RenderError(c, "Invalid request")
	}

	hash := generateHash(request.Text, request.Secret, request.Length, request.NumSymbols, request.NumNumbers)
	return views.RenderHash(c, hash)
}

// Helper function to generate the hash
func generateHash(text, secret string, length, numSymbols, numNumbers int) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(text))
	baseHash := hex.EncodeToString(h.Sum(nil))

	// Create the character pool
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>?/"

	// Shuffle and select characters
	rand.Seed(time.Now().UnixNano())
	pool := letters + numbers + symbols
	result := make([]byte, length)

	// Ensure the baseHash is used
	copy(result, baseHash[:length])

	// Add symbols
	for i := 0; i < numSymbols; i++ {
		index := rand.Intn(length)
		result[index] = symbols[rand.Intn(len(symbols))]
	}

	// Add numbers
	for i := 0; i < numNumbers; i++ {
		index := rand.Intn(length)
		result[index] = numbers[rand.Intn(len(numbers))]
	}

	// Fill the rest with random characters from the pool
	for i := 0; i < length; i++ {
		if result[i] == 0 {
			result[i] = pool[rand.Intn(len(pool))]
		}
	}

	return string(result)
}
