package handler

import (
	"auth_app/internal/app/role"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterRoleHandlers(r chi.Router, s *role.Role) {
	r.Get("/role", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			panic("panic")
		}
	})
}
