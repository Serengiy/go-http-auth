package service

import (
	"auth_app/internal/config"
	"auth_app/internal/dto"
	"auth_app/internal/http/validators"
	"auth_app/internal/models"
	"auth_app/internal/repository"
	"github.com/form3tech-oss/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type AuthService struct {
	rep *repository.Auth
}

type JWTClaims struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserId    uint   `json:"user_id"`
	ExpiresAt int64  `json:"expires_at"`
	Iat       int64  `json:"iat"`
	jwt.StandardClaims
}

func NewAuthService(authRepo *repository.Auth) *AuthService {
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

func (s *AuthService) LoginUser(reqBody dto.LoginRequest) (string, error) {
	const op = "auth service. LoginUser"

	if err := validators.ValidateStruct(reqBody); err != nil {
		return "", ValidationError{Message: "Validation failed: " + err.Error()}
	}

	user, err := s.rep.FindUserByEmail(reqBody.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", ValidationError{Message: "User not found"}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))
	if err != nil {
		return "", ValidationError{Message: "Password is incorrect"}
	}

	token, err := generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func generateToken(user *models.User) (string, error) {
	const op = "auth service. generateToken"

	secretKey := []byte(config.GetSecretKey())

	claims := JWTClaims{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserId:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Iat:       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)

	if err != nil {
		return "", nil
	}

	return signedToken, nil
}
