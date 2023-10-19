package valueObject

import (
	"errors"
	"regexp"
)

var (
	invalidEmail     = errors.New("invalid email")
	validEmailRegExp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type Email struct {
	Email string
}

func valid(email string) bool {
	emailRegex := validEmailRegExp
	return emailRegex.MatchString(email)
}

func NewEmail(email string) (Email, error) {
	if !valid(email) {
		return Email{}, invalidEmail
	}
	return Email{
		Email: email,
	}, nil
}
