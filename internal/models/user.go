package models

type User struct {
	ID        uint   `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"size:255" json:"first_name"`
	LastName  string `gorm:"size:255" json:"last_name"`
	Email     string `gorm:"size:255" json:"email"`
	Password  string `gorm:"size:255" json:"password"`
	Phone     string `gorm:"size:255" json:"phone"`
	Roles     []Role `gorm:"many2many:user_roles;" json:"roles"`
}

func (User) TableName() string {
	return "users"
}
