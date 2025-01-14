package models

type Permission struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"size:255;not null;unique" json:"name"`
}

func (Permission) TableName() string {
	return "permissions"
}
