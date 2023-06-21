package user

type Service interface {
	Create(d *CreateUserDto) (*ResponseUserDto, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(d *CreateUserDto) (*ResponseUserDto, error) {
	id, name, email, password := d.ID, d.Name, d.Email, d.Password

	user, err := CreateUser(id, name, email, password)

	if err != nil {
		return nil, err
	}

	err = s.Repository.Create(user)

	if err != nil {
		return nil, err
	}

	result := &ResponseUserDto{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return result, nil
}
