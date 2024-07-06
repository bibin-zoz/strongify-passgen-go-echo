package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	db "strongify-passgen-go-echo/database"
	"strongify-passgen-go-echo/domain"
	"strongify-passgen-go-echo/helpers"
	models "strongify-passgen-go-echo/model"
	views "strongify-passgen-go-echo/view"

	"github.com/labstack/echo/v4"
)

func GetPhrases(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := helpers.GetTokenFromHeader(authHeader)

	userID, _, err := helpers.ExtractUserIDFromToken(token)
	if err != nil {
		fmt.Println("err", err)
		return views.RenderError(c, "Invalid token")
	}

	var phrases []domain.WordPhrase
	if err := db.DB.Where("user_id = ?", userID).Find(&phrases).Error; err != nil {
		return views.RenderError(c, "failed to retrieve phrases")
	}

	return c.JSON(http.StatusOK, phrases)
}
func AddPhrase(c echo.Context) error {
	var phrase models.WordPhrase
	if err := c.Bind(&phrase); err != nil {
		log.Println("failed to bind data",err)
		return views.RenderError(c, err.Error())
	}

	if len(phrase.Words) < 5 {
		return views.RenderError(c, "phrase must contain at least 5 words")
	}

	authHeader := c.Request().Header.Get("Authorization")
	token := helpers.GetTokenFromHeader(authHeader)

	userID, _, err := helpers.ExtractUserIDFromToken(token)
	if err != nil {
		fmt.Println("err", err)
		return views.RenderError(c, "Invalid token")
	}
	joinedString := strings.Join(phrase.Words, " ")
	userPhrase := &domain.WordPhrase{
		UserID:     int(userID),
		PhraseType: phrase.PhraseType,
		Phrase:     joinedString,
	}
	if err := db.DB.Create(&userPhrase).Error; err != nil {
		return views.RenderError(c, "could not store phrase in database")
	}

	return views.RenderSuccess(c, "phrase added successfully")
}
