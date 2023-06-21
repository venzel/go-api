package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(handleFunc HandleFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := handleFunc(w, r)

		if err != nil {
			render.JSON(w, r, map[string]string{"error": err.Error()})
		}

		render.Status(r, status)

		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
