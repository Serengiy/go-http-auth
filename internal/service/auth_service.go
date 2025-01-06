package service

import (
	"auth_app/internal/dto"
	"auth_app/internal/http/validators"
	"auth_app/internal/models"
	"auth_app/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService struct {
	rep *repository.User
}

func NewAuthService(authRepo *repository.User) *AuthService {
	return &AuthService{
		rep: authRepo,
	}
}

func (s *AuthService) RegisterUser(reqBody dto.RegisterRequest) (*models.User, error) {
	const op = "auth service. RegisterUser"

	if err := validators.ValidateStruct(reqBody); err != nil {
		return nil, ValidationError{Message: "Validation failed: " + err.Error()}
	}

	userExist, err := s.rep.FindUserByEmail(reqBody.Email)
	if userExist != nil {
		return nil, ValidationError{Message: "User already exists"}
	}

	hashedPassword, err := hashPassword(reqBody.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		FirstName: reqBody.FirstName,
		LastName:  reqBody.LastName,
		Email:     reqBody.Email,
		Phone:     reqBody.Phone,
		Password:  hashedPassword,
	}

	err = s.rep.CreateNewUser(user)
	if err != nil {
		return nil, err
	}

	log.Printf("User has been created: %v", user)

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
