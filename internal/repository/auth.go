package repository

import (
	"auth_app/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type Auth struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *Auth {
	return &Auth{
		DB: db,
	}
}

func (repo *Auth) CreateNewUser(user *models.User) error {
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}
	return repo.DB.Create(user).Error
}

func (repo *Auth) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
