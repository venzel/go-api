package user

import (
	"errors"
	"server/internal/register/domain/notification"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func CreateUser(id string, name string, email string, password string) (*User, error) {
	user := &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.isValid()

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) isValid() error {
	notification := notification.CreateNotification("user")

	if message, valid := user.isIdValid(user.ID); !valid {
		notification.AddError(message)
	}

	if message, valid := user.isNameValid(user.Name); !valid {
		notification.AddError(message)
	}

	if message, valid := user.isEmailValid(user.Email); !valid {
		notification.AddError(message)
	}

	if message, valid := user.isPasswordValid(user.Password); !valid {
		notification.AddError(message)
	}

	if notification.HasError() {
		return errors.New(notification.GetErrors())
	}

	return nil
}

func (user *User) isIdValid(id string) (string, bool) {
	if id == "" {
		return "Id is invalid", false
	}

	return "", true
}

func (user *User) isNameValid(name string) (string, bool) {
	if name == "" {
		return "Name is invalid", false
	}

	return "", true
}

func (user *User) isEmailValid(email string) (string, bool) {
	if email == "" {
		return "Email is invalid", false
	}

	return "", true
}

func (user *User) isPasswordValid(password string) (string, bool) {
	if password == "" {
		return "Password is invalid", false
	}

	return "", true
}

func (user *User) EncryptPassword() error {
	password := []byte(user.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	err = user.isValid()

	if err != nil {
		return err
	}

	return nil
}
