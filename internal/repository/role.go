package repository

import (
	"auth_app/internal/models"
	"fmt"
	"gorm.io/gorm"
)

type Role struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *Role {
	return &Role{db: db}
}

func (r *Role) FindRoleByName(name string) (*models.Role, error) {
	const op = "repository.role.FindRoleByName"
	var role models.Role

	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *Role) CreateRole(role *models.Role) error {
	const op = "repository.role.CreateRole"
	if role == nil {
		return fmt.Errorf("role is nil")
	}

	return r.db.Create(role).Error
}

func (r *Role) FindRoleById(id int64) (*models.Role, error) {
	const op = "repository.role.FindRoleById"
	var role models.Role
	if err := r.db.Where("id=?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *Role) DeleteRole(role *models.Role) error {
	const op = "repository.role.DeleteRole"
	if role == nil {
		return fmt.Errorf("role is nil")
	}
	return r.db.Delete(role).Error
}

func (r *Role) AttachPermission(role *models.Role, permission *models.Permission) error {
	const op = "repository.role.AttachPermission"
	if role == nil || permission == nil {
		return fmt.Errorf("role is nil")
	}

	err := r.db.Model(role).Association("Permissions").Append(permission)
	if err != nil {
		return err
	}
	return nil
}
