package users

import "errors"

type Name struct {
	Value string
}

func (n Name) NewName(value string) (Name, error) {
	if !n.isValid(value) {
		return Name{}, errors.New("name must be at least 3 characters long")
	}
	return Name{Value: value}, nil
}

func (n Name) isValid(value string) bool {
	return len(value) >= 3
}
