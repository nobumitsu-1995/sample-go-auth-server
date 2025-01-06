package services

import (
	users "auth-server/Domain/Models/Users"
	"errors"
)

type IUserService interface {
	DuplicateEmail(email string) error
}

type UserService struct {
	userRepository users.IUserRepository
}

func (u *UserService) DuplicateEmail(email string) error {
	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}
	if user.ID != "" {
		return errors.New("user already exists")
	}
	return nil
}
