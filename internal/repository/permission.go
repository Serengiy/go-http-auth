package repository

import (
	"auth_app/internal/dto"
	"auth_app/internal/models"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Permission struct {
	DB *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *Permission {
	return &Permission{
		DB: db,
	}
}

func (p *Permission) FindPermissionByName(name string) (*models.Permission, error) {
	var permission models.Permission

	if err := p.DB.Where("name = ?", name).First(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (p *Permission) FindPermissionByID(id int64) (*models.Permission, error) {
	var permission models.Permission

	if err := p.DB.Where("id = ?", id).First(&permission).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}

func (p *Permission) CreatePermission(perm *models.Permission) error {
	if perm == nil {
		return fmt.Errorf("permission is nil")
	}

	return p.DB.Create(perm).Error
}

func (p *Permission) GetPermissions(filter *dto.PermissionFilterStruct) ([]models.Permission, int64, error) {
	var permissions []models.Permission
	var totalRecords int64

	query := p.DB.Model(&models.Permission{})

	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if err := query.Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	offset := (filter.Page - 1) * filter.PerPage
	query = query.Offset(offset).Limit(filter.PerPage)

	log.Printf("Paginate: %d, PerPage: %d, Offset: %d", filter.Page, filter.PerPage, offset)

	if err := query.Find(&permissions).Error; err != nil {
		return nil, 0, err
	}
	return permissions, totalRecords, nil
}

func (p *Permission) DeletePermissionByID(id int64) error {
	if err := p.DB.Where("id = ?", id).Delete(&models.Permission{}).Error; err != nil {
		return err
	}
	return nil
}
