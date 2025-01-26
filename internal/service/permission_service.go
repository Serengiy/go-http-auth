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

func (s *PermissionService) GetPermissionsByFilter(filters dto.PermissionFilterStruct) ([]models.Permission, int64, error) {
	const op = "service.PermissionService.GetPermissionsByFilter"

	permissions, totalRecords, err := s.rep.GetPermissions(&filters)
	if err != nil {
		return nil, 0, err
	}
	return permissions, totalRecords, nil
}

func (s *PermissionService) GetPermissionById(id int64) (*models.Permission, error) {
	const op = "service.PermissionService.GetPermissionById"

	permission, err := s.rep.FindPermissionByID(id)

	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (s *PermissionService) DeletePermissionById(permission *models.Permission) error {
	const op = "service.PermissionService.DeletePermissionById"

	err := s.rep.DeletePermissionByID(int64(permission.ID))
	if err != nil {
		return err
	}
	return nil
}
