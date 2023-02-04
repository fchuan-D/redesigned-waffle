package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"username" json:"user_name"`
	Password string `gorm:"password" json:"password"`
	Role     string `gorm:"role" json:"role"`
}

// TableName 修改表名映射
func (User User) TableName() string {
	return "testUser"
}
