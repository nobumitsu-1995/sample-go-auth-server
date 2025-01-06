package users

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

type Password struct {
	Value string
}

func NewPassword(value string) (Password, error) {
	password := Password{Value: value}
	if !password.isValid(value) {
		return Password{}, errors.New("password must be at least 8 characters long")
	}

	return Password{Value: value}, nil
}

func (p Password) isValid(value string) bool {
	return len(value) >= 8
}

func (p Password) Hash() string {
	hash := sha256.New()
	hash.Write([]byte(p.Value))

	return hex.EncodeToString(hash.Sum(nil))
}
