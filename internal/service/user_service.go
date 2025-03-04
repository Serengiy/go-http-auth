package service

import (
	"auth_app/internal/repository"
)

type UserService struct {
	rep *repository.User
}

func NewUserService(userRepo *repository.User) *UserService {
	return &UserService{
		rep: userRepo,
	}
}
