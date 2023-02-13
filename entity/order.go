package entity

import (
	"gorm.io/gorm"
)

type Oreder struct {
	gorm.Model
	Status bool `gorm:"comment:支付状态"`
	Amount int  `gorm:"comment:订单金额"`
	Charge uint `gorm:"comment:充电时长"`
	UserID uint `gorm:"comment:订单用户"`
	User   User `gorm:"foreignKey:UserID"`
}

// TableName 修改表名映射
func (Oreder) TableName() string {
	return "orders"
}
