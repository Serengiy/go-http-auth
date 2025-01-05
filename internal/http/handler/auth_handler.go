package handler

import (
	"auth_app/internal/app/auth"
	"auth_app/internal/dto"
	"auth_app/internal/http/validators"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterAuthHandlers(r chi.Router, authService *auth.Auth) {
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var reqBody dto.RegisterRequest

		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := validators.ValidateStruct(reqBody); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := authService.Register(reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
