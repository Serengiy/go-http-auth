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

type PermissionService struct {
	rep *repository.Permission
}

func NewPermissionService(repo *repository.Permission) *PermissionService {
	return &PermissionService{
		rep: repo,
	}
}

func (s *PermissionService) CreateNewPermission(reqBody *dto.PermissionStoreRequest) (*models.Permission, error) {
	const op = "service.PermissionService.CreateNewPermission"

	if err := validators.ValidateStruct(reqBody); err != nil {
		return nil, ValidationError{Message: "Validation failed: " + err.Error()}
	}

	permissionExists, err := s.rep.FindPermissionByName(reqBody.Name)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if permissionExists != nil {
		return nil, ValidationError{Message: fmt.Sprintf("%s has already been created", reqBody.Name)}
	}

	permission := &models.Permission{
		Name: reqBody.Name,
	}

	err = s.rep.CreatePermission(permission)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *PermissionService) GetPermissionsByFilter(filters dto.PermissionFilterStruct) ([]models.Permission, error) {
	const op = "service.PermissionService.GetPermissionsByFilter"

	permissions, err := s.rep.GetPermissions(&filters)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
