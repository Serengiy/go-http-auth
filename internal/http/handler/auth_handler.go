package handler

import (
	"auth_app/internal/dto"
	"auth_app/internal/service"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterAuthHandlers(r chi.Router, authService *service.AuthService) {
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		const op = "Register new user handler"
		var reqBody dto.RegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user, err := authService.RegisterUser(reqBody)
		if err != nil {
			var vErr service.ValidationError
			if errors.As(err, &vErr) {
				http.Error(w, vErr.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	})
}
