package user

import (
	"server/internal/domain/user/dtos"
)

type Service interface {
	Create(d *dtos.CreateUserDto) (*dtos.ResponseUserDto, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(d *dtos.CreateUserDto) (*dtos.ResponseUserDto, error) {
	user, err := CreateUser(d.ID, d.Name, d.Email, d.Password)

	if err != nil {
		return nil, err
	}

	err = s.Repository.Create(user)

	if err != nil {
		return nil, err
	}

	result := &dtos.ResponseUserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return result, nil
}
