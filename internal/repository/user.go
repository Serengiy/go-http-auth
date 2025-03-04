package repository

import (
	"auth_app/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (repo *User) CreateNewUser(user *models.User) error {
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}
	return repo.DB.Create(user).Error
}

func (repo *User) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
