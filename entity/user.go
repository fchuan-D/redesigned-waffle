package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"username"`
	Password string `gorm:"password"`
	Role     string `gorm:"role"`
}

// TableName 修改表名映射
func (User User) TableName() string {
	return "testUser"
}
