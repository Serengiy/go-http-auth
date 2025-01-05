package auth

import (
	"auth_app/internal/dto"
	"auth_app/internal/models"
	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *Auth {
	return &Auth{DB: db}
}

func (a *Auth) Register(r dto.RegisterRequest) (*models.User, error) {
	return nil, nil
}

func (a *Auth) Authenticate(username, password string) error {
	return nil
}
