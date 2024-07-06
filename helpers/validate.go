package helpers

import (
	"errors"
	"regexp"
	"strongify-passgen-go-echo/domain"
)

func ValidateUser(user domain.User) error {
	if user.Email == "" {
		return errors.New("invalid mail id")
	}
	emailPattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(emailPattern).MatchString(user.Email) {
		return errors.New("invalid mail id")
	}
	if user.UserName == "" {
		return errors.New("invalid user  Name")
	}
	return nil
}
