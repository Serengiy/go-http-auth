package models

type Role struct {
	ID          uint         `gorm:"primary_key;auto_increment" json:"id"`
	Name        string       `gorm:"size:255;not null;unique" json:"name"`
	Permissions []Permission `gorm:"many2many:permissions_roles;" json:"permissions"`
}
