package users

type User struct {
	ID       string
	Email    Email
	Password Password
}

type IUserRepository interface {
	FindByID(id string) (User, error)
	FindByEmail(email string) (User, error)
	Create(user User) error
	Update(user User) error
	Delete(user User) error
}

func (u *User) UpdateEmail(newEmail Email) error {
	u.Email = newEmail
	return nil
}

func (u *User) UpdatePassword(newPassword Password) error {
	u.Password = newPassword
	return nil
}
