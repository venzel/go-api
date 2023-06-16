package database

import (
	"server/internal/domain/user"
)

type UserRepository struct{}

func (r *UserRepository) Create(user *user.User) error {
	return nil
}
