package role

import "gorm.io/gorm"

type Role struct {
	DB *gorm.DB
}

func NewRoleService(db *gorm.DB) *Role {
	return &Role{DB: db}
}

func (*Role) CreateRole() *Role {
	return &Role{}
}

func (*Role) AttachRole() *Role {
	return &Role{}
}

func (*Role) DetachRole() *Role {
	return &Role{}
}
