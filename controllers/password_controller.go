
package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	db "strongify-passgen-go-echo/database"
	"strongify-passgen-go-echo/domain"
	models "strongify-passgen-go-echo/model"
	views "strongify-passgen-go-echo/view"

	"github.com/labstack/echo/v4"
)

func GenerateHash(c echo.Context) error {

	request := new(models.HashRequest)

	if err := c.Bind(request); err != nil {
		return views.RenderError(c, "Invalid request")
	}
	var phrase domain.WordPhrase
	result := db.DB.Where("ID=?", request.PhraseId).Find(&phrase)
	if result.Error != nil {
		return views.RenderError(c, "faield to get the details")
	}
	fmt.Println(request)
	if request.Length < 6 {
		return views.RenderError(c, "Password length must be at least 6 characters")
	}
	if request.NumNumbers+request.NumSymbols >= request.Length-3 {

		return views.RenderError(c, "Reduce number of symbols,number or Increase password length")
	}
	secret := os.Getenv("secret")
	hash := generateHash(phrase.Phrase+request.Text, secret, request.Length, request.NumSymbols, request.NumNumbers)
	return views.RenderHash(c, hash)
}

func generateHash(text, secret string, length, numSymbols, numNumbers int) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(text))
	baseHash := hex.EncodeToString(h.Sum(nil))

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>?/"

	result := make([]byte, length)
	copy(result, baseHash[:length])

	addCharacters := func(pool string, count int, offset int) {
		for i := 0; i < count; i++ {
			index := (offset + i) % length
			result[index] = pool[(baseHash[index]+byte(i))%byte(len(pool))]
		}
	}

	addCharacters(letters[26:], 1, 0) 
	addCharacters(numbers, 1, 1)      
	addCharacters(symbols, 1, 2)      

	if numSymbols > 0 {
		numSymbols -= 1
	}
	if numNumbers > 0 {
		numNumbers -= 1
	}

	addCharacters(symbols, numSymbols, 3)
	addCharacters(numbers, numNumbers, 4)

	for i := 0; i < length; i++ {
		if result[i] == 0 {
			result[i] = letters[baseHash[i]%byte(len(letters))]
		}
	}

	return string(result)
}
