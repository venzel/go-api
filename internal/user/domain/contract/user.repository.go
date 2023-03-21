package contract

import "server/internal/user/domain"

type UserRepository interface {
	Create(user *domain.User) error
}
