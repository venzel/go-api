package user

import (
	"errors"
	"server/internal/register/domain/notification"
	"server/internal/register/domain/validator"

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

	if validator.Id(user.ID) {
		notification.AddError("Id is invalid")
	}

	if validator.Name(user.Name) {
		notification.AddError("Name is invalid")
	}

	if validator.Email(user.Email) {
		notification.AddError("Email is invalid")
	}

	if validator.Password(user.Password) {
		notification.AddError("Password is invalid")
	}

	if notification.HasError() {
		return errors.New(notification.GetErrors())
	}

	return nil
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
