package service

import "auth_app/internal/app/auth"

type AuthService struct {
	authRepo *auth.Auth
}

func NewAuthService(authRepo *auth.Auth) *AuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}

func (s *AuthService) RegisterUser(username, password string) (string, error) {
	return s.RegisterUser(username, password)
}
