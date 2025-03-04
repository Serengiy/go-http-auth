package handler

import (
	"auth_app/internal/http/middleware"
	"auth_app/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func registerUserHandler(r chi.Router, userService *service.UserService) {
	r.Route("/me", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("User profile"))
			if err != nil {
				return
			}
		})
	})
}
