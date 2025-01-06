package Users

import users "auth-server/Domain/Models/Users"

type IUserService interface {
	CreateUser(user users.User) error
	UpdateUser(user users.User) error
	DeleteUser(user users.User) error
}

type UserService struct {
	userRepository users.IUserRepository
	userFactory    users.IUserFactory
}

func (us *UserService) CreateUser(email string, password string) error {
	user, err := us.userFactory.Create(
		users.Email{Value: email},
		users.Password{Value: password},
	)
	if err != nil {
		return err
	}
	err = us.userRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdateUser(id string, name string, email string, password string) error {
	user, err := us.userRepository.FindByID(id)
	if err != nil {
		return err
	}

	user.UpdateName(users.Name{Value: name})
	user.UpdateEmail(users.Email{Value: email})
	user.UpdatePassword(users.Password{Value: password})

	err = us.userRepository.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) DeleteUser(id string) error {
	err := us.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
