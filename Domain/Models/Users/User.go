package users

type User struct {
	ID       string
	Name     Name
	Email    Email
	Password Password
}

type IUserRepository interface {
	FindByID(id string) (User, error)
	FindByEmail(email Email) (User, error)
	Save(user User) error
	Delete(id string) error
}

type IUserFactory interface {
	Create(email Email, password Password) (User, error)
}

func (u *User) UpdateName(newName Name) error {
	u.Name = newName
	return nil
}

func (u *User) UpdateEmail(newEmail Email) error {
	u.Email = newEmail
	return nil
}

func (u *User) UpdatePassword(newPassword Password) error {
	u.Password = newPassword
	return nil
}
