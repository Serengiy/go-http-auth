package permission

import "gorm.io/gorm"

type Permission struct {
	DB *gorm.DB
}

func NewPermissionService(db *gorm.DB) *Permission {
	return &Permission{DB: db}
}

func (p *Permission) Permission() (*Permission, error) {
	return nil, nil
}

func (p *Permission) Permissions() ([]*Permission, error) {
	return nil, nil
}

func CreatePermission() (*Permission, error) {
	return &Permission{}, nil
}

func DeletePermission() (*Permission, error) {
	return &Permission{}, nil
}
