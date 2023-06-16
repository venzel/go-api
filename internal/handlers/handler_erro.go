package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type HandleFn func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func HandlerError(handleFn HandleFn) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, status, err := handleFn(w, r)

		if err != nil {
			render.JSON(w, r, map[string]string{"error": err.Error()})
		}

		render.Status(r, status)

		if obj != nil {
			render.JSON(w, r, obj)
		}
	})
}
