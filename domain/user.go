package domain

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName string
	Email    string
	Password string
}
type Phrase struct {
	ID       int
	Category string
	Phrase   []string
}
