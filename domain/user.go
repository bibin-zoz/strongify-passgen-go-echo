package domain

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type WordPhrase struct {
	ID         int
	UserID     int
	PhraseType string
	Phrase     string
}
