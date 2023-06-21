package user

import (
	"fmt"
	"server/internal/domain/user"
	"server/internal/handlers"
	"server/internal/infra/database"
)

type FactoryUserImpl struct{}

func (f *FactoryUserImpl) Of() {
	userRepository := database.UserRepository{}

	userService := user.ServiceImpl{
		Repository: &userRepository,
	}

	userHandler := handlers.Handler{
		UserService: &userService,
	}

	fmt.Println(userHandler)
}
