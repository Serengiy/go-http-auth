package dto

type RegisterRequest struct {
	FirstName       string `json:"first_name" validate:"required,min=3,max=20,alphanum"`
	LastName        string `json:"last_name" validate:"required,min=3,max=20,alphanum"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=20"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	Phone           string `json:"phone" validate:"required,max=20"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}
