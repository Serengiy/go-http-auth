package dto

type RoleRequest struct {
	Name string `json:"name" validate:"required,min=3,max=32,alphanum"`
}
