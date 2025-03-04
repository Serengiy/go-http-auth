package service

import (
	"auth_app/internal/dto"
	"auth_app/internal/http/validators"
	"auth_app/internal/models"
	"auth_app/internal/repository"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type RoleService struct {
	rep *repository.Role
}

func NewRoleService(repo *repository.Role) *RoleService {
	return &RoleService{
		rep: repo,
	}
}

func (s *RoleService) CreateNewRole(request dto.RoleRequest) (*models.Role, error) {
	const op = "RoleService.CreateNewRole"

	err := validators.ValidateStruct(request)
	if err != nil {
		return nil, ValidationError{Message: "Validation failed: " + err.Error()}
	}

	roleExists, err := s.rep.FindRoleByName(request.Name)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if roleExists != nil {
		return nil, ValidationError{Message: fmt.Sprintf("Role with name %s already exists", request.Name)}
	}

	role := &models.Role{
		Name: request.Name,
	}

	err = s.rep.CreateRole(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleService) FindRoleByID(id int64) (*models.Role, error) {
	const op = "RoleService.FindRoleByID"

	role, err := s.rep.FindRoleById(id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RoleService) DeleteRole(role *models.Role) error {
	const op = "RoleService.DeleteRole"
	if role == nil {
		return fmt.Errorf("role is nil")
	}

	err := s.rep.DeleteRole(role)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleService) AttachPermission(role *models.Role, permission *models.Permission) error {
	const op = "RoleService.AttachPermission"
	if role == nil || permission == nil {
		return fmt.Errorf("role or permission is nil")
	}

	err := s.rep.AttachPermission(role, permission)
	if err != nil {
		return err
	}
	return nil
}
