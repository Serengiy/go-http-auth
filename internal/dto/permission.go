package dto

type PermissionStoreRequest struct {
	Name string `json:"name" validate:"required,min=3,max=20,alphanum"`
}

type PermissionFilterStruct struct {
	PerPage int    `json:"per_page" `
	Page    int    `json:"page" `
	Name    string `json:"name" `
}
