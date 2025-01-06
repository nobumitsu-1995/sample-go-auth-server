package users

import (
	"errors"
	"regexp"
)

type Email struct {
	Value string
}

func NewEmail(address string) (Email, error) {
	email := Email{Value: address}
	if !email.isValid(address) {
		return Email{}, errors.New("invalid email address")
	}

	return Email{Value: address}, nil
}

func (e Email) isValid(address string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(address)
}
