package helpers

import (
	"errors"
	"strongify-passgen-go-echo/domain"
)

func ValidateUser(user domain.User) error {
	if user.Email == "" {
		return errors.New("invalid mail id")
	}
	if user.UserName == "" {
		return errors.New("invalid user  Name")
	}
	return nil
}
