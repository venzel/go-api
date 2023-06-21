package handlers

import (
	"net/http"
	"server/internal/domain/user"

	"github.com/go-chi/render"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var createUserDto user.CreateUserDto

	render.DecodeJSON(r.Body, &createUserDto)

	result, err := h.UserService.Create(&createUserDto)

	if err != nil {
		return nil, 500, err
	}

	return result, 201, nil
}
