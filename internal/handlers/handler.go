package handlers

import "server/internal/domain/user"

type Handler struct {
	UserService user.Service
}
